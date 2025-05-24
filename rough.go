package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "os"
)

type Address struct {
    Country string `json:"country"`
    City string `json:"city"`
    Street string `json:"street"`
    HouseNumber string `json:"house_number"`
}

type User struct {
    Name string    `json:"name"`
    Age int `json:"age"`
    Address Address `json:"address"`
}

func router() *gin.Engine {
    r := gin.Default()
    r.GET("/:name&mne", func(c *gin.Context) {
        user := c.Param("name&mne")
        designation := c.Param("mne")
        c.String(http.StatusOK, fmt.Sprintf("Hello %s, you are an %s", user, designation))
    })

    // Post request for uploading files
    r.POST("/login", func(c *gin.Context) {
        user := User{}
        if err := c.BindJSON(&user); err != nil {
            c.AbortWithError(http.StatusBadRequest, err)
            return
        }
        fmt.Println(user)
        c.JSON(http.StatusAccepted, &user)
    })
    r.POST("/upload", func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            fmt.Printf("An error occurred: %s\n", err)
            c.AbortWithError(http.StatusBadRequest, err)
            return
        }
        fmt.Println(file.Filename)

        upload_dir := "./uploads"

        os.MkdirAll(upload_dir, os.ModePerm)
        c.SaveUploadedFile(file, upload_dir)
        c.String(http.StatusAccepted, fmt.Sprintf("File %s saved successfully", file.Filename))
    } )
    return r
}

func main() {
    router().Run()
}