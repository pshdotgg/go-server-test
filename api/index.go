package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func root(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Running")
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	root(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}