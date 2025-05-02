package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	// json tag to deserialize json body
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
	// binding helps with post data validation
	// required to ensure it's required and email to specify it's an email
}

// Binding custom types to the body of post request
func router() *gin.Engine {
	r := gin.Default()
	userRoute := r.Group("/user")
	{
		userRoute.GET("/hello/:name", func(c *gin.Context) {
			user := c.Param("name")
			response := fmt.Sprintf("hello, %s", user)
			c.String(http.StatusOK, response)
		})
		userRoute.POST("/post", func(c *gin.Context) {
			body := Message{}
			if err := c.BindJSON(&body); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			fmt.Println(body)
			c.JSON(http.StatusAccepted, &body)
		})
	}
	return r
}

func main() {
	router().Run()
}
