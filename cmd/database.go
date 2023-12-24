package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Track struct {
	ID      string   `gorm:"primary_key;column:track_id" binding:"-" json:"track_id,omitempty"`
	ISRC    string   `gorm:"column:isrc;uniqueIndex;not null" json:"isrc"`
	Title   string   `gorm:"column:title" json:"title,omitempty"`
	Images  []Image  `gorm:"column:images;foreignKey:url;references:track_id" json:"images,omitempty"`
	Artists []Artist `gorm:"column:artists;foreignKey:artist_id;references:track_id" json:"artists,omitempty"`
}

type Image struct {
	Height int    `gorm:"column:height" json:"height,omitempty"`
	Width  int    `gorm:"column:width" json:"width,omitempty"`
	URL    string `gorm:"primary_key;column:url" json:"url,omitempty"`
}

type Artist struct {
	ID   string `gorm:"primary_key;column:artist_id" binding:"-" json:"artist_id,omitempty"`
	Name string `gorm:"column:name" json:"name,omitempty"`
	URI  string `gorm:"column:uri" json:"uri,omitempty"`
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

	DB.Exec("PRAGMA foreign_keys = ON")
	DB.AutoMigrate(&Image{}, &Artist{}, &Track{})
}
