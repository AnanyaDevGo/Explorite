package client

import (
	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/config"
	pb "ExploriteGateway/pkg/pb/post"
	"ExploriteGateway/pkg/utils/models"
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"google.golang.org/grpc"
)

type postClient struct {
	Client pb.PostClient
}

func NewPostClient(cfg config.Config) (interfaces.PostClient, error) {
	fmt.Println("client")
	fmt.Println("auth", cfg.ExploritePost)
	grpcConnection, err := grpc.Dial(cfg.ExploritePost, grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial error", err)
		return nil, err
	}

	grpcClient := pb.NewPostClient(grpcConnection)

	return &postClient{
		Client: grpcClient,
	}, nil
}

func (pc *postClient) AddPost(post models.AddPost, file *multipart.FileHeader) error {
	fmt.Println("hereee client gateway")
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	fileData, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	_, err = pc.Client.AddPost(context.Background(), &pb.AddPostRequest{
		Caption:   post.Caption,
		UserId:    post.UserId,
		MediaUrl:  post.MediaURL,
		MediaData: fileData,
	})
	if err != nil {
		fmt.Println("erorrrr inside handler", err)
		return err
	}
	return nil
}

func (pc *postClient) ListPost() ([]models.AddPost, error) {
	resp, err := pc.Client.ListPost(context.Background(), &pb.ListPostRequest{})
	if err != nil {
		return nil, err
	}

	var posts []models.AddPost
	for _, p := range resp.Posts {
		post := models.AddPost{
			Caption:  p.Caption,
			Media:    p.MediaData,
			UserId:   p.UserId,
			MediaURL: p.MediaUrl,
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (pc *postClient) EditPost(id int, post models.EditPost) error {
	fmt.Println("hereeeeeeee")
	_, err := pc.Client.EditPost(context.Background(), &pb.EditPostRequest{
		UserId:  int32(id),
		Caption: post.Caption,
		PostId:  post.PostId,
	})
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

func (pc *postClient) DeletePost(id int) error {
	fmt.Println("hereeee 2")
	fmt.Println("iddddd", id)
	_, err := pc.Client.DeletePost(context.Background(), &pb.DeletePostRequest{
		PostId: int32(id),
	})
	if err != nil {
		return err
	}
	return nil
}

// func (pc *postClient) SavePost(postID int) error {
//     _, err := pc.Client.SavePost(context.Background(), &pb.SavePostRequest{
//         PostId: int32(postID),
//     })
//     if err != nil {
//         return err
//     }
//     return nil
// }

// func (uc *userClient) UnSavePost(postID int) error {
//     _, err := uc.Client.UnSavePost(context.Background(), &pb.UnSavePostRequest{
//         PostId: int32(postID),
//     })
//     if err != nil {
//         return err
//     }
//     return nil
// }

func (pc *postClient) CreateCommentPost(postId, userId int, comment string) (bool, error) {
	resp, err := pc.Client.CreateCommentPost(context.Background(), &pb.CreateCommentRequest{
		PostId:  uint64(postId),
		UserId:  uint64(userId),
		Comment: comment,
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func (pc *postClient) UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error) {
	resp, err := pc.Client.UpdateCommentPost(context.Background(), &pb.UpdateCommentRequest{
		PostId:    uint64(postId),
		UserId:    uint64(userId),
		Comment:   comment,
		CommentId: uint64(commentId),
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func (pc *postClient) DeleteCommentPost(postId, userId, commentId int) (bool, error) {
	resp, err := pc.Client.DeleteCommentPost(context.Background(), &pb.DeleteCommentRequest{
		PostId:    uint64(postId),
		UserId:    uint64(userId),
		CommentId: uint64(commentId),
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}
func (pc *postClient) UpvotePost(userID, postID int) error {
	_, err := pc.Client.UpvotePost(context.Background(), &pb.UpvotePostRequest{
		UserId: int32(userID),
		PostId: int32(postID),
	})
	if err != nil {
		return err
	}
	return nil
}

func (pc *postClient) DownvotePost(userID, postID int) error {
	_, err := pc.Client.DownvotePost(context.Background(), &pb.DownvotePostRequest{
		UserId: int32(userID),
		PostId: int32(postID),
	})
	if err != nil {
		return err
	}
	return nil
}
