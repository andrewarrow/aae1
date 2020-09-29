package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeIndex(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func AboutUs(c *gin.Context) {
	c.String(http.StatusOK, "we are a new website")
}

func main() {
	router := gin.Default()
	router.GET("/", WelcomeIndex)
	router.GET("/about-us", AboutUs) // < - notice the new route
	port := 3000
	router.Run(fmt.Sprintf(":%d", port))
}
