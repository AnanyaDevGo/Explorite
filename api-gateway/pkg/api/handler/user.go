package handler

import (
	logging "ExploriteGateway/Logging"
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

// UserSignUp godoc
// @Summary User Sign Up
// @Description Create a new user account
// @Tags User
// @Accept json
// @Produce json
// @Param userDetails body models.UserSignup true "User Sign Up Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user/signup [post]
func (uh *UserHandler) UserSignUp(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()
	var userDetails models.UserSignup

	fmt.Println("gateway", userDetails.Email)

	if err := c.ShouldBindJSON(&userDetails); err != nil {
		logrusLogger.Error("Error in getting details", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := uh.GRPC_Client.UserSignUp(userDetails)
	if err != nil {
		fmt.Println(err)
		logrusLogger.Error("Error in user signup", err)
		anerr := errors.New("error in services")
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, anerr.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logrusLogger.Info("User successfully signedup")
	success := response.ClientResponse(http.StatusOK, "User created successfully", user, nil)
	c.JSON(http.StatusOK, success)
}

// UserLogin godoc
// @Summary User Login
// @Description Authenticate a user
// @Tags User
// @Accept json
// @Produce json
// @Param userDetails body models.UserLogin true "User Login Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user/login [post]
func (uh *UserHandler) UserLogin(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()
	var userDetails models.UserLogin
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := uh.GRPC_Client.UserLogin(userDetails)
	if err != nil {
		logrusLogger.Error("Error in user login", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logrusLogger.Info("User loggedin succesfully")
	success := response.ClientResponse(http.StatusOK, "User authenticated successfully", user, nil)
	c.JSON(http.StatusOK, success)
}

// UserOTPLogin godoc
// @Summary User OTP Login
// @Description Authenticate a user
// @Tags User
// @Accept json
// @Produce json
// @Param userDetails body models.UserOTPLogin true "User Login Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user/otplogin [post]

func (uh *UserHandler) UserOTPLogin(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	var userDetails models.UserOTPLogin
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	otp, err := uh.GRPC_Client.UserOTPLogin(userDetails.Email)
	if err != nil {
		logrusLogger.Error("Error in generating OTP", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to generate OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logrusLogger.Info("OTP generated successfully")
	success := response.ClientResponse(http.StatusOK, "OTP generated successfully", map[string]string{"OTP": otp}, nil)
	c.JSON(http.StatusOK, success)
}

// VerifyOTP godoc
// @Summary Verify OTP
// @Description Verify the OTP sent to the user's email
// @Tags User
// @Accept json
// @Produce json
// @Param otpVerification body models.OtpVerification true "OTP Verification Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/otpverify [post]
func (ot *UserHandler) VerifyOTP(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	var code models.OtpVerification
	if err := c.BindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	verified, err := ot.GRPC_Client.OtpVerification(code.Email, code.Otp)
	if err != nil {
		logrusLogger.Error("Error in verifying OTP", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("OTP verified successfully")
	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", verified, nil)
	c.JSON(http.StatusOK, successRes)
}

// AddProfile godoc
// @Summary Add User Profile
// @Description Add a new profile for the user
// @Tags User
// @Accept json
// @Produce json
// @Param profile body models.UserProfile true "User Profile Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/profile/add [post]
func (uh *UserHandler) AddProfile(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	id, _ := c.Get("id")
	var profile models.UserProfile
	if err := c.ShouldBindJSON(&profile); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := uh.GRPC_Client.AddProfile(id.(int), profile); err != nil {
		logrusLogger.Error("Error in adding profile", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add profile", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Profile added successfully")
	successRes := response.ClientResponse(http.StatusOK, "Successfully added profile", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetProfile godoc
// @Summary Get User Profile
// @Description Get the profile of the user
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/profile/get [get]
func (u *UserHandler) GetProfile(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	addresses, err := u.GRPC_Client.GetProfile(id)
	if err != nil {
		logrusLogger.Error("Error in retrieving profile", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Profile retrieved successfully")
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", addresses, nil)
	c.JSON(http.StatusOK, successRes)
}

// EditProfile godoc
// @Summary Edit User Profile
// @Description Edit the profile details of the user
// @Tags User
// @Accept json
// @Produce json
// @Param editProfile body models.EditProfile true "Edit Profile Details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/profile/edit [patch]
func (u *UserHandler) EditProfile(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	var model models.EditProfile
	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err := validator.New().Struct(model)
	if err != nil {
		err = errors.New("missing constraints for email id")
		errorRes := response.ClientResponse(http.StatusBadRequest, "Email id is not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = u.GRPC_Client.EditProfile(id, model)
	if err != nil {
		logrusLogger.Error("Error in editing profile", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Error updating the values", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Profile edited successfully")
	successRes := response.ClientResponse(http.StatusCreated, "Details edited successfully", nil, nil)
	c.JSON(http.StatusCreated, successRes)
}

// ChangePassword godoc
// @Summary Change User Password
// @Description Change the password of the user
// @Tags User
// @Accept json
// @Produce json
// @Param changePassword body models.ChangePassword true "Change Password Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/profile/change-password [patch]
func (u *UserHandler) ChangePassword(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	var ChangePassword models.ChangePassword
	if err := c.BindJSON(&ChangePassword); err != nil {
		logrusLogger.Error("Error  in getting details", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := u.GRPC_Client.ChangePassword(id, ChangePassword.Oldpassword, ChangePassword.Password, ChangePassword.Repassword); err != nil {
		logrusLogger.Error("Error in changing password", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("password changed succesfully")
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
