package handlers

import (
	"apirest/database"
	"apirest/models"
	"apirest/repository"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	format := vars["format"]

	database.Connect()
	repository.Init(database.DB)
	users := repository.GetAllUsers()

	var output []byte

	switch format {
	case "xml":
		output, _ = xml.Marshal(users)
		w.Header().Set("Content-Type", "text/xml")
	case "yaml":
		output, _ = yaml.Marshal(users)
	default:
		output, _ = json.Marshal(users)
		w.Header().Set("Content-Type", "application/json")
	}

	database.Close()
	fmt.Fprintln(w, string(output))
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
	user := models.GetUser(database.DB, id)
	user.Delete(database.DB)
	database.Close()
	output, _ := json.Marshal(user)
	fmt.Fprintln(w, string(output))
}
