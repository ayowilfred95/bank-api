package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ayowilfred95/helpers"
	"github.com/ayowilfred95/users"
	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email string
	Password string
}

type ErrResponse struct{
	Message string
}

func login(w http.ResponseWriter, r*http.Request) {
	// Read the body

	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)


	// handle Login
	var formattedBody Login
	err = json.Unmarshal(body,&formattedBody)
	helpers.HandleErr(err)
	login:= users.Login(formattedBody.Username, formattedBody.Password)

	// prepare response 
	if login["message"] =="all is fine" {
		resp:=login
		json.NewEncoder(w).Encode(resp)
	} else {
		// handle error
		resp:= ErrResponse{Message: "Wrong username and password"}
		json.NewEncoder(w).Encode(resp)
	}

}


// register route function
func register(w http.ResponseWriter, r*http.Request) {
	// Read the body

	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)


	// handle Register
	var formattedBody Register
	err = json.Unmarshal(body,&formattedBody)
	helpers.HandleErr(err)
	register:= users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	// prepare response 
	if register["message"] =="all is fine" {
		resp:=register
		json.NewEncoder(w).Encode(resp)
	} else {
		// handle error
		resp:= ErrResponse{Message: "Cannot perform registration"}
		json.NewEncoder(w).Encode(resp)
	}

}
 



// create our route
func StartApi() {
	router:= mux.NewRouter()

	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888",router))

}

