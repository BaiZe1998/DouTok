// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: comment.proto

package api

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

// ShortVideoCoreCommentServiceClient is the client API for ShortVideoCoreCommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortVideoCoreCommentServiceClient interface {
	CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
	ListComment(ctx context.Context, in *ListCommentRequest, opts ...grpc.CallOption) (*ListCommentResponse, error)
}

type shortVideoCoreCommentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortVideoCoreCommentServiceClient(cc grpc.ClientConnInterface) ShortVideoCoreCommentServiceClient {
	return &shortVideoCoreCommentServiceClient{cc}
}

func (c *shortVideoCoreCommentServiceClient) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	out := new(CommentActionResponse)
	err := c.cc.Invoke(ctx, "/api.ShortVideoCoreCommentService/CommentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortVideoCoreCommentServiceClient) ListComment(ctx context.Context, in *ListCommentRequest, opts ...grpc.CallOption) (*ListCommentResponse, error) {
	out := new(ListCommentResponse)
	err := c.cc.Invoke(ctx, "/api.ShortVideoCoreCommentService/ListComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortVideoCoreCommentServiceServer is the server API for ShortVideoCoreCommentService service.
// All implementations must embed UnimplementedShortVideoCoreCommentServiceServer
// for forward compatibility
type ShortVideoCoreCommentServiceServer interface {
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error)
	ListComment(context.Context, *ListCommentRequest) (*ListCommentResponse, error)
	mustEmbedUnimplementedShortVideoCoreCommentServiceServer()
}

// UnimplementedShortVideoCoreCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShortVideoCoreCommentServiceServer struct {
}

func (UnimplementedShortVideoCoreCommentServiceServer) CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedShortVideoCoreCommentServiceServer) ListComment(context.Context, *ListCommentRequest) (*ListCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComment not implemented")
}
func (UnimplementedShortVideoCoreCommentServiceServer) mustEmbedUnimplementedShortVideoCoreCommentServiceServer() {
}

// UnsafeShortVideoCoreCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortVideoCoreCommentServiceServer will
// result in compilation errors.
type UnsafeShortVideoCoreCommentServiceServer interface {
	mustEmbedUnimplementedShortVideoCoreCommentServiceServer()
}

func RegisterShortVideoCoreCommentServiceServer(s grpc.ServiceRegistrar, srv ShortVideoCoreCommentServiceServer) {
	s.RegisterService(&ShortVideoCoreCommentService_ServiceDesc, srv)
}

func _ShortVideoCoreCommentService_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortVideoCoreCommentServiceServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ShortVideoCoreCommentService/CommentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortVideoCoreCommentServiceServer).CommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortVideoCoreCommentService_ListComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortVideoCoreCommentServiceServer).ListComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ShortVideoCoreCommentService/ListComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortVideoCoreCommentServiceServer).ListComment(ctx, req.(*ListCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortVideoCoreCommentService_ServiceDesc is the grpc.ServiceDesc for ShortVideoCoreCommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortVideoCoreCommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ShortVideoCoreCommentService",
	HandlerType: (*ShortVideoCoreCommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentAction",
			Handler:    _ShortVideoCoreCommentService_CommentAction_Handler,
		},
		{
			MethodName: "ListComment",
			Handler:    _ShortVideoCoreCommentService_ListComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}