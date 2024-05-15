package service

import (
	"context"
	"fmt"
	pb "postservice/pkg/pb/post"
	interfaces "postservice/pkg/usecase/interface"
	"postservice/pkg/utils/models"
)

type PostServer struct {
	postUsecase interfaces.PostUsecase
	pb.UnimplementedPostServer
}

func NewPostServer(usecase interfaces.PostUsecase) pb.PostServer {
	return &PostServer{
		postUsecase: usecase,
	}
}

func (ps *PostServer) AddPost(ctx context.Context, req *pb.AddPostRequest) (*pb.AddPostResponse, error) {

	fmt.Println("Hereeeeeeeeee")
	post := models.AddPost{
		Caption:  req.Caption,
		UserId:   req.UserId,
		MediaURL: req.MediaUrl,
		Media:    req.MediaData,
	}
	err := ps.postUsecase.AddPost(post)
	if err != nil {
		return &pb.AddPostResponse{
			Error: err.Error(),
		}, err
	}
	return &pb.AddPostResponse{
		Error: "",
	}, nil
}

func (ps *PostServer) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostResponse, error) {
	posts, err := ps.postUsecase.ListPost()
	if err != nil {
		return nil, err
	}

	var pbPosts []*pb.Posts
	for _, post := range posts {
		pbPost := &pb.Posts{
			Caption:   post.Caption,
			UserId:    post.UserId,
			MediaUrl:  post.MediaURL,
			MediaData: post.Media,
		}
		pbPosts = append(pbPosts, pbPost)
	}

	response := &pb.ListPostResponse{
		Posts: pbPosts,
	}
	return response, nil
}

func (ps *PostServer) EditPost(ctx context.Context, req *pb.EditPostRequest) (*pb.EditPostResponse, error) {
	id := req.UserId
	caption := req.Caption
	postId := req.PostId

	editPost := models.EditPost{
		Caption: caption,
		PostId:  postId,
	}

	err := ps.postUsecase.EditPost(int(id), editPost)
	if err != nil {
		return nil, err
	}

	response := &pb.EditPostResponse{
		Error: "",
	}
	return response, nil
}

func (ps *PostServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	postID := req.PostId

	err := ps.postUsecase.DeletePost(int(postID))
	if err != nil {
		return nil, err
	}
	response := &pb.DeletePostResponse{
		Message: "Post deleted successfully",
	}
	return response, nil
}

// func (ps *PostServer) SavePost(ctx context.Context, req *pb.SavePostRequest) (*pb.SavePostResponse, error) {
//     postID := req.PostId

//     err := ps.postUsecase.SavePost(postID)
//     if err != nil {
//         return &pb.SavePostResponse{
//             Success:      false,
//             ErrorMessage: err.Error(),
//         }, nil
//     }

//     return &pb.SavePostResponse{
//         Success: true,
//     }, nil
// }

// func (ps *PostServer) UnSavePost(ctx context.Context, req *pb.UnSavePostRequest) (*pb.UnSavePostResponse, error) {
//     postID := req.PostId

//     err := ps.postUsecase.UnSavePost(postID)
//     if err != nil {
//         return &pb.UnSavePostResponse{
//             Success:      false,
//             ErrorMessage: err.Error(),
//         }, nil
//     }

//     return &pb.UnSavePostResponse{
//         Success: true,
//     }, nil
// }

func (ps *PostServer) UpvotePost(ctx context.Context, req *pb.UpvotePostRequest) (*pb.UpvotePostResponse, error) {
	fmt.Println("hereeeee")
	userID := req.UserId
	postID := req.PostId

	err := ps.postUsecase.UpvotePost(int(userID), int(postID))
	if err != nil {
		return &pb.UpvotePostResponse{
			Error: err.Error(),
		}, err
	}

	return &pb.UpvotePostResponse{}, nil
}

func (ps *PostServer) DownvotePost(ctx context.Context, req *pb.DownvotePostRequest) (*pb.DownvotePostResponse, error) {
	userID := req.UserId
	postID := req.PostId

	err := ps.postUsecase.DownvotePost(int(userID), int(postID))
	if err != nil {
		return &pb.DownvotePostResponse{
			Error: err.Error(),
		}, err
	}

	return &pb.DownvotePostResponse{}, nil
}
