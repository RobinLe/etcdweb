package main

import (
	"net/http"

	"github.com/robinle/etcdoper/operation"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("ui/*")
	router.GET("/", index)

	router.GET("/:key", getKey)
	// router.POST("/:key", getKey)
	// router.PUT("/:key", getKey)
	// router.DELETE("/:key", getKey)

	router.Run(":8080")
}

// Index index

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
	// c.JSON(http.StatusOK, gin.H{})
}

func getKey(c *gin.Context) {
	key := c.Param("key")
	keys, err := operation.GetDirKeys(key)
	if err != nil {
		keyValue, err := operation.GetKeyValue(key)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
		}
		c.JSON(http.StatusOK, keyValue)
	}
	c.JSON(http.StatusOK, gin.H{"data": keys})
}
