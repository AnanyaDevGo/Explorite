package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoCallHandler struct{}

func NewVideoCallHandler() *VideoCallHandler {
	return &VideoCallHandler{}
}

func (v *VideoCallHandler) ExitPage(c *gin.Context) {
	c.HTML(http.StatusOK, "exit.html", nil)
}

func (v *VideoCallHandler) ErrorPage(c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", nil)
}

func (v *VideoCallHandler) IndexedPage(c *gin.Context) {
	call := c.DefaultQuery("call", "")

	c.HTML(http.StatusOK, "index.html", gin.H{"call": call})
}
