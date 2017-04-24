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
	router.GET("/raw/:key", getKey)
	router.GET("/raw/:key/*subkey", getKey)
	router.DELETE("/raw/:key", deleteKey)
	router.DELETE("/raw/:key/*subkey", deleteKey)
	router.GET("/web/:key", renderData)
	router.GET("/web/:key/*subkey", renderData)

	router.Run(":8080")
}

// return etcd config page
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "config.html", gin.H{})
}

// config etcd and return the root key
func setEtcdServer(c *gin.Context) {
	config := Config{}
	if c.Bind(&config) == nil {
		etcdClient.InitClient(config.Endpoint)
	}
	c.HTML(http.StatusOK, "table.html", gin.H{"key": "//"})
}

// render data table of etcd
func renderData(c *gin.Context) {
	key := c.Param("key")
	subkey := c.Param("subkey")
	key = key + subkey
	c.HTML(http.StatusOK, "table.html", gin.H{"key": key})
}

// get etcd data
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

// delete etcd data
func deleteKey(c *gin.Context) {
	key := c.Param("key")
	subkey := c.Param("subkey")
	key = key + subkey
	err := etcdClient.Delete(key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"delete": "failed"})
	}
	c.JSON(http.StatusOK, gin.H{})
}
