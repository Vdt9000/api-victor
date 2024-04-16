package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Nascimento string `json:"date_of_birth"`
}

var users []User

func main() {
	router := httprouter.New()
	router.GET("/users", listUsers)
	router.POST("/users/create", createUser)
	router.DELETE("/users/delete/:id", deleteUser)
	router.PUT("/users/edit/:id", editUser)
	http.ListenAndServe(":8080", router)

	//http.HandleFunc("/users", listUsers)
	//http.HandleFunc("/users/create", createUser)
	//http.HandleFunc("/users/edit/:id", editUser)
	//http.HandleFunc("/users/delete/:id", deleteUser)

	//fmt.Println("API is running on :8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func readUsers(id int) ([]User, error) {
	var users []User
	file, err := os.ReadFile("file.json")
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(file, &users)
	if err != nil {
		return users, err
	}

	fmt.Print(id)

	return users, nil

}

func listUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := readUsers(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := readUsers(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1

	users = append(users, newUser)

	err = WriteUsers(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func editUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "ID do usuário não fornecido", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID informado é diferente do esperado", http.StatusBadRequest)
		return
	}
	users, err := readUsers(1)
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

	// 3. Atualiza os campos do usuário encontrado com os novos dados

	for i, user := range users {
		if user.ID == userID {
			users[i].Name = updateUser.Name
			users[i].Nascimento = updateUser.Nascimento
			err = WriteUsers(users)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	// 4. Retorna erro se o usuário não for encontrado
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Implementa exclusão de usuário aqui
	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "ID do usuário não fornecido", http.StatusBadRequest)
		return
	}
	// 2. Pesquisa o usuário na lista de usuários pelo ID
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID do usuário inválido", http.StatusBadRequest)
		return
	}

	users, err := readUsers(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			err = WriteUsers(users)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// 4. Retorna erro se o usuário não for encontrado
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func WriteUsers(users []User) error {
	jsonUser, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = os.WriteFile("file.json", jsonUser, 0644)
	if err != nil {
		return err
	}

	return nil

}
