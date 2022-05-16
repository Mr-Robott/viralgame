package server

import (
	"github.com/viralgame/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/viralgame/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
	log    *log.Logger
}

var Handler Server

func (s *Server) setupHandlers() {

	s.Router.HandleFunc("/user", SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/user/{id}/state", SetMiddlewareJSON(s.UpdateGameState)).Methods("PUT")
	s.Router.HandleFunc("/user/{id}/state", SetMiddlewareJSON(s.LoadGameState)).Methods("GET")
	s.Router.HandleFunc("/user/{id}/friends", SetMiddlewareJSON(s.AddFriends)).Methods("PUT")
	s.Router.HandleFunc("/user/{id}/friends", SetMiddlewareJSON(s.GetFriends)).Methods("GET")
	s.Router.HandleFunc("/user", SetMiddlewareJSON(s.GetAllUsers)).Methods("GET")

}

func (s *Server) Run() {

	db := &server.DatabaseConfig{}
	s.log = log.New(os.Stdout, "Viral Game ", log.LstdFlags)
	s.Router = mux.NewRouter().StrictSlash(true)
	s.setupHandlers()

	db.InitializeConfig()
	s.ConnectDB(db.GetConnectionString(), db.DBDriver)
	s.log.Println("Starting server at 5000")
	s.log.Fatal(http.ListenAndServe(":5000", s.Router))
}

func (s *Server) ConnectDB(connection, driver string) {
	var err error
	s.DB, err = gorm.Open(driver, connection)
	if err != nil {
		s.log.Fatalf("Unable to connect %s database : %s", err.Error())
	} else {
		s.log.Println("Database Connected")
	}
	s.DB.AutoMigrate(&models.User{}, &models.UserFriends{})
}
