package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	CreateBookmark(*Bookmark) error
	DeleteBookmark(int) error
	UpdateBookmark(*Bookmark) error
	GetBookmarkByID(int) (*Bookmark, error)
}

type SQLLiteStore struct {
	db *sql.DB
}

func NewSQLLiteStore() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(version)
}
