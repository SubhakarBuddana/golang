package main

import (
	"log"
	"net/http"

	"github.com/SubhakarBuddana/golang/userservice"
)

func main() {

	http.HandleFunc("/sign-up", userservice.Signup)
	http.HandleFunc("/login", userservice.Login)
	http.HandleFunc("/changepassword", userservice.Chanpassword)
	http.HandleFunc("/changepassword", userservice.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
