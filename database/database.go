package database

import (
	"database/sql"
	"fmt"
	"log"
)

func MyCreateTable(schema string) {
	DB.Exec(schema)
}

// Ferificar si una tabla existe o no
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

// Crear una tabla en la base de datos
func CreateTable(schema, name string) {

	if !ExistsTable(name) {
		_, err := DB.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Eliminara Tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	log.Printf("%+v\n",sql )
	DB.Exec(sql)
}

// Polimorfismo a Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := DB.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// Polimorfismo a Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
