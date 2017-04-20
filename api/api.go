package main

import (
	"net/http"

	"github.com/robinle/etcdoper/operation"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("ui/*")

	router.GET("/get/:key", getKey)
	router.GET("/get/:key/*subkey", getKey)
	router.GET("/web/:key", index)

	router.Run(":8080")
}

// Index index
func index(c *gin.Context) {
	key := c.Param("key")
	c.HTML(http.StatusOK, "index.html", gin.H{"key": key})
}

func getKey(c *gin.Context) {
	key := c.Param("key")
	subkey := c.Param("subkey")
	key = key + subkey
	keys, err := operation.GetDirKeys(key)
	if err != nil {
		keyValue, err := operation.GetKeyValue(key)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
		}
		c.JSON(http.StatusOK, keyValue)
	}

	// c.String(http.StatusOK, key)
	c.JSON(http.StatusOK, gin.H{"data": keys})
}
