package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	album := Album{}

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&album); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, album)
	c.IndentedJSON(http.StatusCreated, album)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, alb := range albums {
		if alb.ID == id {
			c.IndentedJSON(http.StatusOK, alb)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/albums", getAlbums)
	r.POST("/albums", postAlbums)
	r.GET("/albums/:id", getAlbumByID)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
