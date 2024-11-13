package main

import "github.com/gin-gonic/gin"

type Recipe struct {
	Name         string   `json:"name"`
	Tags         []string `json:"tags"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt  string   `json:"publishedAt"`
}

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func IndexNameParamHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}

func main() {
	router := gin.Default()
	//router.GET("/", IndexHandler)
	router.GET("/:name", IndexNameParamHandler)
	router.Run()
}
