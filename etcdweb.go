package main

import (
	"net/http"

	"github.com/robinle/etcdweb/operation"

	"gopkg.in/gin-gonic/gin.v1"
)

// Config etcd config
type Config struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Endpoint string `form:"etcd-endpoint" binding:"required"`
}

var etcdClient = operation.EtcdClient{}

func main() {
	etcdClient.InitClient("http://127.0.0.1:2379")
	router := gin.Default()
	router.LoadHTMLGlob("ui/html/*")

	router.GET("/", index)
	router.POST("/setetcd", setEtcdServer)
	router.GET("/get/:key", getKey)
	router.GET("/get/:key/*subkey", getKey)
	router.GET("/web/:key", renderData)
	router.GET("/web/:key/*subkey", renderData)

	router.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "config.html", gin.H{})
}

func setEtcdServer(c *gin.Context) {
	config := Config{}
	if c.Bind(&config) == nil {
		etcdClient.InitClient(config.Endpoint)
	}
	c.HTML(http.StatusOK, "table.html", gin.H{"key": "//"})
}

func renderData(c *gin.Context) {
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
