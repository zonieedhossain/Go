package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {

	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed Database")
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("Database could not connect")
	}

	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("Database could not connect")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Created Success")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("Database could not connect")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name : ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User successfully Deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("Database could not connect")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name : ?", name).Find(&user)

	user.Email = email

	db.Save(&user)

	fmt.Fprintf(w, "User successfully Updated")
}
