Database:
	sqlite3 is used instead of MySQL
	3 tables created - tracks, images, artists

Code & build:
	; latest code is on dev branch
	git clone https://github.com/noorulhkhan/Go-Spotify-API.git
	cd Go-Spotify-API
	git checkout dev
	cd cmd
	go build
	.\cmd.exe

Swagger:
	http://localhost:8080/docs/index.html
	Parameters for,
		SearchTrackByTitle:	holiday,	isrc:GBAYE0601477,	isrc:USWB11403680,	isrc:USEE10001992
		findTracksByArtist:	The Beatles,	Laufey,			Mariah Carey

POSTMAN:
	GET	http://localhost:8080/track/search/holiday
	GET	http://localhost:8080/track/search/isrc:GBAYE0601477
	GET	http://localhost:8080/track/find/The Beatles
	GET	http://localhost:8080/track

