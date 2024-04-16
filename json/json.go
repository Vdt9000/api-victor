package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Nascimento string `json:"date_of_birth"`
}

func main() {
	data, err := os.ReadFile("file.json")
	if err != nil {
		panic(err)
	}

	var user []User
	err = json.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
