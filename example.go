package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":10086") // listen and serve on 0.0.0.0:8080
	// router.Run(":3000") for a hard coded port
}
