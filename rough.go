package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type Message struct {
    Name string    `json:"name"`
}

func router() *gin.Engine {
    r := gin.Default()
    r.GET("/:name&mne", func(c *gin.Context) {
        user := c.Param("name&mne")
        designation := c.Param("mne")
        c.String(http.StatusOK, fmt.Sprintf("Hello %s, you are an %s", user, designation))
    })
    return r
}

func main() {
    router().Run()
}