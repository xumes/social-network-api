package controllers

import (
	"api/src/controllers/repositories"
	"api/src/database"
	"api/src/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(201)
	w.Write([]byte(fmt.Sprintf("Inserted id: %d", userId)))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Get all users"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Get user by id"))
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Updating user by id"))
}

func RemoveUserById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("deleting an user by Id"))
}
