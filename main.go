package main

import "github.com/gin-gonic/gin"

func IndexHanlder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHanlder)
	router.Run()
}
