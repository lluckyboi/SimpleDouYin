// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: action.proto

package pb

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

// ActionClient is the client API for Action service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionClient interface {
	Favorite(ctx context.Context, in *FavoriteReq, opts ...grpc.CallOption) (*FavoriteResp, error)
	FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
	Comment(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*CommentResp, error)
	CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
}

type actionClient struct {
	cc grpc.ClientConnInterface
}

func NewActionClient(cc grpc.ClientConnInterface) ActionClient {
	return &actionClient{cc}
}

func (c *actionClient) Favorite(ctx context.Context, in *FavoriteReq, opts ...grpc.CallOption) (*FavoriteResp, error) {
	out := new(FavoriteResp)
	err := c.cc.Invoke(ctx, "/action.Action/Favorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionClient) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	out := new(FavoriteListResp)
	err := c.cc.Invoke(ctx, "/action.Action/FavoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionClient) Comment(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*CommentResp, error) {
	out := new(CommentResp)
	err := c.cc.Invoke(ctx, "/action.Action/Comment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionClient) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	out := new(CommentListResp)
	err := c.cc.Invoke(ctx, "/action.Action/CommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionServer is the server API for Action service.
// All implementations must embed UnimplementedActionServer
// for forward compatibility
type ActionServer interface {
	Favorite(context.Context, *FavoriteReq) (*FavoriteResp, error)
	FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error)
	Comment(context.Context, *CommentReq) (*CommentResp, error)
	CommentList(context.Context, *CommentListReq) (*CommentListResp, error)
	mustEmbedUnimplementedActionServer()
}

// UnimplementedActionServer must be embedded to have forward compatible implementations.
type UnimplementedActionServer struct {
}

func (UnimplementedActionServer) Favorite(context.Context, *FavoriteReq) (*FavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Favorite not implemented")
}
func (UnimplementedActionServer) FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedActionServer) Comment(context.Context, *CommentReq) (*CommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comment not implemented")
}
func (UnimplementedActionServer) CommentList(context.Context, *CommentListReq) (*CommentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedActionServer) mustEmbedUnimplementedActionServer() {}

// UnsafeActionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionServer will
// result in compilation errors.
type UnsafeActionServer interface {
	mustEmbedUnimplementedActionServer()
}

func RegisterActionServer(s grpc.ServiceRegistrar, srv ActionServer) {
	s.RegisterService(&Action_ServiceDesc, srv)
}

func _Action_Favorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).Favorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/action.Action/Favorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).Favorite(ctx, req.(*FavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Action_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/action.Action/FavoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).FavoriteList(ctx, req.(*FavoriteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Action_Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/action.Action/Comment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).Comment(ctx, req.(*CommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Action_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/action.Action/CommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).CommentList(ctx, req.(*CommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Action_ServiceDesc is the grpc.ServiceDesc for Action service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Action_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "action.Action",
	HandlerType: (*ActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Favorite",
			Handler:    _Action_Favorite_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _Action_FavoriteList_Handler,
		},
		{
			MethodName: "Comment",
			Handler:    _Action_Comment_Handler,
		},
		{
			MethodName: "CommentList",
			Handler:    _Action_CommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "action.proto",
}