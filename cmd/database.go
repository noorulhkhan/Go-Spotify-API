package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Track struct {
	// gorm.Model
	ID      uint     `gorm:"primaryKey"`
	TrackID string   `gorm:"column:track_id;uniqueIndex;not null;ON CONFLICT DO NOTHING" binding:"-" json:"track_id,omitempty"`
	ISRC    string   `gorm:"column:isrc;uniqueIndex;not null" json:"isrc"`
	Title   string   `gorm:"column:title;not null" json:"title,omitempty"`
	Images  []Image  `gorm:"column:images;many2many:track_images" json:"images,omitempty"`
	Artists []Artist `gorm:"column:artists;many2many:track_artists" json:"artists,omitempty"`
	// CreatedAt time.Time `gorm:"coulmn:created_at;autoCreateTime" json:"-"`
	// UpdatedAt time.Time `gorm:"coulmn:created_at;autoUpdateTime" json:"-"`
	// DeletedAt time.Time `gorm:"coulmn:created_at;index" json:"-"`
}

type Image struct {
	// gorm.Model
	ID     uint   `gorm:"primaryKey"`
	Height int    `gorm:"column:height;not null" json:"height,omitempty"`
	Width  int    `gorm:"column:width;not null" json:"width,omitempty"`
	URL    string `gorm:"column:url;uniqueIndex;not null;ON CONFLICT DO NOTHING" json:"url,omitempty"`
	// CreatedAt time.Time `gorm:"coulmn:created_at;autoCreateTime" json:"-"`
	// UpdatedAt time.Time `gorm:"coulmn:created_at;autoUpdateTime" json:"-"`
	// DeletedAt time.Time `gorm:"coulmn:created_at;index" json:"-"`
}

type Artist struct {
	// gorm.Model
	ID       uint   `gorm:"primaryKey"`
	ArtistID string `gorm:"column:artist_id;uniqueIndex;not null;ON CONFLICT DO NOTHING" binding:"-" json:"artist_id,omitempty"`
	Name     string `gorm:"column:name;not null" json:"name,omitempty"`
	URI      string `gorm:"column:uri;not null" json:"uri,omitempty"`
	// CreatedAt time.Time `gorm:"coulmn:created_at;autoCreateTime" json:"-"`
	// UpdatedAt time.Time `gorm:"coulmn:created_at;autoUpdateTime" json:"-"`
	// DeletedAt time.Time `gorm:"coulmn:created_at;index" json:"-"`
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
	// DB.Exec("PRAGMA table_info('tracks')")
	// DB.Exec("PRAGMA table_info('images')")
	// DB.Exec("PRAGMA table_info('artists')")
	fmt.Println("Database tables migrated...")
}
