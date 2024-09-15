package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ghenoo/microservices-go/cmd/db"
)

// @Summary Lista todos os usuários
// @Description Retorna uma lista de usuários
// @Tags users
// @Produce  json
// @Success 200 {array} User
// @Router /users [get]

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        users, err := db.GetUsers()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(users)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// @Summary Adiciona um novo usuário
// @Description Adiciona um usuário com base no payload enviado
// @Tags users
// @Accept  json
// @Produce  json
// @Success 201 {object} User
// @Router /users/add [post]

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var user db.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        err := db.InsertUser(user.Name, user.Email)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
