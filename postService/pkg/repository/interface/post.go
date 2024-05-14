package interfaces

type PostRepository interface {
	AddPost(userID int, caption string, url string) error
	//ListPost() ([]models.AddPost, error) 
	//UpdatePostByID(postID int, updatedPost models.EditPost) error
	// DeletePostByID(postID int) error 
}
