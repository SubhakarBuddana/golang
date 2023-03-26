package dbconnection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBconnection() (*sql.DB, error) {
	const (
		db1 = "mysql"
		db2 = "root:Cadbury@2000@tcp(localhost:3306)/sampledb"
	)

	db, err := sql.Open(db1, db2)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success!")
	return db, nil
}
