package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Username string    `json:"username"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
	Id       int       `json:"id"`
}

func createMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	message.Date = time.Now()

	err = store.CreateMessage(&message)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/list/", http.StatusFound)
}

func getMessageHandler(w http.ResponseWriter, r *http.Request) {
	messagees, err := store.GetMessages()

	messageListBytes, err := json.Marshal(messagees)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(messageListBytes)
}

func updateMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{}
	err := json.NewDecoder(r.Body).Decode(&message)

	err = store.UpdateMessage(&message)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/list/", http.StatusFound)
}

func deleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{}
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = store.DeleteMessage(&message)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/list/", http.StatusFound)
}
