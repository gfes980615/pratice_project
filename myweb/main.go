package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/permutations/:words", WebRoot)

	engine.Run(":9205")
}

// WebRoot ...
func WebRoot(c *gin.Context) {
	words := c.Param("words")

	result := Permutations(words)
	c.JSON(http.StatusOK, result)
}
