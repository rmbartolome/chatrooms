package main

import (
	"database/sql"
	"fmt"
)

type Store interface {
	CreateMessage(message *Message) error
	GetMessages() ([]*Message, error)
	UpdateMessage(message *Message) error
	DeleteMessage(message *Message) error
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateMessage(message *Message) error {
	_, err := store.db.Query("INSERT INTO message(username, content, date) VALUES($1, $2, $3)", message.Username, message.Content, message.Date)
	return err
}

func (store *dbStore) GetMessages() ([]*Message, error) {
	messagees := []*Message{}

	rows, err := store.db.Query("SELECT id, username, content, date FROM message")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		message := &Message{}

		if err := rows.Scan(&message.Id, &message.Username, &message.Content, &message.Date); err != nil {
			return nil, err
		}

		messagees = append(messagees, message)
	}

	return messagees, nil
}

func (store *dbStore) UpdateMessage(message *Message) error {
	fmt.Println(message)
	_, err := store.db.Query("UPDATE message SET username = $1, content=$2, date=$3 WHERE id=$4", message.Username, message.Content, message.Date, message.Id)
	return err
}

func (store *dbStore) DeleteMessage(message *Message) error {
	_, err := store.db.Query("DELETE FROM message WHERE id=$1", message.Id)
	return err
}

var store Store

func InitStore(s Store) {
	store = s
}
