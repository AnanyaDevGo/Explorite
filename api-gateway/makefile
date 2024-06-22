run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/user/user.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/admin/admin.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/auth.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/post/post.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/chat/chat.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/notification/notification.proto

swag:
	swag init -g cmd/main.go -o ./cmd/docs
