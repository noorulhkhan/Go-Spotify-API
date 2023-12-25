package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	_ "golang.org/x/oauth2/clientcredentials"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	"gorm.io/gorm/clause"
)

var (
	ctx    context.Context
	client *spotify.Client
)

func authorize() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	prepare()
}

func getToken() (*oauth2.Token, error) {
	ctx = context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}
	return token, err
}

func prepare() {
	token, _ := getToken()
	httpClient := spotifyauth.New().Client(ctx, token)
	client = spotify.New(httpClient)
}

// ...
func FetchTrackByTitle(title string) (Track, error) {
	track := Track{}
	trackLight := Track{}

	if title == "" {
		return track, errors.New("malformed request (no title found)")
	}
	// if !strings.HasPrefix(title, "isrc") {
	// 	return track, errors.New("malformed request (invalid title found)")
	// }

	results, err := client.Search(ctx, title, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}
	// handle track results
	if results.Tracks != nil {
		fmt.Println("A Track by ISRC code:")
		popularity := 0
		for _, item := range results.Tracks.Tracks {
			//max popularity
			popularity = max(popularity, item.Popularity)
		}
		for _, item := range results.Tracks.Tracks {
			if item.Popularity == popularity {
				track = Track{TrackID: item.ID.String(), ISRC: item.ExternalIDs["isrc"], Images: GetImageUrlOfTrack(item.Album.Images, item.ID.String()), Title: item.Name, Artists: GetArtistsOfTrack(item.Artists, item.ID.String())}
				trackLight = Track{TrackID: item.ID.String(), ISRC: item.ExternalIDs["isrc"], Images: nil, Title: item.Name, Artists: nil}

				// todo: insert records to database
				tx := DB.Begin()
				if err = tx.Save(&trackLight).Error; err != nil {
					tx.Rollback()
					log.Println("Inside getTrack:", err.Error())
				}
				if err = tx.Save(GetImageUrlOfTrack(item.Album.Images, track.TrackID)).Error; err != nil {
					tx.Rollback()
					log.Println("Inside getTrack:", err.Error())
				}
				if err = tx.Save(GetArtistsOfTrack(item.Artists, track.TrackID)).Error; err != nil {
					tx.Rollback()
					log.Println("Inside getTrack:", err.Error())
				}
				if err := tx.Commit().Error; err != nil {
					tx.Rollback()
				}
				// if err := DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&track).Error; err != nil {
				// 	fmt.Println("DEBUG: Something wrong went")
				// }
			}
		}
	}
	return track, nil
}

func FetchTracksByArtist(artist string) ([]Track, error) {
	tracks := make([]Track, 0)
	tracksLight := make([]Track, 0)
	images := make([]Image, 0)
	artists := make([]Artist, 0)

	if artist == "" {
		return tracks, errors.New("malformed request (no artist found)")
	}
	results, err := client.Search(ctx, artist, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}
	// handle track results
	if results.Tracks != nil {
		log.Println("Tracks by Artist:")
		for _, item := range results.Tracks.Tracks {
			track := Track{TrackID: item.ID.String(), ISRC: item.ExternalIDs["isrc"], Images: GetImageUrlOfTrack(item.Album.Images, item.ID.String()), Title: item.Name, Artists: GetArtistsOfTrack(item.Artists, item.ID.String())}
			tracks = append(tracks, track)
			trackLight := Track{TrackID: item.ID.String(), ISRC: item.ExternalIDs["isrc"], Images: nil, Title: item.Name, Artists: nil}
			tracksLight = append(tracksLight, trackLight)
			for _, image := range *GetImageUrlOfTrack(item.Album.Images, item.ID.String()) {
				images = append(images, Image{Height: image.Height, Width: image.Width, URL: image.URL, TrackID: item.ID.String()})
			}
			for _, artist := range *GetArtistsOfTrack(item.Artists, item.ID.String()) {
				artists = append(artists, Artist{ID: artist.ID, Name: artist.Name, URI: artist.URI, TrackID: item.ID.String()})
			}
		}
	}
	// todo: insert records to database
	tx := DB.Begin()
	if err = tx.Save(&tracksLight).Error; err != nil {
		tx.Rollback()
		log.Println("Inside getTracks:", err.Error())
	}
	if err = tx.Save(&images).Error; err != nil {
		tx.Rollback()
		log.Println("Inside getTracks:", err.Error())
	}
	if err = tx.Save(&artists).Error; err != nil {
		tx.Rollback()
		log.Println("Inside getTracks:", err.Error())
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
	}

	// if err = DB.Debug().Save(&tracks).Error; err != nil {
	// if err := DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&tracksLight).Error; err != nil {
	// 	log.Println("Inside getTrack:", err.Error())
	// }
	return tracks, nil
}

func FetchTracks() ([]Track, error) {
	var tracks []Track
	// if err = DB.Preload("Image").Preload("Artist").Find(&tracks).Error; err != nil {
	// err = DB.Joins("images").Joins("artists").Find(&tracks).Error
	// err = DB.Debug().Preload("images").Preload("artists").Find(&tracks).Error
	err = DB.Debug().Preload(clause.Associations).Find(&tracks).Error
	if err != nil {
		log.Println("Inside getTracks:", err.Error())
	}
	return tracks, nil
}

func GetArtistsOfTrack(simpleArtist []spotify.SimpleArtist, trackid string) *[]Artist {
	artists := make([]Artist, 0)
	for _, item := range simpleArtist {
		artists = append(artists, Artist{ID: item.ID.String(), Name: item.Name, URI: string(item.URI), TrackID: trackid})
	}
	return &artists
}

func GetImageUrlOfTrack(spotifyImages []spotify.Image, trackid string) *[]Image {
	images := make([]Image, 0)
	for _, image := range spotifyImages {
		images = append(images, Image{Height: image.Height, Width: image.Width, URL: image.URL, TrackID: trackid})
	}
	return &images
}
