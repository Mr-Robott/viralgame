package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viralgame/models"
)

type AllUsers struct {
	Users []models.UserDetails `json:"users"`
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	// reading request body of the POST request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
	}

	// convert body data into concrete form
	userDetails := models.UserDetails{}
	err = json.Unmarshal(body, &userDetails)
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}

	// initialize and validate user
	user := models.User{}
	user.UserDetails = userDetails
	err = user.ValidateCreateUser()
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.PopulateUser()

	// create the record in user table
	userCreated, err := user.SaveUser(s.DB)
	if err != nil {
		s.ToError(w, http.StatusInternalServerError, err)
		return
	}

	s.log.Println("Successfully Created User :",user.UserDetails.Name)
	// return response with appropriate headers
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, userDetails.ID))
	s.ToJSON(w, http.StatusCreated, userCreated.UserDetails)
}

func (s *Server) UpdateGameState(w http.ResponseWriter, r *http.Request) {

	// reading query params
	vars := mux.Vars(r)

	// check if valid params exist and so is user with id param
	user := s.ValidateParams(w, vars)

	// if user exist ready body data for updated fields
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
	}

	// convert body data into concrete form
	state := models.GameState{}
	err = json.Unmarshal(body, &state)
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}

	// validate the body data for constraints and integrity
	err = state.ValidateState()
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}

	// perform the update operation on database
	user.GameState = state
	user, err = user.UpdateAUser(s.DB, user.UserDetails.ID)
	if err != nil {
		s.ToError(w, http.StatusInternalServerError, err)
		return
	}

	s.log.Println("Game State Updated For User:",user.UserDetails.Name)
	// return response with appropriate headers
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, user.UserDetails.ID))
	s.ToJSON(w, http.StatusOK, user)

}

func (s *Server) LoadGameState(w http.ResponseWriter, r *http.Request) {
	var err error

	// reading query params
	vars := mux.Vars(r)

	// check if valid params exist and so is user with id param
	user := &models.User{}
	if uid, exist := vars["id"]; !exist {
		s.ToError(w, http.StatusUnprocessableEntity, errors.New("no id param found"))
		return
	} else {
		user, err = user.FindUserByID(s.DB, uid)
		if err != nil {
			s.ToError(w, http.StatusUnprocessableEntity, err)
			return
		}
	}

	s.log.Println("Game State Fetched For User:",user.UserDetails.Name)
	// return response with appropriate headers
	s.ToJSON(w, http.StatusOK, user.GameState)

}

func (s *Server) GetAllUsers(w http.ResponseWriter, _ *http.Request) {

	var userDetails AllUsers
	users, err := models.FindAllUsers(s.DB)
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}

	// update response with each user & its details
	for idx := range users {
		userDetails.Users = append(userDetails.Users, users[idx].UserDetails)
	}

	s.log.Println("Fetching Details For All User:")
	// return response with appropriate headers
	s.ToJSON(w, http.StatusOK, userDetails)
}
