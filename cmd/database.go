package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Track struct {
	ID      uint     `gorm:"primaryKey"`
	TrackID string   `gorm:"column:track_id;uniqueIndex;not null;ON CONFLICT DO NOTHING" binding:"-" json:"track_id,omitempty"`
	ISRC    string   `gorm:"column:isrc;uniqueIndex;not null" json:"isrc"`
	Title   string   `gorm:"column:title;not null" json:"title,omitempty"`
	Images  []Image  `gorm:"column:images;many2many:track_images" json:"images,omitempty"`
	Artists []Artist `gorm:"column:artists;many2many:track_artists" json:"artists,omitempty"`
}

type Image struct {
	ID     uint   `gorm:"primaryKey"`
	Height int    `gorm:"column:height;not null" json:"height,omitempty"`
	Width  int    `gorm:"column:width;not null" json:"width,omitempty"`
	URL    string `gorm:"column:url;uniqueIndex;not null;ON CONFLICT DO NOTHING" json:"url,omitempty"`
}

type Artist struct {
	ID       uint   `gorm:"primaryKey"`
	ArtistID string `gorm:"column:artist_id;uniqueIndex;not null;ON CONFLICT DO NOTHING" binding:"-" json:"artist_id,omitempty"`
	Name     string `gorm:"column:name;not null" json:"name,omitempty"`
	URI      string `gorm:"column:uri;not null" json:"uri,omitempty"`
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
	fmt.Println("Database tables migrated...")
}
