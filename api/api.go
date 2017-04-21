package main

import (
	"net/http"

	"github.com/robinle/etcdoper/operation"

	"gopkg.in/gin-gonic/gin.v1"
)

var etcdClient = operation.EtcdClient{}

func main() {
	etcdClient.InitClient("http://192.168.14.166:32379")
	router := gin.Default()
	router.LoadHTMLGlob("ui/*")

	router.GET("/", setEndpoint)
	router.GET("/get/:key", getKey)
	router.GET("/get/:key/*subkey", getKey)
	router.GET("/web/:key", index)
	router.GET("/web/:key/*subkey", index)

	router.Run(":8080")
}

func setEndpoint(c *gin.Context) {
	etcdClient.InitClient("http://192.168.14.166:32379")
	c.JSON(http.StatusOK, gin.H{"200": "ok"})
}

func index(c *gin.Context) {
	key := c.Param("key")
	subkey := c.Param("subkey")
	key = key + subkey
	c.HTML(http.StatusOK, "index.html", gin.H{"key": key})
}

func getKey(c *gin.Context) {
	key := c.Param("key")
	subkey := c.Param("subkey")
	key = key + subkey
	keys, err := etcdClient.GetDirKeys(key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"400": "Not Found!"})
	}

	// c.String(http.StatusOK, key)
	c.JSON(http.StatusOK, gin.H{"data": keys})
}
