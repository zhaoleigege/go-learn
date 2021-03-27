package main

type User struct {
	Id       int    `json:"id"`
	Status   int    `json:"status"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
