package handlers

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/ghenoo/microservices-go/cmd/api/docs"
	"github.com/ghenoo/microservices-go/cmd/producer"
	httpSwagger "github.com/swaggo/http-swagger"
)

var KafkaWriter *producer.KafkaWriter

func ListEndpointsHandler(w http.ResponseWriter, r *http.Request) {
    endpoints := []string{
        "GET /users",
        "POST /users/add",
        "POST /send-message",
        "GET /status",
    }
    for _, endpoint := range endpoints {
        fmt.Fprintln(w, endpoint)
    }
}

func SetupRoutes() {
    http.HandleFunc("/endpoints", ListEndpointsHandler) 
    http.HandleFunc("/users", GetUsersHandler)
    http.HandleFunc("/users/add", PostUserHandler)
    http.HandleFunc("/send-message", SendMessageHandler)
    http.HandleFunc("/status", StatusHandler)

	http.HandleFunc("/swagger/*any", httpSwagger.WrapHandler)
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := "Mensagem enviada para o Kafka!"
	err := KafkaWriter.WriteMessages(context.Background(),
		producer.NewMessage([]byte("Key"), []byte(message)),
	)
	if err != nil {
		http.Error(w, "Failed to send message to Kafka", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, message)
}
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Servidor está funcionando!")
}
