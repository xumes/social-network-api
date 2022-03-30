package controllers

import (
	"api/src/auth"
	"api/src/database"
	responses "api/src/http-responses"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userFromDatabase, err := repository.FindByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	if err = security.ComparePassword(userFromDatabase.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.CreateToken(userFromDatabase.Id)
	fmt.Println(token)

	responses.JSON(w, http.StatusNoContent, nil)
}
