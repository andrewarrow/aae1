package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" // <--- 1. gin is major golang open source repo
)

func WelcomeIndex(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func main() {
	router := gin.Default()
	router.GET("/", WelcomeIndex) // <-- 2. We define where request to
	//                                      http://yourdomain:3000/ go
	port := 3000
	router.Run(fmt.Sprintf(":%d", port))
}
