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

//User holds the user data
type User struct {
	gorm.Model
	Name  string
	Email string
}

//InitialMigration of the DB
func InitialMigration() {

	//connect to DB. fail in unable to connect
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	//setup the initial tables
	//does nothing is already there
	db.AutoMigrate(&User{})

}

//AllUsers returns all users
func AllUsers(w http.ResponseWriter, r *http.Request) {

	//connect to DB. fail in unable to connect
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to database")
	}
	defer db.Close()

	//define the slice of users and get from DB
	var users []User
	db.Find(&users)

	//return users
	json.NewEncoder(w).Encode(users)

}

//NewUser creates a new user
func NewUser(w http.ResponseWriter, r *http.Request) {

	//connect
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to database")
	}
	defer db.Close()

	//get the params
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	//create user object
	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "New User Sucessfully created")

}

//DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	//connect
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to database")
	}
	defer db.Close()

	//get the vars
	vars := mux.Vars(r)	
	name := vars["Name"]
	var user User

	//find and delete the user
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User sucessfully deleted")

}

//UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	//connect
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to database")
	}
	defer db.Close()

	//get vars
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User

	//find the user and update email form post
	db.Where("name = ?", name).Find(&user)
	user.Email = email

	//update the user in DB
	db.Save(&user)

	fmt.Fprintf(w, "Successfully updated the user")

}
