package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var userName string
	connStr := "user=mydbuser dbname=mydbname port=5432 password=mydbpass sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	userID := 41
	sqlStatement := "SELECT name FROM users WHERE id = $1 limit 1"
  
	row := db.QueryRow(sqlStatement, userID)
	row.Scan(&userName)
	fmt.Println("Here is the fetched username:", userName)

	fmt.Println("Here is end of file")
}
