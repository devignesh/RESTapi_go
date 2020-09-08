package main

import (
	_ "RESTapi_go/AgiN/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("test")

	r.GET("/test", handler.pingget)

	r.Run()
}
