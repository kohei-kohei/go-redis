package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kohei-kohei/go-redis/db"
)

func main() {
	r := gin.Default()
	r.GET("/bread", getBread)
	r.Run()
}

func getBread(c *gin.Context) {
	var message string

	bread, err := db.GetBreadName()
	if err != nil {
		message = err.Error()
	} else {
		message = bread.Name
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}
