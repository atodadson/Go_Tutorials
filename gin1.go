package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// // Gin Route without parameter
// func router() *gin.Engine {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(200, "hello")
// 	})
// 	return r
// }

// // Gin Route with Parameter
// func router() *gin.Engine {
// 	r := gin.Default()
// 	r.GET("/:name", func(c *gin.Context) {
// 		user := c.Param("name")
// 		c.String(200, fmt.Sprintf("hello, %s", user))
// 	})
// 	return r
// }

// // Grouping Routes together
// func router() *gin.Engine {
// 	r := gin.Default()
// 	userRoute := r.Group("/user")
// 	{
// 		userRoute.GET("/hello/:name", func(c *gin.Context) {
// 			user := c.Param("name")
// 			c.String(200, fmt.Sprintf("hello, %s", user))
// 		})
// 	}
// 	return r
// }

type Message struct {
	// Json tag to de-serialize json body
	Name string `json:"name"`
}

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
