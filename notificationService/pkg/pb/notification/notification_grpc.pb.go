// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/pb/notification/notification.proto

package notification

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
	NotificationService_GetNotification_FullMethodName     = "/notification.NotificationService/GetNotification"
	NotificationService_ReadNotification_FullMethodName    = "/notification.NotificationService/ReadNotification"
	NotificationService_MarkAllAsRead_FullMethodName       = "/notification.NotificationService/MarkAllAsRead"
	NotificationService_GetAllNotifications_FullMethodName = "/notification.NotificationService/GetAllNotifications"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error)
	ReadNotification(ctx context.Context, in *ReadNotificationRequest, opts ...grpc.CallOption) (*ReadNotificationResponse, error)
	MarkAllAsRead(ctx context.Context, in *MarkAllAsReadRequest, opts ...grpc.CallOption) (*MarkAllAsReadResponse, error)
	GetAllNotifications(ctx context.Context, in *GetAllNotificationsRequest, opts ...grpc.CallOption) (*GetAllNotificationsResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error) {
	out := new(GetNotificationResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ReadNotification(ctx context.Context, in *ReadNotificationRequest, opts ...grpc.CallOption) (*ReadNotificationResponse, error) {
	out := new(ReadNotificationResponse)
	err := c.cc.Invoke(ctx, NotificationService_ReadNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) MarkAllAsRead(ctx context.Context, in *MarkAllAsReadRequest, opts ...grpc.CallOption) (*MarkAllAsReadResponse, error) {
	out := new(MarkAllAsReadResponse)
	err := c.cc.Invoke(ctx, NotificationService_MarkAllAsRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetAllNotifications(ctx context.Context, in *GetAllNotificationsRequest, opts ...grpc.CallOption) (*GetAllNotificationsResponse, error) {
	out := new(GetAllNotificationsResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetAllNotifications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationResponse, error)
	ReadNotification(context.Context, *ReadNotificationRequest) (*ReadNotificationResponse, error)
	MarkAllAsRead(context.Context, *MarkAllAsReadRequest) (*MarkAllAsReadResponse, error)
	GetAllNotifications(context.Context, *GetAllNotificationsRequest) (*GetAllNotificationsResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}
func (UnimplementedNotificationServiceServer) ReadNotification(context.Context, *ReadNotificationRequest) (*ReadNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadNotification not implemented")
}
func (UnimplementedNotificationServiceServer) MarkAllAsRead(context.Context, *MarkAllAsReadRequest) (*MarkAllAsReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAllAsRead not implemented")
}
func (UnimplementedNotificationServiceServer) GetAllNotifications(context.Context, *GetAllNotificationsRequest) (*GetAllNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotifications not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotification(ctx, req.(*GetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ReadNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ReadNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_ReadNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ReadNotification(ctx, req.(*ReadNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_MarkAllAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAllAsReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).MarkAllAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_MarkAllAsRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).MarkAllAsRead(ctx, req.(*MarkAllAsReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetAllNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetAllNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetAllNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetAllNotifications(ctx, req.(*GetAllNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotification",
			Handler:    _NotificationService_GetNotification_Handler,
		},
		{
			MethodName: "ReadNotification",
			Handler:    _NotificationService_ReadNotification_Handler,
		},
		{
			MethodName: "MarkAllAsRead",
			Handler:    _NotificationService_MarkAllAsRead_Handler,
		},
		{
			MethodName: "GetAllNotifications",
			Handler:    _NotificationService_GetAllNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/notification/notification.proto",
}
