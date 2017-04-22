package main

import (
	"net/http"

	"github.com/robinle/etcdweb/operation"

	"gopkg.in/gin-gonic/gin.v1"
)

type Config struct {
	Endpoint string `form:"etcd-endpoint" binding:"required"`
}

var etcdClient = operation.EtcdClient{}

func main() {
	etcdClient.InitClient("http://192.168.14.166:32379")
	router := gin.Default()
	router.LoadHTMLGlob("ui/*")

	router.GET("/", index)
	router.POST("/setendpoint", setEndpoint)
	router.GET("/get/:key", getKey)
	router.GET("/get/:key/*subkey", getKey)
	router.GET("/web/:key", renderHTML)
	router.GET("/web/:key/*subkey", renderHTML)

	router.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{})
}

func setEndpoint(c *gin.Context) {
	config := Config{}
	if c.Bind(&config) == nil {
		etcdClient.InitClient(config.Endpoint)
	}
	c.HTML(http.StatusOK, "table.html", gin.H{"key": "//"})
}

func renderHTML(c *gin.Context) {
	key := c.Param("key")
	subkey := c.Param("subkey")
	key = key + subkey
	c.HTML(http.StatusOK, "table.html", gin.H{"key": key})
}

func getKey(c *gin.Context) {
	key := c.Param("key")
	subkey := c.Param("subkey")
	key = key + subkey
	keys, err := etcdClient.GetDirKeys(key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"400": "Not Found!"})
	}
	c.JSON(http.StatusOK, gin.H{"data": keys})
}
