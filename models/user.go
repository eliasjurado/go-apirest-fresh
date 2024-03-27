package models

import "database/sql"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE IF NOT EXISTS users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(64) NOT NULL,
	email VARCHAR(50),
	create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

// Construir usuario
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

// Crear usuario e insertar
func CreateUser(db *sql.DB, username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save(db)
	return user
}

// Insertar Registro
func (user *User) insert(db *sql.DB) {
	sql := "INSERT users SET username=?, password=?,email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId() //devuelve el id insertado
}

// Obtener todo el registro
func ListUsers(db *sql.DB) Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

// Obtener un Registro
func GetUser(db *sql.DB, id int) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

// Actualizar Registro
func (user *User) update(db *sql.DB) {
	sql := "UPDATE users SET username=?, password=?,email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

// Guardar o editar registro
func (user *User) Save(db *sql.DB) {
	if user.Id == 0 {
		user.insert(db)
	} else {
		user.update(db)
	}
}

// Eliminar registro
func (user *User) Delete(db *sql.DB) {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}
