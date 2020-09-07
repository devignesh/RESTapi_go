package main

import (
	_ "RESTapi_go/AgiN/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", handler.pingget)

	r.Run()
}
