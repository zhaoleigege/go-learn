package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"reflect"
)

var scheme = `
CREATE TABLE person(
    id int auto_increment primary key,
    name varchar(20) not null,
    age int not null 
);
`

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var persons []Person
	fmt.Println("slice为nil:", reflect.ValueOf(&persons).IsNil())
}

func createTable(db *sqlx.DB) (int64, error) {
	result := db.MustExec(scheme)

	affect, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affect, nil
}

func insertTable(db *sqlx.DB) (int64, error) {
	result := db.MustExec(`INSERT INTO person (name, age) VALUES (?, ?)`, "test1", 22)

	affect, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affect, nil
}

func getTable(db *sqlx.DB) (*Person, error) {
	person := &Person{}

	err := db.Get(person, "SELECT * FROM person where id = ?", 1)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func selectTable(db *sqlx.DB) ([]Person, error) {
	var persons []Person // 初始化slice为nil，但是&slice不是nil

	err := db.Select(&persons, "SELECT * FROM person")
	if err != nil {
		return nil, err
	}

	return persons, nil
}
