package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

var names = make([]string, 100)

func main() {
	r := gin.Default()
	r.GET("/:no", func(c *gin.Context) {
		no, err := strconv.Atoi(c.Param("no"))
		if err != nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": names[no],
		})
	})
	r.GET("/:no/@:name", func(c *gin.Context) {
		no, err := strconv.Atoi(c.Param("no"))
		if err != nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		names[no] = c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"name": names[no],
		})
	})
	r.Run()
}
