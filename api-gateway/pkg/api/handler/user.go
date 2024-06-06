package handler

import (
	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/utils/models"
	"ExploriteGateway/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		anerr := errors.New("error in services")
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, anerr.Error())
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

func (uh *UserHandler) UserOTPLogin(c *gin.Context) {
	var userDetails models.UserOTPLogin
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	otp, err := uh.GRPC_Client.UserOTPLogin(userDetails.Email)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to generate OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "OTP generated successfully", map[string]string{"OTP": otp}, nil)
	c.JSON(http.StatusOK, success)
}
func (ot *UserHandler) VerifyOTP(c *gin.Context) {
	var code models.OtpVerification
	if err := c.BindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	verified, err := ot.GRPC_Client.OtpVerification(code.Email, code.Otp)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", verified, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uh *UserHandler) AddProfile(c *gin.Context) {
	id, _ := c.Get("id")
	var profile models.UserProfile
	if err := c.ShouldBindJSON(&profile); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	fmt.Println("idddd", id)
	if err := uh.GRPC_Client.AddProfile(id.(int), profile); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add profile", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added profile", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (u *UserHandler) GetProfile(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	fmt.Println("iddd", id)
	addresses, err := u.GRPC_Client.GetProfile(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", addresses, nil)
	c.JSON(http.StatusOK, successRes)
}
func (u *UserHandler) EditProfile(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	var model models.EditProfile
	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err := validator.New().Struct(model)
	if err != nil {
		err = errors.New("missing constraints for email id")
		errRes := response.ClientResponse(http.StatusBadRequest, "email id is not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = u.GRPC_Client.EditProfile(id, model)
	fmt.Println("errrorr", err)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error updating the values", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "details edited succesfully", nil, nil)

	c.JSON(http.StatusCreated, successRes)
}
func (u *UserHandler) ChangePassword(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	var ChangePassword models.ChangePassword
	if err := c.BindJSON(&ChangePassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := u.GRPC_Client.ChangePassword(id, ChangePassword.Oldpassword, ChangePassword.Password, ChangePassword.Repassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "password changed Successfully ", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// func (u *UserHandler) FollowReq(c *gin.Context) {
// 	id, exists := c.Get("id")
// 	if !exists {
// 		errorRes := response.ClientResponse(http.StatusUnauthorized, "User ID not found", nil, "User ID is required")
// 		c.JSON(http.StatusUnauthorized, errorRes)
// 		return
// 	}

// 	useridStr := c.Param("userid")
// 	userid, err := strconv.Atoi(useridStr)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid user ID", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	err = u.GRPC_Client.SendFollowReq(id.(int), userid)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Error in follow request", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Successfully followed user", nil, nil)
// 	c.JSON(http.StatusOK, successRes)
// }

// func (uh *UserHandler) AcceptFollowreq(c *gin.Context) {
// 	id, exists := c.Get("id")
// 	if !exists {
// 		errorRes := response.ClientResponse(http.StatusUnauthorized, "User ID not found", nil, "User ID is required")
// 		c.JSON(http.StatusUnauthorized, errorRes)
// 		return
// 	}

// 	useridStr := c.Param("userid")
// 	userid, err := strconv.Atoi(useridStr)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid user ID", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	err = uh.GRPC_Client.AcceptFollowRequest(id.(int), userid)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Error in accepting follow request", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Successfully accepted follow request", nil, nil)
// 	c.JSON(http.StatusOK, successRes)
// }

// func (uh *UserHandler) Followers(c *gin.Context) {
// 	id, exists := c.Get("id")
// 	if !exists {
// 		errorRes := response.ClientResponse(http.StatusUnauthorized, "User ID not found", nil, "User ID is required")
// 		c.JSON(http.StatusUnauthorized, errorRes)
// 		return
// 	}

// 	followers, err := uh.GRPC_Client.GetFollowers(id.(int))
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Error retrieving followers", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved followers", followers, nil)
// 	c.JSON(http.StatusOK, successRes)
// }

// func (uh *UserHandler) Followings(c *gin.Context) {
// 	id, exists := c.Get("id")
// 	if !exists {
// 		errorRes := response.ClientResponse(http.StatusUnauthorized, "User ID not found", nil, "User ID is required")
// 		c.JSON(http.StatusUnauthorized, errorRes)
// 		return
// 	}

// 	followings, err := uh.GRPC_Client.GetFollowings(id.(int))
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Error retrieving followings", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved followings", followings, nil)
// 	c.JSON(http.StatusOK, successRes)
// }
