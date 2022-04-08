package controllers

import (
	"api/src/auth"
	"api/src/database"
	responses "api/src/http-responses"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Posts
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorId = userId

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	post.Id, err = repository.Create(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {

}

func GetPostById(w http.ResponseWriter, r *http.Request) {

}

func UpdatePostById(w http.ResponseWriter, r *http.Request) {

}

func RemovePostById(w http.ResponseWriter, r *http.Request) {

}
