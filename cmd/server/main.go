package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	chatrooms "github.com/rbartolome/chatrooms/pkg"
	"github.com/rbartolome/chatrooms/pkg/kafka"
)

type Request struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/messages", createMessageHandler).Methods("POST")
	r.HandleFunc("/messages", getMessageHandler).Methods("GET")
	r.HandleFunc("/messages", updateMessageHandler).Methods("PUT")
	r.HandleFunc("/messages", deleteMessageHandler).Methods("DELETE")

	return r
}

func main() {
	connString := "dbname=chatroom sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	var (
		brokers = os.Getenv("KAFKA_BROKERS")
		topic   = os.Getenv("KAFKA_TOPIC")
	)

	publisher := kafka.NewPublisher(strings.Split(brokers, ","), topic)

	r := gin.Default()
	// nr := newRouter()

	r.POST("/publish", publishHandler(publisher))
	r.POST("/join", joinHandler(publisher))

	// fmt.Println("Servidor corriendo por el puerto 8080")
	// http.ListenAndServe(":8080", nr)

	_ = r.Run()
}

func joinHandler(publisher chatrooms.Publisher) func(*gin.Context) {
	return func(c *gin.Context) {
		var req Request
		err := json.NewDecoder(c.Request.Body).Decode(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		message := chatrooms.NewSystemMessage(fmt.Sprintf("%s has joined the room!", req.Username))

		if err := publisher.Publish(context.Background(), message); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "message published"})
	}
}

func publishHandler(publisher chatrooms.Publisher) func(*gin.Context) {
	return func(c *gin.Context) {
		var req Request
		err := json.NewDecoder(c.Request.Body).Decode(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		message := chatrooms.NewMessage(req.Username, req.Message)

		msg := Message{Username: req.Username, Content: req.Message}
		err = store.CreateMessage(&msg)
		if err != nil {
			fmt.Println(err)
		}

		if err := publisher.Publish(context.Background(), message); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "message published"})
	}
}
