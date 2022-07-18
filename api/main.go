package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kohei-kohei/go-redis/cache"
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

	if bread == nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	key := fmt.Sprintf("BreadID:%d", id)
	if err := cache.Set(c, key, bread.Name); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": bread.Name,
	})
}
