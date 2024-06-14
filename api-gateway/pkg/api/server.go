package server

import (
	"ExploriteGateway/pkg/api/handler"
	"ExploriteGateway/pkg/api/middleware"
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, userHandler *handler.UserHandler, postHandler *handler.PostHandler, chatHandler *handler.ChatHandler, videocallHandler *handler.VideoCallHandler) *ServerHTTP {
	router := gin.New()
	router.Use(gin.Logger())
	router.Static("/static", "./static")
	router.LoadHTMLGlob("template/*")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	

	router.GET("/exit", videocallHandler.ExitPage)
	router.GET("/error", videocallHandler.ErrorPage)
	router.GET("/index", videocallHandler.IndexedPage)

	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignUp)

	router.POST("/user/signup", userHandler.UserSignUp)
	router.POST("/user/login", userHandler.UserLogin)
	router.POST("/user/otplogin", userHandler.UserOTPLogin)
	router.POST("/user/otpverify", userHandler.VerifyOTP)

	router.Use(middleware.UserAuthMiddleware())

	userprofile := router.Group("/user/profile")
	{
		userprofile.POST("/add", userHandler.AddProfile)
		userprofile.GET("/get", userHandler.GetProfile)
		userprofile.PATCH("/edit", userHandler.EditProfile)
		userprofile.PATCH("/change-password", userHandler.ChangePassword)
	}
	post := router.Group("/user/post")
	{
		post.POST("/add", postHandler.AddPost)
		post.GET("/list", postHandler.ListPost)
		post.PATCH("/edit", postHandler.EditPost)
		post.DELETE("/delete", postHandler.DeletePost)
		post.PATCH("/upvote", postHandler.UpvotePost)
		post.PATCH("/downvote", postHandler.DownvotePost)
	}
	chat := router.Group("/user/chat")
	{
		chat.GET("", chatHandler.FriendMessage)
		chat.GET("/message", chatHandler.GetChat)
	}

	router.Use(middleware.AdminAuthMiddleware())

	usermanagement := router.Group("/admin")
	{
		usermanagement.GET("/list", adminHandler.GetUsers)
		usermanagement.PATCH("/block", adminHandler.BlockUser)
		usermanagement.PATCH("/unblock", adminHandler.UnBlockUser)
	}

	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :8000")
	err := s.engine.Run(":8000")
	if err != nil {
		log.Printf("error while starting the server: %v", err)
	}
}
