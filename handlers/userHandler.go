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

	vars := mux.Vars(r)
	format := vars["format"]

	database.Connect()
	repository.Init(database.DB)
	users := repository.GetAllUsers()
	// if users==nil || len(users)<=0{
		models.SendData(w,users,format)
	// }
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	database.Connect()
	repository.Init(database.DB)
	user := repository.GetOneUser(id)
	output, _ := json.Marshal(user)

	database.Close()
	fmt.Fprintln(w, string(output))
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
	user,_ := models.GetUser(database.DB, id)
	user.Delete(database.DB)
	database.Close()
	output, _ := json.Marshal(user)
	fmt.Fprintln(w, string(output))
}
