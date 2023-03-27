package userservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SubhakarBuddana/golang/dbconnection"
	"github.com/SubhakarBuddana/golang/types"
	"github.com/SubhakarBuddana/golang/validations"

	_ "github.com/go-sql-driver/mysql"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()

		var person types.Person
		err = json.Unmarshal(body, &person)
		if err != nil {
			panic(err)
		}

		//fmt.Println(p.Username, p.Firstname, p.Lastname)
		err = validations.SignUpValidation(person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = dbconnection.InsertUser(person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}

}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer r.Body.Close()
		var loginuser types.Login
		err = json.Unmarshal(body, &loginuser)
		if err != nil {
			panic(err)
		}
		err = validations.LoginRequestValidation(loginuser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = dbconnection.LoginUser(loginuser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		http.Error(w, "Invalid Method Request", http.StatusMethodNotAllowed)
		return
	}

}
func Chanpassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("varma")
	if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer r.Body.Close()
		var request types.ChangePasswordRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			panic(err)
		}
		err = validations.ChangePasswordRequestValidation(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = dbconnection.UserChangePassword(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		http.Error(w, "Invalid Method Request", http.StatusMethodNotAllowed)
	}
}
