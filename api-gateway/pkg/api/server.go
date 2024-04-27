package server

import (
	"ExploriteGateway/pkg/api/handler"
	"ExploriteGateway/pkg/api/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, userHandler *handler.UserHandler) *ServerHTTP {

	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignUp)

	router.POST("/user/signup", userHandler.UserSignUp)
	router.POST("/user/login", userHandler.UserLogin)

	router.Use(middleware.AdminAuthMiddleware())
	{
		usermanagement := router.Group("/user")
		{
			usermanagement.GET("list", adminHandler.GetUsers)
			usermanagement.PATCH("/block", adminHandler.BlockUser)
			usermanagement.PATCH("/unblock", adminHandler.UnBlockUser)
		}

	}

	return &ServerHTTP{engine: router}

}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :8000")
	err := s.engine.Run(":8000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
