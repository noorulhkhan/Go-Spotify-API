package main

import (
	"net/http"

	utils "ltimindtree/utils"

	_ "ltimindtree/cmd/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRouter() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/track/search/:title", getTrack)
	router.GET("/track/find/:artist", findTracks)
	router.GET("/track", getTracks)

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

// findTracks godoc
// @Summary Get tracks collection by Artist
// @Description Get tracks collection by Artist
// @Accept  json
// @Produce  json
// @Param artist path string true "Artist"
// @Success 200 {object} []main.Track
// @Failure 400 {object} utils.ErrResp
// @Failure 404 {object} utils.ErrResp
// @Failure 500 {object} utils.ErrResp
// @Router /track/find/{artist} [get]
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

// getTrack godoc
// @Summary Gets tracks collection
// @Description Gets tracks collection
// @Accept  json
// @Produce  json
// @Success 200 {object} []main.Track
// @Failure 400 {object} utils.ErrResp
// @Failure 404 {object} utils.ErrResp
// @Failure 500 {object} utils.ErrResp
// @Router /track [get]
func getTracks(c *gin.Context) {
	var tracks []Track
	var err error
	tracks, err = FetchTracks()
	if err != nil {
		errresp := utils.ErrResp{Error: err.Error(), Message: "Unable to fetch data"}
		c.JSON(http.StatusBadRequest, errresp)
		return
	}
	c.IndentedJSON(http.StatusOK, tracks)
}
