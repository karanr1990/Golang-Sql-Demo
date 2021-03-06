package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "golangdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
UPDATE users
SET first_name = $2, last_name = $3
WHERE id = $1;`

	_, err = db.Exec(sqlStatement, 15, "NewFirst", "NewLast")
	if err != nil {
		panic(err)
	}
	fmt.Println("record updated")
}
