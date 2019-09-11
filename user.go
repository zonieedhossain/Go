package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	user  string
	email string
}

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed Database")
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("test", "gorm.db")
	if err != nil {
		panic("Database could not connect")
	}

	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create New User")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User")
}
