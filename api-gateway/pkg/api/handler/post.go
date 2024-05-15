package handler

import (
	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/utils/models"
	"ExploriteGateway/pkg/utils/response"
	"errors"
	"fmt"
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

func (ph *PostHandler) AddPost(c *gin.Context) {
	userID, _ := c.Get("id")

	var post models.AddPost
	post.Caption = c.PostForm("caption")
	post.MediaURL = c.PostForm("mediaurl")

	fmt.Printf("user id %v\n %T", userID, userID)

	userId, ok := userID.(int)
	fmt.Println("iddddd", userID)
	if !ok {
		err := errors.New("user id convertion failed")
		errorRes := response.ClientResponse(http.StatusBadRequest, "convertion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	userIdStr := strconv.Itoa(userId)
	post.UserId = userIdStr

	file, err := c.FormFile("image")
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in getting data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	fileContent, err := file.Open()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in opening file", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	defer fileContent.Close()

	imageBytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in reading data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	post.Media = imageBytes

	if err := ph.GRPC_Client.AddPost(post, file); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add post", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added post", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ph *PostHandler) ListPost(c *gin.Context) {
	posts, err := ph.GRPC_Client.ListPost()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Posts cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully listed posts", posts, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ph *PostHandler) EditPost(c *gin.Context) {
	userID, _ := c.Get("id")

	var editPost models.EditPost
	editPost.Caption = c.PostForm("caption")
	editPost.PostId = c.PostForm("postid")

	userId, ok := userID.(int)
	if !ok {
		err := errors.New("user id conversion failed")
		errorRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	editPost.UserId = strconv.Itoa(userId)

	if err := ph.GRPC_Client.EditPost(userId, editPost); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not edit post", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully edited post", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (ph *PostHandler) DeletePost(c *gin.Context) {
	postID := c.Param("postid")
	fmt.Println("post id ", postID)

	if postID == "" {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, "Post ID is empty")
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := ph.GRPC_Client.DeletePost(postIDInt); err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to delete post", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
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

//     successRes := response.ClientResponse(http.StatusOK, "Post unsaved successfully", nil, nil)
//     c.JSON(http.StatusOK, successRes)
// }
func (ph *PostHandler) UpvotePost(c *gin.Context) {
	userIDInterface, exists := c.Get("id")
	if !exists {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to get user ID from context", nil, "")
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to convert userID to int", nil, "")
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	postID := c.Param("postid")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := ph.GRPC_Client.UpvotePost(userID, postIDInt); err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to upvote post", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Post upvoted successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ph *PostHandler) DownvotePost(c *gin.Context) {
	userIDInterface, exists := c.Get("id")
	if !exists {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to get user ID from context", nil, "")
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to convert userID to int", nil, "")
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	postID := c.Param("postid")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid post ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := ph.GRPC_Client.DownvotePost(userID, postIDInt); err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to downvote post", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Post downvoted successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
