package main

import (
	"./cors"
	"github.com/gin-gonic/gin"
	"github.com/ngerakines/ginpongo2"
	
)

func main() {
	r := gin.Default()
	r.Use(cors.Cors())
	r.Use(ginpongo2.Pongo2())

	r.GET("/", func(c *gin.Context) {
		c.Set("template", "index.htm")
		c.Set("data", map[string]interface{}{"name": "Hello World!"})
	})

	r.Run(":6060")
}
