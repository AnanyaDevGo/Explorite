package middleware

import (
	"ExploriteGateway/pkg/helper"
	"ExploriteGateway/pkg/utils/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("authorization")
		//fmt.Println(tokenHeader, "this is the token header user")

		if tokenHeader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
			fmt.Println("autherisatiuon is working", tokenHeader)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenPart := splitted[1]
		tokenClaims, err := helper.ValidateTokenUser(tokenPart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		fmt.Println("id", tokenClaims.ID)
		c.Set("id", tokenClaims.ID)

		fmt.Println("claims id", tokenClaims.ID)
		c.Set("tokenClaims", tokenClaims)
		c.Next()
	}
}
