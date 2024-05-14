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

// func (pc *postClient) ListPost() ([]models.AddPost, error) {
// 	resp, err := pc.Client.ListPost(context.Background(), &pb.ListPostRequest{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var posts []models.AddPost
// 	for _, p := range resp.Posts {
// 		post := models.AddPost{
// 			Caption:  p.Caption,
// 			Media:    p.Media,
// 			UserId:   p.UserId,
// 			MediaURL: p.MediaURL,
// 		}
// 		posts = append(posts, post)
// 	}
// 	return posts, nil
// }

// func (pc *postClient) EditPost(id int, post models.EditPost) error {
// 	_, err := pc.Client.EditPost(context.Background(), &pb.EditPostRequest{
// 		Id:      int32(id),
// 		Caption: post.Caption,
// 		PostId:  post.PostId,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (pc *postClient) DeletePost(id int) error {
//     _, err := pc.Client.DeletePost(context.Background(), &pb.DeletePostRequest{
//         PostId: int32(id),
//     })
//     if err != nil {
//         return err
//     }
//     return nil
// }
