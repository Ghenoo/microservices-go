package handlers

import "net/http"

func SetupRoutes() {
    http.HandleFunc("/users", GetUsersHandler)
    http.HandleFunc("/users/add", PostUserHandler)
}
