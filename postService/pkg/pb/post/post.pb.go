// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: pkg/pb/post/post.proto

package post

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caption   string `protobuf:"bytes,1,opt,name=caption,proto3" json:"caption,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MediaUrl  string `protobuf:"bytes,3,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	MediaData []byte `protobuf:"bytes,4,opt,name=media_data,json=mediaData,proto3" json:"media_data,omitempty"`
}

func (x *AddPostRequest) Reset() {
	*x = AddPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostRequest) ProtoMessage() {}

func (x *AddPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostRequest.ProtoReflect.Descriptor instead.
func (*AddPostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *AddPostRequest) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *AddPostRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddPostRequest) GetMediaUrl() string {
	if x != nil {
		return x.MediaUrl
	}
	return ""
}

func (x *AddPostRequest) GetMediaData() []byte {
	if x != nil {
		return x.MediaData
	}
	return nil
}

type AddPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddPostResponse) Reset() {
	*x = AddPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostResponse) ProtoMessage() {}

func (x *AddPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostResponse.ProtoReflect.Descriptor instead.
func (*AddPostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{1}
}

func (x *AddPostResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_pb_post_post_proto protoreflect.FileDescriptor

var file_pkg_pb_post_post_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x7f,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x55, 0x72, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x27, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x3e, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_pb_post_post_proto_rawDescOnce sync.Once
	file_pkg_pb_post_post_proto_rawDescData = file_pkg_pb_post_post_proto_rawDesc
)

func file_pkg_pb_post_post_proto_rawDescGZIP() []byte {
	file_pkg_pb_post_post_proto_rawDescOnce.Do(func() {
		file_pkg_pb_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_post_post_proto_rawDescData)
	})
	return file_pkg_pb_post_post_proto_rawDescData
}

var file_pkg_pb_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_pb_post_post_proto_goTypes = []interface{}{
	(*AddPostRequest)(nil),  // 0: post.AddPostRequest
	(*AddPostResponse)(nil), // 1: post.AddPostResponse
}
var file_pkg_pb_post_post_proto_depIdxs = []int32{
	0, // 0: post.Post.AddPost:input_type -> post.AddPostRequest
	1, // 1: post.Post.AddPost:output_type -> post.AddPostResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_post_post_proto_init() }
func file_pkg_pb_post_post_proto_init() {
	if File_pkg_pb_post_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_post_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_post_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_pb_post_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_post_post_proto_goTypes,
		DependencyIndexes: file_pkg_pb_post_post_proto_depIdxs,
		MessageInfos:      file_pkg_pb_post_post_proto_msgTypes,
	}.Build()
	File_pkg_pb_post_post_proto = out.File
	file_pkg_pb_post_post_proto_rawDesc = nil
	file_pkg_pb_post_post_proto_goTypes = nil
	file_pkg_pb_post_post_proto_depIdxs = nil
}
