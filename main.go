package main

import (
	"github.com/arogyaGurkha/GurkhaFabricAPI/docs"
	lc "github.com/arogyaGurkha/GurkhaFabricAPI/lifecycle"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"
	"os/exec"
	"strings"
)

// peerVersion represents version information of peer
type peerVersion struct {
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
}

// albums represents data about a record album
type album struct {
	ID     string  `json:"ID"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @title Hyperledger Fabric Gurkhaman API
// @description API to run fabric binaries
// @version 0.1
// @contact.name arogya.Gurkha
// @contact.url https://github.com/arogyaGurkha
func main() {

	updateSwagger()

	router := gin.Default()

	// Routes
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:ID", getAlbumByID)
	router.GET("/fabric/peer/", getPeerVersion)
	router.POST("/fabric/lifecycle/admin/:organization", lc.SetAdmin)
	router.GET("/fabric/lifecycle/admin", lc.GetAdmin)

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // Points to the API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("ID")

	// Loop over the list of albums, looking for an album whose ID value matches
	// the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getPeerVersion checks the current version of peer binary
// @Summary Get the current peer binary version
// @Description `peer version` is executed through `exec.Command()` to return the current peer version.
// @Produce json
// @Tags peer
// @Success 200 {object} peerVersion "successful operation"
// @Router /fabric/peer/ [get]
func getPeerVersion(c *gin.Context) {
	var versionResponse peerVersion

	cmd := exec.Command("peer", "version")
	output, _ := cmd.Output()

	outputList := strings.Split(string(output), "\n")
	version := strings.SplitAfter(outputList[1], ":")[1][1:]      // "Version: 2.4.0" -> "2.4.0"
	architecture := strings.SplitAfter(outputList[4], ":")[1][1:] // "OS/Arch: darwin/amd64" -> "darwin/amd64"

	versionResponse.Version = version
	versionResponse.Architecture = architecture

	c.IndentedJSON(http.StatusOK, versionResponse)
}

func updateSwagger() {
	cmd := exec.Command("swag", "init", "-g", "main.go", "--output", "docs")
	cmd.Run()
}
