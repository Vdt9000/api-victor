package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Nascimento string `json:"date_of_birth"`
}

var users = []User{
	{ID: 1, Name: "Victor", Nascimento: "18-09-2004"},
	{ID: 2, Name: "Nick", Nascimento: "20-01-2000"},
	{ID: 3, Name: "Vitinho", Nascimento: "30-03-1999"},
	{ID: 4, Name: "Vini", Nascimento: "25-07-1995"},
	{ID: 5, Name: "Luan", Nascimento: "25-07-1995"},
}

func main() {
	http.HandleFunc("/users", listUsers)
	http.HandleFunc("/users/create", createUser)
	http.HandleFunc("/users/edit", editUser)
	http.HandleFunc("/users/delete", deleteUser)

	fmt.Println("API is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1

	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	// Implementar edição de usuário aqui
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID do usuário não fornecido", http.StatusBadRequest)
		return
	}

	// 2. Pesquisar o usuário na lista de usuários pelo ID
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID do usuário inválido", http.StatusBadRequest)
		return
	}

	var updateUser User
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Atualizar os campos do usuário encontrado com os novos dados
	for i, user := range users {
		if user.ID == userID {
			users[i] = updateUser
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	// 4. Retornar erro se o usuário não for encontrado
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Implementar exclusão de usuário aqui
	// 1. Extrair o ID do usuário a ser excluído da URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID do usuário não fornecido", http.StatusBadRequest)
		return
	}

	// 2. Pesquisar o usuário na lista de usuários pelo ID
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID do usuário inválido", http.StatusBadRequest)
		return
	}

	// 3. Remover o usuário da lista de usuários, se encontrado
	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// 4. Retornar erro se o usuário não for encontrado
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}
