package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var dB *sql.DB;

func DbInit() error {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/klook")
	if err != nil {
		return err
	}

	dB = db

	return nil
}

func DbDestroy() {
	if dB != nil {
		dB.Close()
	}
}

func (user *User) selectUser() error {
	if dB == nil {
		return errors.Errorf("数据库连接没有建立")
	}

	return dB.QueryRow("select name, password, status from users where id = ?",
		user.Id).Scan(&user.Name, &user.Password, &user.Status)
}

func selectUsers() ([]User, error) {
	if dB == nil {
		return nil, errors.Errorf("数据库连接没有建立")
	}

	rows, err := dB.Query("select id, name, password, status from users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Status); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func insertUser(user User) error {
	if dB == nil {
		return errors.Errorf("数据库连接没有建立")
	}

	_, err := dB.Exec("insert into users(id, name, password, status) VALUES (?, ?, ?, ?)",
		user.Id, user.Name, user.Password, user.Status)

	return err
}

func updateUser(user User) error {
	_, err := dB.Exec("update users set name = ?, password = ?, status = ? where id = ?",
		user.Name, user.Password, user.Status, user.Id)

	return err
}

func deleteUser(id int) error {
	if dB == nil {
		return errors.Errorf("数据库连接没有建立")
	}

	_, err := dB.Exec("delete from users where id = ?", id)

	return err
}
