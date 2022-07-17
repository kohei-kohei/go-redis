package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kohei-kohei/go-redis/db"
)

func main() {
	r := gin.Default()
	r.GET("/bread/:id", getBreadName)
	r.Run()
}

func getBreadName(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	bread, err := db.GetBread(id)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": bread.Name,
	})
}
