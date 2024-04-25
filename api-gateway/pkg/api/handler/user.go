package handler

import (
	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/utils/models"
	"ExploriteGateway/pkg/utils/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	GRPC_Client interfaces.UserClient
}

func NewUserHandler(userClient interfaces.UserClient) *UserHandler {
	return &UserHandler{
		GRPC_Client: userClient,
	}
}

func (uh *UserHandler) UserSignUp(c *gin.Context) {
	var userDetails models.UserSignup

	fmt.Println("gateway", userDetails.Email)

	if err := c.ShouldBindJSON(&userDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := uh.GRPC_Client.UserSignUp(userDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User created successfully", user, nil)
	c.JSON(http.StatusOK, success)
}

func (uh *UserHandler) UserLogin(c *gin.Context) {
	var userDetails models.UserLogin
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := uh.GRPC_Client.UserLogin(userDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User authenticated successfully", user, nil)
	c.JSON(http.StatusOK, success)
}
