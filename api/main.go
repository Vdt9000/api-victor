package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/createusers", getUsers)
	http.HandleFunc("/deleteusers", getUsers)
	http.HandleFunc("/updateusers", getUsers)
	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	func (createUser(w http.http.ResponseWriter, r *http.Request)){



	}
		}
	
		
	

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			ID:   1,
			Name: "Rafael",
		},

		{
			ID:   2,
			Name: "Pipoca",
		}})
}
