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
	bread, err := db.GetBreadName()
	if err != nil {
		c.String(500, "Server Error")
		return
	}

	c.JSON(200, gin.H{
		"message": bread.Name,
	})
}
