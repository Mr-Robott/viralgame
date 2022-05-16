package server

import (
	"encoding/json"
	"errors"
	"github.com/viralgame/models"
	"net/http"
)

func (s *Server) ValidateParams(w http.ResponseWriter, vars map[string]string) *models.User {
	var err error
	// check if valid params exist and so is user with id param
	user := &models.User{}
	if uid, exist := vars["id"]; !exist || uid == "" {
		s.log.Println("Unable to Validate Params")
		s.ToError(w, http.StatusUnprocessableEntity, errors.New("no id param found"))
		return nil
	} else {
		user, err = user.FindUserByID(s.DB, uid)
		if err != nil {
			s.ToError(w, http.StatusUnprocessableEntity, err)
			return nil
		}
	}

	return user
}

func (s *Server) ToJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		s.log.Println(w, "%s", err.Error())
	}
}

func (s *Server) ToError(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		s.ToJSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	s.ToJSON(w, http.StatusBadRequest, nil)
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
