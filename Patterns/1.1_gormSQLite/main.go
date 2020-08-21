/*

#########################
## Testing using Postman
#########################

	Can test using Postman with following calls:

	POST - create new user
	localhost:8001/user/jsmith/j@smith.com

	GET - get users
	localhost:8001/users

	DELETE - delete users
	http://localhost:8001/user/jsmith

	PUT - update users
	http://localhost:8001/user/jsmith/j@smith.com


#########################
## Testing using curl
#########################

	Get all users: curl http://localhost:8001/users
	Create a new user: curl -X POST http://localhost:8001/user/jsmith/j@smith.com
	Update user's email: curl -X PUT http://localhost:8001/user/jsmith/john@smith.com
	Delete User: curl -X DELETE http://localhost:8001/user/jsmith


#####################
## Database
#####################

	Using sqlite DB. The DB file is called test.db
	CAn use an app like TablePlus to access the DB

*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//homeHandler is executed when user hits "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

//handleRequest configures the router and starts the server
func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	//home
	myRouter.HandleFunc("/", homeHandler).Methods("GET")

	//various handlers to list, create, delete and update users
	//handlers are in user.go file
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	//start the server
	log.Fatal(http.ListenAndServe(":8001", myRouter))

}

func main() {

	fmt.Println("GORM & REST Tutorial...")

	//initial DB migration - creates initial tables
	//only if the tables don't already exist
	InitialMigration()

	//start the server and handle requests
	handleRequest()

}
