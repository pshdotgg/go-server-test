package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Running!!!",
		})
	})
	r.Run()
}

func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}