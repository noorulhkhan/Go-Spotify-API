package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Track struct {
	ID      string   `json:"track_id,omitempty" gorm:"primary_key;column:track_id" binding:"-"`
	ISRC    string   `json:"isrc" gorm:"column:irsc;uniqueIndex;not null"`
	Title   string   `json:"title,omitempty" gorm:"column:title"`
	Images  []Image  `json:"image,omitempty"`
	Artists []Artist `json:"artists,omitempty"`
}

type Image struct {
	Hight int    `json:"height,omitempty" gorm:"column:high"`
	Width int    `json:"width,omitempty" gorm:"column:width"`
	URL   string `json:"url,omitempty" gorm:"column:url"`
}

type Artist struct {
	ID   string `json:"artist_id,omitempty" gorm:"primary_key;column_artist_id" binding:"-"`
	Name string `json:"name,omitempty" gorm:"column:name"`
	URI  string `json:"uri,omitempty" gorm:"column:uri"`
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

	DB.AutoMigrate(&Image{}, &Artist{}, &Track{})
}
