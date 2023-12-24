package main

import (
	"net/http"

	"ltimindtree/cmd/utils"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	router := gin.Default()
	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/track/search/:title", getTrack)
	router.GET("/track/find/:artist", findTracks)

	router.Run("localhost:8080")
}

// getTrack godoc
// @Summary Gets a track by Title or ISRC
// @Description Gets a track by Title or ISRC
// @Accept  json
// @Produce  json
// @Param title path string true "Title"
// @Success 200 {object} main.Track
// @Failure 400 {object} utils.ErrResp
// @Failure 404 {object} utils.ErrResp
// @Failure 500 {object} utils.ErrResp
// @Router /track/search/{title} [get]
func getTrack(c *gin.Context) {
	var track Track
	var err error
	title := c.Param("title")
	track, err = FetchTrackByTitle(title)
	if err != nil {
		errresp := utils.ErrResp{Error: err.Error(), Message: "Unable to fetch data"}
		c.JSON(http.StatusBadRequest, errresp)
		return
	}
	c.IndentedJSON(http.StatusOK, track)
}

// getPassenger godoc
// @Summary Get tracks collection by artist
// @Description Get tracks collection by artist
// @Accept  json
// @Produce  json
// @Param artist path string true "artist"
// @Success 200 {object} []main.Track
// @Failure 400 {object} utils.ErrResp
// @Failure 404 {object} utils.ErrResp
// @Failure 500 {object} utils.ErrResp
// @Router /track/find/{phrase} [get]
func findTracks(c *gin.Context) {
	var err error
	var tracks []Track
	artist := c.Param("artist")
	tracks, err = FetchTracksByArtist(artist)
	if err != nil {
		errresp := utils.ErrResp{Error: err.Error(), Message: "Unable to fetch data"}
		c.JSON(http.StatusBadRequest, errresp)
		return
	}
	c.IndentedJSON(http.StatusOK, tracks)
}
