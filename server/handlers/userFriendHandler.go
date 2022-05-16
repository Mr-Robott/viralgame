package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viralgame/models"
)

type FriendsID struct {
	Friends []string `json:"friends"`
}

func (s *Server) AddFriends(w http.ResponseWriter, r *http.Request) {

	// reading query params
	vars := mux.Vars(r)

	// check if valid params exist and so is user with id param
	user := s.ValidateParams(w, vars)

	// if user exist ready body data for updated fields
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
	}

	var allFriends FriendsID
	err = json.Unmarshal(body, &allFriends)
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}

	for _, ids := range allFriends.Friends {
		var u models.User
		_, err = u.FindUserByID(s.DB, ids)
		if err != nil {
			s.ToError(w, http.StatusUnprocessableEntity, err)
			return
		}

		uf := models.UserFriends{}
		uf.PopulateUserFriend()
		uf.UserOne = user.UserDetails.ID
		uf.UserTwo = ids
		_, err = uf.SaveUserFriend(s.DB)
		if err != nil {
			s.ToError(w, http.StatusUnprocessableEntity, err)
			return
		}
	}
	// return response with appropriate headers
	s.ToJSON(w, http.StatusCreated, user)
}

func (s *Server) GetFriends(w http.ResponseWriter, r *http.Request) {

	// reading query params
	vars := mux.Vars(r)

	// check if valid params exist and so is user with id param
	user := s.ValidateParams(w, vars)

	uf, err := models.FindFriendsByID(s.DB, user.UserDetails.ID)
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}

	// filter friends and user from records
	var friends []string
	for idx := range uf {
		if uf[idx].UserOne == user.UserDetails.ID {
			friends = append(friends, uf[idx].UserTwo)
		} else {
			friends = append(friends, uf[idx].UserOne)
		}
	}

	ufs, err := user.GetAllFriends(s.DB, friends)
	if err != nil {
		s.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}

	// return response with appropriate headers
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, user.UserDetails.ID))
	s.ToJSON(w, http.StatusCreated, ufs)
}
