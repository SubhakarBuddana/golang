package main
import(
	"net/http"
	"log"
)
func Main(){
	http.HandleFunc("/sign-up",signup)
    http.HandleFunc("/login",login)
    http.HandleFunc("/changepassword",chanpassword)
    log.Fatal(http.ListenAndServe(":8080", nil))
}