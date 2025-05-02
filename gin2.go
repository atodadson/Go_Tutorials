package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	// json tag to deserialize json body
	Name        string `json:"name"`
	Email       string `json:"email" binding:"required,email"`
	Age         int    `json:"age" binding:"gte=10,lte=100"`                    // Range 10 to 100
	FatherAge   int    `json:"fatherage" binding:"gte=10,lte=100,gtfield=Age"`  // Range 10 to 100 and greater than age
	Description string `json:"description" binding:"omitempty,max=255,min=10"`  // Checking min and max string length
	Status      string `json:"status" binding:"omitempty,oneof=married single"` // Response part of a list
	// binding helps with post data validation
	// required to ensure it's required and email to specify it's an email
}

/*
POPULAR VALIDATION TAGS
email, make sure the field is in email format (does not check the
validity of the email itself)

 email
 gte=10,lte=1000   ==>  When binding to an integer, validate the range of the value
 max=255               maximum length of a string
 min=18                minimum length of a string
 oneof=married single  one in a set of values (here, married or single)

time_format:"2006-01-02" Useful for defining dates
ltefield=OtherDate" time_format:"2006-01-02"  make sure the date comes before the other date, defined as
OtherDate in the same struct

startswith=MAC,len=9   make sure the string is of length 9 and starts with MAC
Uppercase / lowercase   make sure the field is uppercase or lowercase only
 alphanum / alpha      only accept english letters and numerals
 contains=key           make sure the string contains a key
 endswith=.             make sure the string ends with a period (.)

*/

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
				c.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()}) // .Error(), "status code": err.StatusOK
				// c.AbortWithError(http.StatusBadRequest, err)
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
