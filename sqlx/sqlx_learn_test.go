package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"log"
	"testing"
)

var db *sqlx.DB

func init() {
	var err error

	db, err = sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/sqlx")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateTable(t *testing.T) {
	_, err := createTable(db)
	if err != nil {
		log.Fatal(err)
	}
}

func TestInsertTable(t *testing.T) {
	_, err := insertTable(db)
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetTable(t *testing.T) {
	person, err := getTable(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("得到数据%+v\n", person)
}

func TestSelectTable(t *testing.T) {
	persons, err := selectTable(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, person := range persons {
		log.Printf("得到数据%+v\n", person)
	}
}
