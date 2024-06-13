package handler

import (
	logging "ExploriteGateway/Logging"
	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/utils/models"
	"ExploriteGateway/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	GRPC_Client interfaces.AdminClient
}

func NewAdminHandler(adminClient interfaces.AdminClient) *AdminHandler {
	return &AdminHandler{
		GRPC_Client: adminClient,
	}
}

// AdminSignUp godoc
// @Summary Admin Sign Up
// @Description Create a new admin account
// @Tags Admin
// @Accept json
// @Produce json
// @Param adminDetails body models.AdminSignUp true "Admin Sign Up Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response
// @Router /admin/signup [post]
func (ad *AdminHandler) AdminSignUp(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()
	var adminDetails models.AdminSignUp

	fmt.Println("gateway", adminDetails.Email)

	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		logrusLogger.Error("Error in getting details", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminSignUp(adminDetails)
	if err != nil {
		logrusLogger.Error("Error in admin signup", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logrusLogger.Info("Admin successfully signedup")
	success := response.ClientResponse(http.StatusOK, "Admin created successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}

// LoginHandler godoc
// @Summary Admin Login
// @Description Authenticate an admin account
// @Tags Admin
// @Accept json
// @Produce json
// @Param adminDetails body models.AdminLogin true "Admin Login Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/login [post]
func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()
	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		logrusLogger.Error("Error in getting details", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminLogin(adminDetails)
	if err != nil {
		logrusLogger.Error("Error in admin login", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logrusLogger.Info("Admin loggedin succesfully")
	success := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}

// GetUsers godoc
// @Summary Get Users
// @Description Retrieve a paginated list of users
// @Tags Admin
// @Produce json
// @Param page query int true "Page number"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/list [get]
func (ad *AdminHandler) GetUsers(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		logrusLogger.Error("Error in parsing page number", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ad.GRPC_Client.GetUsers(page)
	if err != nil {
		logrusLogger.Error("Error in retrieving users", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Successfully retrieved the users")
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", users, nil)
	c.JSON(http.StatusOK, successRes)
}

// BlockUser godoc
// @Summary Block a User
// @Description Block a user by ID
// @Tags Admin
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/block [post]
func (ad *AdminHandler) BlockUser(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	id := c.Query("id")
	err := ad.GRPC_Client.BlockUser(id)
	if err != nil {
		logrusLogger.Error("Missing user ID in request")
		errorRes := response.ClientResponse(http.StatusBadRequest, "user could not be blocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Successfully blocked the users")
	successRes := response.ClientResponse(http.StatusOK, "Successfully blocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// UnBlockUser godoc
// @Summary Unblock a User
// @Description Unblock a user by ID
// @Tags Admin
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/unblock [post]
func (ad *AdminHandler) UnBlockUser(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	id := c.Query("id")
	fmt.Println("userid", id)
	err := ad.GRPC_Client.UnBlockUser(id)

	if err != nil {
		logrusLogger.Error("Missing user ID in request")
		errorRes := response.ClientResponse(http.StatusBadRequest, "user could not be unblocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Successfully unblocked the users")
	successRes := response.ClientResponse(http.StatusOK, "Successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
