package main

import (
	"fmt"
    "log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var Users []User

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Users)
}

func returnUserById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: " + key)
}

func handleRequests() {
	// Replace this routing in favour of mux
    // http.HandleFunc("/", homePage)
	// http.HandleFunc("/users", returnAllUsers)
	
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsers)
	myRouter.HandleFunc("/users/{id}", returnUserById)

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Users = []User{
		User{Id: "1", Name: "Sam", Surname: "Wells", Age:"25"},
		User{Id: "2", Name: "Sarah", Surname: "Scott", Age:"26"},
	}

    handleRequests()
}

type User struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Surname string `json:"Surname"`
	Age string `json:"Age"`
}
