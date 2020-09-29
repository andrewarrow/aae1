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

func Custom404(c *gin.Context) {
	c.String(404, "we have no page for this url, yet.")
}

func main() {
	router := gin.Default()
	router.GET("/", WelcomeIndex)
	router.GET("/about-us", AboutUs)
	router.NoRoute(Custom404) //         <------- notice the special NoRotue func
	port := 3000
	router.Run(fmt.Sprintf(":%d", port))
}
