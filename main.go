package main

import (
	"apirest/database"
	"apirest/handlers"
	"apirest/repository"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(0)
	database.Connect()
	repository.Init(database.DB)

	mux := mux.NewRouter()	
	mux.HandleFunc("/api/user/", handlers.GetAllUsers).Methods("GET")	
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetOneUser).Methods("GET")	
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.RemoveUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))

	// database.Connect()

	// //Create tables
	// // models.Init(database.DB)
	// database.CreateTable(models.UserSchema, "user")

	// // database.TruncateTable("users")

	// //Create
	// // user := models.User{Username: "pretelcarla", Password:"1234567689", Email:"pretel.carla@gmail.com"}
	// // user.Save(database.DB)

	// //Update
	// // user := models.User{Id: 1,Username: "pretelcarla", Password:"1234567689", Email:"pretel.carla@gmail.com"}
	// // user.Save(database.DB)

	// //Delete
	// // user := models.User{Id: 1}
	// // user.Delete(database.DB)

	// log.Printf("%+v\n", models.ListUsers(database.DB))
	// log.Printf("%+v\n", models.GetUser(database.DB, 1))
	// defer database.Close()

}
