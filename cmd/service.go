package main

import (
	"context"
	"encoding/json"
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
				track = Track{ID: item.ID.String(), ISRC: item.ExternalIDs["isrc"], Images: GetImageUrlOfTrack(item.Album.Images), Title: item.Name, Artists: GetArtistsOfTrack(item.Artists)}

				// todo: insert records to database
				if err := DB.Save(&track).Error; err != nil {
					log.Println("Inside getTrack:", err.Error())
				}
			}
		}
	}
	return track, nil
}

func FetchTracksByArtist(artist string) ([]Track, error) {
	tracks := make([]Track, 0)

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
			track := Track{ID: item.ID.String(), ISRC: item.ExternalIDs["isrc"], Images: GetImageUrlOfTrack(item.Album.Images), Title: item.Name, Artists: GetArtistsOfTrack(item.Artists)}
			tracks = append(tracks, track)
		}
	}
	// todo: insert records to database
	if err = DB.Save(&tracks).Error; err != nil {
		log.Println("Inside getTrack:", err.Error())
	}
	return tracks, nil
}

func GetArtistsOfTrack(simpleArtist []spotify.SimpleArtist) string {
	var str string
	// if len(simpleArtist) != 0 {
	// 	item := simpleArtist[0]
	// 	artist := Artist{ID: item.ID.String(), Name: item.Name, URI: string(item.URI)}
	// 	bytes, _ := json.Marshal(artist)
	// 	str = fmt.Sprint(string(bytes))
	// }
	artists := make([]Artist, 0)
	for _, item := range simpleArtist {
		artist := Artist{ID: item.ID.String(), Name: item.Name, URI: string(item.URI)}
		artists = append(artists, artist)
	}
	bytes, _ := json.Marshal(artists)
	str = fmt.Sprint(string(bytes))
	return str
}

func GetImageUrlOfTrack(spotifyImages []spotify.Image) string {
	var str string
	// if len(spotifyImages) != 0 {
	// 	item := spotifyImages[0]
	// 	image := Image{Height: item.Height, Width: item.Width, URL: item.URL}
	// 	bytes, _ := json.Marshal(image)
	// 	str = fmt.Sprint(string(bytes))
	// }
	images := make([]Image, 0)
	for _, item := range spotifyImages {
		image := Image{Height: item.Height, Width: item.Width, URL: item.URL}
		images = append(images, image)
	}
	bytes, _ := json.Marshal(images)
	str = fmt.Sprint(string(bytes))
	return str
}
