package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

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
	// Find .env file
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
	// fmt.Println(os.Getenv("CLIENT_ID"), " :---:", os.Getenv("CLIENT_SECRET"))
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
	if !strings.HasPrefix(title, "isrc") {
		return track, errors.New("malformed request (invalid title found)")
	}

	results, err := client.Search(ctx, title, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}
	// handle track results
	if results.Tracks != nil {
		fmt.Println("Tracks:")
		for _, item := range results.Tracks.Tracks {
			// fmt.Println("   item.ID:ID=> ", item.ID.String())
			// fmt.Println("   item.Name:ISRC=> ", item.ExternalIDs["isrc"])
			// fmt.Println("   item.Name:ImageURL=> ", item.ExternalURLs, item.Album.Images[0])
			// fmt.Println("   item.Name:Title=> ", item.Name)
			// fmt.Println("   item.Artists:=> ", item.Artists[0].Name, item.Artists[0].URI)
			// fmt.Printf("\n   item:=> %#v\n", item)
			track = Track{ID: item.ID.String(), ISRC: item.ExternalIDs["isrc"], Images: GetImageUrlOfTrack(item.Album.Images), Title: item.Name, Artists: GetArtistsOfTrack(item.Artists)}
			break
		}
	}
	if err = DB.Create(&track).Error; err != nil {
		log.Println("Inside getTrack:", err.Error())
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
		fmt.Println("Tracks:")
		for _, item := range results.Tracks.Tracks {
			// fmt.Println("   item.ID:ID=> ", item.ID.String())
			// fmt.Println("   item.Name:ISRC=> ", item.ExternalIDs["isrc"])
			// fmt.Println("   item.Name:ImageURL=> ", item.ExternalURLs, item.Album.Images[0].URL, GetImageUrlOfTrack(item.Album.Images))
			// fmt.Println("   item.Name:Title=> ", item.Name)
			// fmt.Println("   item.Artists:=> ", GetArtistsOfTrack(item.Artists))
			// fmt.Println("   item.Artists:=> ", item.Artists[0].Name, item.Artists[0].URI)
			track := Track{ID: item.ID.String(), ISRC: item.ExternalIDs["isrc"], Images: GetImageUrlOfTrack(item.Album.Images), Title: item.Name, Artists: GetArtistsOfTrack(item.Artists)}
			tracks = append(tracks, track)
		}
	}
	if err = DB.Create(&tracks).Error; err != nil {
		log.Println("Inside getTrack:", err.Error())
	}
	return tracks, nil
}

func GetArtistsOfTrack(simpleArtist []spotify.SimpleArtist) []Artist {
	artists := make([]Artist, 0)
	for _, item := range simpleArtist {
		artists = append(artists, Artist{ID: item.ID.String(), Name: item.Name, URI: string(item.URI)})
	}
	return artists
}

func GetImageUrlOfTrack(spotifyImages []spotify.Image) []Image {
	images := make([]Image, 0)
	for _, image := range spotifyImages {
		images = append(images, Image{Height: image.Height, Width: image.Width, URL: image.URL})
	}
	return images
}
