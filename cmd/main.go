package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mateushenriquedasilva/caching-proxy-go/api"
)

func init() {
	api.InitRedis()
	r := gin.Default()

	r.GET("/products", api.Middleware, api.GetProducts)

	r.Run()
}

func main() {}
