package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/test/:name", WebRoot)

	engine.Run(":9205")
}

// WebRoot ...
func WebRoot(c *gin.Context) {
	str := c.Param("name")

	result := Permutations(str)
	c.JSON(http.StatusOK, result)
	// context.String(http.StatusOK, "hello, world %v", "HI")
}

// Permutations ...
func Permutations(str string) interface{} {
	tmp := []string{}
	for _, s := range str {
		tmp = append(tmp, string(s))
	}

	for i := 0; i < len(tmp); i++ {
		
	}

	return tmp
}
