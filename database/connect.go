package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Guarda la conexion
var DB *sql.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	dns := fmt.Sprintf("%v:%v@(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	//Abrir conexión a la database
	connection, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = connection

	//Verificar conexión a la db
	Ping()

	log.Printf("%+v\n", "conexion exitosa")
	
}

// Verificar la conexion
func Ping() {
	if err := DB.Ping(); err != nil {
		log.Fatal(err.Error())
	}
}

// Cerrar la Conexion
func Close() {
	DB.Close()
}
