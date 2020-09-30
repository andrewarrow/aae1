package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeIndex(c *gin.Context) {
	c.Data(http.StatusOK, "text/html;",
		[]byte("<html><body><h1 style='font-size:72px;color:red;'>Hello</h1></body></html>")) // <-- real HTML
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
	router.NoRoute(Custom404)
	port := 3000
	router.Run(fmt.Sprintf(":%d", port))
}
