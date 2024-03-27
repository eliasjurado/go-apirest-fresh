package repository

import (
	"apirest/models"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init(connection *sql.DB) {
	db = connection
}


// Obtener todo el registro
func GetAllUsers() models.Users {
	sql := "SELECT id, username, password, email FROM users"
	users := models.Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

//Obtener un Registro
func GetOneUser(id int) *models.User {
	user := models.NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}


// func GetOneUser(id int) {
// 	query := "select * from contact where id = ?"
// 	row := db.QueryRow(query, id)

// 	c := models.Contact{}

// 	log.Printf("%v\n", "Lista de Contacto")
// 	log.Printf("%v\n", "------------------")
// 	var valueEmail sql.NullString
// 	var valuePhone sql.NullString
// 	err := row.Scan(&c.Id, &c.Name, &valueEmail, &valuePhone)

// 	if valueEmail.Valid {
// 		c.Email = valueEmail.String
// 	} else {
// 		c.Email = ""
// 	}

// 	if valuePhone.Valid {
// 		c.Phone = valuePhone.String
// 	} else {
// 		c.Phone = ""
// 	}
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			log.Fatalf("No se encontró ningún contacto con el ID %d", id)
// 		}
// 	}

// 	log.Printf("ID: %v, Name: %v, Email: %v, Phone: %v\n", c.Id, c.Name, c.Email, c.Phone)
// 	log.Printf("%v\n", "------------------")
// }

func CreateUser(c models.Contact) {
	query := "insert into contact (name, email, phone) values(?,?,?)"

	_, err := db.Exec(query, c.Name, c.Email, c.Phone)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", "Nuevo contacto registrado con éxito")
}

func UpdateUser(c models.Contact) {
	query := "update contact set name =?, email=?, phone=? where id=?"

	_, err := db.Exec(query, c.Name, c.Email, c.Phone, c.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", "Contacto actualizado con éxito")
}

func DeleteUser(id int) {
	query := "delete from contact where id=?"

	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", "Contacto eliminado con éxito")
}
