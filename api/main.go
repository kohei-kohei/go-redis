package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/kohei-kohei/go-redis/cache"
	"github.com/kohei-kohei/go-redis/db"
	"github.com/kohei-kohei/go-redis/domain"
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

	key := fmt.Sprintf("BreadID:%d", id)
	value, err := cache.Get(c, key)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	if value != "" {
		b := domain.Bread{}
		json.Unmarshal([]byte(value), &b)
		c.JSON(http.StatusOK, gin.H{
			"bread": value,
		})
		return
	}

	time.Sleep(time.Second * 3)

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

	serialized, _ := json.Marshal(bread)
	if err := cache.Set(c, key, string(serialized)); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"bread": bread.Name,
	})
}
