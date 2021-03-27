package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getDb() *DB {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/klook")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetUser(id int) User {
	db := getDb()
	rows, err := db.Query("select * from users where id = ?", id)

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &password, &status)
		if err != nil {
			log.Fatal(err)
		}

		user := new(User)
		user.Id = id
		user.Status = status
		user.Name = name
		user.Password = password
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		id       int
		name     string
		password string
		status   int
	)

	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/klook")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from users")

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &password, &status)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, password, status)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
