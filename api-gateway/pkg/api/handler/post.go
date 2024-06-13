package handler

import (
	logging "ExploriteGateway/Logging"
	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/utils/models"
	"ExploriteGateway/pkg/utils/response"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	GRPC_Client interfaces.PostClient
}

func NewPostHandler(postClient interfaces.PostClient) *PostHandler {
	return &PostHandler{
		GRPC_Client: postClient,
	}
}

// AddPost godoc
// @Summary Add a new post
// @Description Create a new post with caption and media
// @Tags Post
// @Accept multipart/form-data
// @Produce json
// @Param caption formData string true "Post Caption"
// @Param mediaurl formData string true "Media URL"
// @Param image formData file true "Post Image"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/post/add [post]
func (ph *PostHandler) AddPost(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	userID, _ := c.Get("id")
	var post models.AddPost
	post.Caption = c.PostForm("caption")
	post.MediaURL = c.PostForm("mediaurl")

	userId, ok := userID.(int)
	if !ok {
		err := errors.New("user id conversion failed")
		logrusLogger.Error(err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	post.UserId = strconv.Itoa(userId)

	file, err := c.FormFile("image")
	if err != nil {
		logrusLogger.Error("Error in getting data", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Error in getting data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	fileContent, err := file.Open()
	if err != nil {
		logrusLogger.Error("Error in opening file", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Error in opening file", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	defer fileContent.Close()

	imageBytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		logrusLogger.Error("Error in reading data", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Error in reading data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	post.Media = imageBytes

	if err := ph.GRPC_Client.AddPost(post, file); err != nil {
		logrusLogger.Error("Could not add post", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add post", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Successfully added post")
	successRes := response.ClientResponse(http.StatusOK, "Successfully added post", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// ListPost godoc
// @Summary List all posts
// @Description Get a list of all posts
// @Tags Post
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/post/list [get]
func (ph *PostHandler) ListPost(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	posts, err := ph.GRPC_Client.ListPost()
	if err != nil {
		logrusLogger.Error("Posts cannot be displayed", err)
		errRes := response.ClientResponse(http.StatusBadRequest, "Posts cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	logrusLogger.Info("Successfully listed posts")
	successRes := response.ClientResponse(http.StatusOK, "Successfully listed posts", posts, nil)
	c.JSON(http.StatusOK, successRes)
}

// EditPost godoc
// @Summary Edit a post
// @Description Edit the caption of an existing post
// @Tags Post
// @Accept multipart/form-data
// @Produce json
// @Param postid formData string true "Post ID"
// @Param caption formData string true "Post Caption"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/post/edit [patch]
func (ph *PostHandler) EditPost(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	userID, _ := c.Get("id")
	var editPost models.EditPost
	editPost.Caption = c.PostForm("caption")
	editPost.PostId = c.PostForm("postid")

	userId, ok := userID.(int)
	if !ok {
		err := errors.New("user id conversion failed")
		logrusLogger.Error(err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	editPost.UserId = strconv.Itoa(userId)

	if err := ph.GRPC_Client.EditPost(userId, editPost); err != nil {
		logrusLogger.Error("Could not edit post", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not edit post", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Successfully edited post")
	successRes := response.ClientResponse(http.StatusOK, "Successfully edited post", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// DeletePost godoc
// @Summary Delete a post
// @Description Delete a post by ID
// @Tags Post
// @Produce json
// @Param postid path string true "Post ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/post/{postid} [delete]
func (ph *PostHandler) DeletePost(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	postID := c.Param("postid")
	if postID == "" {
		err := errors.New("post ID is empty")
		logrusLogger.Error(err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		logrusLogger.Error("Invalid post ID", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := ph.GRPC_Client.DeletePost(postIDInt); err != nil {
		logrusLogger.Error("Failed to delete post", err)
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to delete post", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	logrusLogger.Info("Post deleted successfully")
	successRes := response.ClientResponse(http.StatusOK, "Post deleted successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// func (ph *PostHandler) SavePost(c *gin.Context) {
//     postID := c.Param("postid")
//     postIDInt, err := strconv.Atoi(postID)
//     if err != nil {
//         errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, err.Error())
//         c.JSON(http.StatusBadRequest, errorRes)
//         return
//     }

//     if err := ph.GRPC_Client.SavePost(postIDInt); err != nil {
//         errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to save post", nil, err.Error())
//         c.JSON(http.StatusInternalServerError, errorRes)
//         return
//     }

//     successRes := response.ClientResponse(http.StatusOK, "Post saved successfully", nil, nil)
//     c.JSON(http.StatusOK, successRes)
// }

// func (ph *PostHandler) UnSavePost(c *gin.Context) {
//     postID := c.Param("postid")
//     postIDInt, err := strconv.Atoi(postID)
//     if err != nil {
//         errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, err.Error())
//         c.JSON(http.StatusBadRequest, errorRes)
//         return
//     }

//     if err := u.GRPC_Client.UnSavePost(postIDInt); err != nil {
//         errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to unsave post", nil, err.Error())
//         c.JSON(http.StatusInternalServerError, errorRes)
//         return
//     }

//	    successRes := response.ClientResponse(http.StatusOK, "Post unsaved successfully", nil, nil)
//	    c.JSON(http.StatusOK, successRes)
//	}

// UpvotePost godoc
// @Summary Upvote a Post
// @Description Upvote a post by ID
// @Tags Post
// @Produce json
// @Param postid query string true "Post ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /post/upvote [post]
func (ph *PostHandler) UpvotePost(c *gin.Context) {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	userIDInterface, exists := c.Get("id")
	if !exists {
		err := errors.New("failed to get user ID from context")
		logrusLogger.Error(err)
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to get user ID from context", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		err := errors.New("failed to convert userID to int")
		logrusLogger.Error(err)
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to convert userID to int", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	postID := c.Query("postid")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		logrusLogger.Error("Invalid post ID", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := ph.GRPC_Client.UpvotePost(userID, postIDInt); err != nil {
		logrusLogger.Error("Failed to upvote post", err)
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to upvote post", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	logrusLogger.Info("Post upvoted successfully")
	successRes := response.ClientResponse(http.StatusOK, "Post upvoted successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// DownvotePost godoc
// @Summary Downvote a Post
// @Description Downvote a post by ID
// @Tags Post
// @Produce json
// @Param postid query string true "Post ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /post/downvote [post]

func (ph *PostHandler) DownvotePost(c *gin.Context) {
    logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
    defer logrusLogFile.Close()

    userIDInterface, exists := c.Get("id")
    if !exists {
        err := errors.New("failed to get user ID from context")
        logrusLogger.Error(err)
        errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to get user ID from context", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errorRes)
        return
    }
    userID, ok := userIDInterface.(int)
    if !ok {
        err := errors.New("failed to convert userID to int")
        logrusLogger.Error(err)
        errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to convert userID to int", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errorRes)
        return
    }

    postID := c.Query("postid")
    postIDInt, err := strconv.Atoi(postID)
    if err != nil {
        logrusLogger.Error("Invalid post ID", err)
        errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, err.Error())
        c.JSON(http.StatusBadRequest, errorRes)
        return
    }

    if err := ph.GRPC_Client.DownvotePost(userID, postIDInt); err != nil {
        logrusLogger.Error("Failed to downvote post", err)
        errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to downvote post", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errorRes)
        return
    }

    logrusLogger.Info("Post downvoted successfully")
    successRes := response.ClientResponse(http.StatusOK, "Post downvoted successfully", nil, nil)
    c.JSON(http.StatusOK, successRes)
}