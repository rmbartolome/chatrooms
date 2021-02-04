package main

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	store *dbStore
	db    *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	connString := "dbname=chatroom_temp sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	_, err := s.db.Query("DELETE FROM message")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateMessage() {
	d, _ := time.Parse("2006-02-01", "2020-05-04")

	s.store.CreateMessage(&Message{
		Username: "test username",
		Content:  "test content",
		Date:     d,
	})

	res, err := s.db.Query("SELECT count(*) FROM message WHERE username='test username'")
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		if err := res.Scan(&count); err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

func (s *StoreSuite) TestGetMessage() {
	_, err := s.db.Query("INSERT INTO message(username, content, date) VALUES('test', 'test content', '2020-05-05')")
	if err != nil {
		s.T().Fatal(err)
	}

	messages, err := s.store.GetMessages()
	if err != nil {
		s.T().Fatal(err)
	}

	nMessages := len(messages)
	if nMessages != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nMessages)
	}

	d, _ := time.Parse("2006-02-01", "2020-05-05")
	expectedMessage := Message{Username: "test", Content: "test content", Date: d, Id: messages[0].Id}
	messages[0].Date = d

	if *messages[0] != expectedMessage {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedMessage, *messages[0])
	}

}
