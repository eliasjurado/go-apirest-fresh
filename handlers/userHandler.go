package handlers

import (
	"apirest/database"
	"apirest/models"
	"apirest/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	
	format := r.Header.Get("x-format")

	database.Connect()
	users := repository.GetAllUsers()
	models.SendData(w, users, format)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	format := r.Header.Get("x-format")
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		models.SendNotProcesableEntity(w, format)
		return
	}

	database.Connect()
	user := repository.GetOneUser(id)
	if user.Id == 0 {
		models.SendNotFound(w, format)
		return
	}

	models.SendData(w, user, format)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintln(w, http.StatusUnprocessableEntity)
		return
	}
	database.Connect()
	user.Save(database.DB)
	database.Connect()
	output, _ := json.Marshal(user)
	fmt.Fprintln(w, string(output))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintln(w, http.StatusUnprocessableEntity)
		return
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		fmt.Fprintln(w, http.StatusUnprocessableEntity)
		return
	}

	user.Id = int64(id)
	database.Connect()
	user.Save(database.DB)
	database.Connect()
	output, _ := json.Marshal(user)
	fmt.Fprintln(w, string(output))
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintln(w, http.StatusUnprocessableEntity)
		return
	}
	database.Connect()
	user, _ := models.GetUser(database.DB, id)
	user.Delete(database.DB)
	database.Close()
	output, _ := json.Marshal(user)
	fmt.Fprintln(w, string(output))
}
