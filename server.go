package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleDefault(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", "World")

}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", handleDefault)

	r.Run(":8080")
}
