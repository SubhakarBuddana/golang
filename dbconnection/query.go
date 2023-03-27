package dbconnection

import (
	"database/sql"
	"errors"
	"log"

	"github.com/SubhakarBuddana/golang/types"
)

func InsertUser(person types.Person) error {
	db, err := DBconnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO sampletable(Username, Firstname, Lastname, Email, Password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(person.Username, person.Firstname, person.Lastname, person.Email, person.Password)
	if err != nil {
		return errors.New("Username already exist")
	}
	return nil

}

func LoginUser(L types.Login) error {
	var dbUsername, dbPassword string
	db, err := DBconnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.QueryRow("SELECT username, password FROM sampletable WHERE username = ?", L.Username).Scan(&dbUsername, &dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("User not found")
		}
	}
	if dbPassword != L.Password {
		return errors.New("Incorrect Password")
	}

	return nil
}

func UserChangePassword(N types.ChangePasswordRequest) error {
	var dbUsername, dbPassword string
	db, err := DBconnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.QueryRow("SELECT username, password FROM sampletable WHERE username = ?", N.Username).Scan(&dbUsername, &dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("User not found")
		}
	}
	if dbPassword == N.Oldpassword {
		_, err = db.Exec("UPDATE sampletable SET Password = ? WHERE Username = ?", N.Newpassword, N.Username)
		if err != nil {
			return err
		}

	}
	if dbPassword != N.Oldpassword {
		return errors.New("Incorrect Password")
	}
	return nil
}
