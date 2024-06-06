// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/pb/chat/chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChatService_GetAllChats_FullMethodName   = "/chat.ChatService/GetAllChats"
	ChatService_GetFriendChat_FullMethodName = "/chat.ChatService/GetFriendChat"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	GetAllChats(ctx context.Context, in *GetAllChatsRequest, opts ...grpc.CallOption) (*GetAllChatsResponse, error)
	GetFriendChat(ctx context.Context, in *GetFriendChatRequest, opts ...grpc.CallOption) (*GetFriendChatResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) GetAllChats(ctx context.Context, in *GetAllChatsRequest, opts ...grpc.CallOption) (*GetAllChatsResponse, error) {
	out := new(GetAllChatsResponse)
	err := c.cc.Invoke(ctx, ChatService_GetAllChats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetFriendChat(ctx context.Context, in *GetFriendChatRequest, opts ...grpc.CallOption) (*GetFriendChatResponse, error) {
	out := new(GetFriendChatResponse)
	err := c.cc.Invoke(ctx, ChatService_GetFriendChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	GetAllChats(context.Context, *GetAllChatsRequest) (*GetAllChatsResponse, error)
	GetFriendChat(context.Context, *GetFriendChatRequest) (*GetFriendChatResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) GetAllChats(context.Context, *GetAllChatsRequest) (*GetAllChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllChats not implemented")
}
func (UnimplementedChatServiceServer) GetFriendChat(context.Context, *GetFriendChatRequest) (*GetFriendChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendChat not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_GetAllChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetAllChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetAllChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetAllChats(ctx, req.(*GetAllChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetFriendChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetFriendChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetFriendChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetFriendChat(ctx, req.(*GetFriendChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllChats",
			Handler:    _ChatService_GetAllChats_Handler,
		},
		{
			MethodName: "GetFriendChat",
			Handler:    _ChatService_GetFriendChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/chat/chat.proto",
}