package main
import (
    "strings"
    "encoding/json"
    "fmt"
    "database/sql"
    "log"
    "io/ioutil"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
type Person struct{
    Username    string `json:"username"`
	Firstname     string    `json:"first"`
	Lastname string `json:"last"`
    Email string `json:"email"`
    Password string `json:"password"`
    Confirmpassword string `json:"confirm"`
}
func signup(w http.ResponseWriter,r *http.Request){
    if(r.Method=="POST"){
        body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			defer r.Body.Close()
            // fmt.Println(string(body))

            var p Person
            err = json.Unmarshal(body, &p)
            if err != nil {
                panic(err)
            }
        
            fmt.Println(p.Username, p.Firstname, p.Lastname)
            if(p.Username==""){
                http.Error(w,"Invalid Username.",http.StatusBadRequest)
                return 
            }
            if(p.Firstname==""){
                http.Error(w,"Invalid Firstname.",http.StatusBadRequest)
                return 
            }
            if(p.Lastname==""){
                http.Error(w,"Invalid Lastname.",http.StatusBadRequest)
                return
            }
            if(p.Email==""){
                http.Error(w,"Invalid Email.",http.StatusBadRequest)
                return 
            }
            if(p.Password==""){
                http.Error(w,"Password must not be empty.",http.StatusBadRequest)
                return
            }
            if(p.Password!=p.Confirmpassword){
                http.Error(w,"Password doesn't match.",http.StatusBadRequest)
                return
            }
            stmt, err := db.Prepare("INSERT INTO sampletable(Username, Firstname, Lastname, Email, Password) VALUES (?, ?, ?, ?, ?)")
            if err != nil {
                log.Fatal(err)
            }
            defer stmt.Close()
            _, err = stmt.Exec(p.Username, p.Firstname, p.Lastname,p.Email,p.Password)
            if err != nil {
                words := strings.Split(err.Error(), " ")
                for i:=0;i<len(words);i++{
                    if(words[i]=="Duplicate" && words[i+1]=="entry"){
                        http.Error(w,"Username already exist",http.StatusBadRequest)
                    }
                }

            }

    }else{
        http.Error(w,"Invalid Request Method",http.StatusMethodNotAllowed)
    }

}
type Login struct{
    Username string `json:"name"`
    Password string `json:"password"`
}
func login(w http.ResponseWriter,r *http.Request){
    if(r.Method=="POST"){
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Fatal(err)
            return 
        }
        defer r.Body.Close()
        var L Login
        err = json.Unmarshal(body, &L)
        if err != nil {
            panic(err)
        }
        if (L.Username==" "){
            http.Error(w,"Enter Username",http.StatusBadRequest)
            return 
        }
        if (L.Password==" "){
            http.Error(w,"Enter Password",http.StatusBadRequest)
            return 
        }
        var dbUsername, dbPassword string
    err = db.QueryRow("SELECT username, password FROM sampletable WHERE username = ?", L.Username).Scan(&dbUsername, &dbPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w,"User Not Found",http.StatusBadRequest)
            return 
        }
    }
    if dbPassword == L.Password {
        w.WriteHeader(http.StatusOK)
        return 
    }
    http.Error(w,"Incorrect Password",http.StatusBadRequest)
        



    }else{
        http.Error(w,"Invalid Method Request",http.StatusMethodNotAllowed)
        return 
    }

}
type Password struct{
    Username string `json:"name"`
    Oldpassword string`json:"opass"`
    Newpassword string `json:"npass"`
}
func chanpassword(w http.ResponseWriter,r *http.Request){
    fmt.Println("varma")
    if(r.Method=="PUT"){
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Fatal(err)
            return 
        }
        defer r.Body.Close()
        var N Password
        err = json.Unmarshal(body, &N)
        if err != nil {
            panic(err)
        }
        if (N.Username==" "){
            http.Error(w,"Enter Username",http.StatusBadRequest)
            return 
        }
        if (N.Oldpassword==" "){
            http.Error(w,"Enter Password",http.StatusBadRequest)
            return 
        }
        var dbUsername, dbPassword string
        err = db.QueryRow("SELECT username, password FROM sampletable WHERE username = ?", N.Username).Scan(&dbUsername, &dbPassword)
        if err != nil {
            if err == sql.ErrNoRows {
                http.Error(w,"User Not Found",http.StatusBadRequest)
                return 
        }
    }
    if dbPassword == N.Oldpassword {
        _, err = db.Exec("UPDATE sampletable SET Password = ? WHERE Username = ?", N.Newpassword, N.Username)
    if err != nil {
       log.Fatal(err) 
    }

        return 
    }
    http.Error(w,"Incorrect Password",http.StatusBadRequest)
        


    }else{
        http.Error(w,"Invalid Method Request",http.StatusMethodNotAllowed)
    }
}
func main() {
    db, err = sql.Open("mysql", "root:Cadbury@2000@tcp(localhost:3306)/sampledb")
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Success!")
    defer db.Close()
    http.HandleFunc("/sign-up",signup)
    http.HandleFunc("/login",login)
    http.HandleFunc("/changepassword",chanpassword)
    log.Fatal(http.ListenAndServe(":8080", nil))

}