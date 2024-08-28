package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Crazy Train", Artist: "Ozzy Osbourne", Price: 56.99},
    {ID: "2", Title: "Peru", Artist: "Ed Sheeran", Price: 17.99},
    {ID: "3", Title: "Somebody to Love", Artist: "Queen", Price: 39.99},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8000")
}
