package controllers

import "net/http"

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("Creating a new user"))
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
