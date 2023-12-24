package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Track struct {
	ID      string   `json:"track_id,omitempty"`
	ISRC    string   `json:"isrc" gorm:"uniqueIndex,not null"`
	Title   string   `json:"title,omitempty"`
	Images  []Image  `json:"image,omitempty"`
	Artists []Artist `json:"artists,omitempty"`
}

type Image struct {
	Hight int    `json:"height,omitempty"`
	Width int    `json:"width,omitempty"`
	URL   string `json:"url,omitempty"`
}

type Artist struct {
	ID   string `json:"artist_id,omitempty"`
	Name string `json:"name,omitempty"`
	URI  string `json:"uri,omitempty"`
}

var DB *gorm.DB
var err error

func InitialMigration() {
	flag := false
	dbfile := "database.db"
	if _, err := os.Stat(dbfile); err == nil {
		// flag = true
		os.Remove(dbfile)
		fmt.Println("Old", dbfile, "deleted.")
	}
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to open the SQLite database.")
	}
	if flag {
		return
	}

	// DB.AutoMigrate(&Image{}, &Artist{}, &Track{})
}
