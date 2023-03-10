// Code generated by goctl. DO NOT EDIT.
// Source: action.proto

package action

import (
	"context"

	"SimpleDouYin/app/service/action/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Author           = pb.Author
	Comment          = pb.Comment
	CommentListReq   = pb.CommentListReq
	CommentListResp  = pb.CommentListResp
	CommentReq       = pb.CommentReq
	CommentResp      = pb.CommentResp
	FavoriteListReq  = pb.FavoriteListReq
	FavoriteListResp = pb.FavoriteListResp
	FavoriteReq      = pb.FavoriteReq
	FavoriteResp     = pb.FavoriteResp
	FollowListReq    = pb.FollowListReq
	FollowListResp   = pb.FollowListResp
	FollowReq        = pb.FollowReq
	FollowResp       = pb.FollowResp
	FollowerListReq  = pb.FollowerListReq
	FollowerListResp = pb.FollowerListResp
	FriendListReq    = pb.FriendListReq
	FriendListResp   = pb.FriendListResp
	Video            = pb.Video

	Action interface {
		Favorite(ctx context.Context, in *FavoriteReq, opts ...grpc.CallOption) (*FavoriteResp, error)
		FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
		Comment(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*CommentResp, error)
		CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
		Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error)
		FollowList(ctx context.Context, in *FollowListReq, opts ...grpc.CallOption) (*FollowListResp, error)
		FollowerLost(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListResp, error)
		FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error)
	}

	defaultAction struct {
		cli zrpc.Client
	}
)

func NewAction(cli zrpc.Client) Action {
	return &defaultAction{
		cli: cli,
	}
}

func (m *defaultAction) Favorite(ctx context.Context, in *FavoriteReq, opts ...grpc.CallOption) (*FavoriteResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.Favorite(ctx, in, opts...)
}

func (m *defaultAction) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.FavoriteList(ctx, in, opts...)
}

func (m *defaultAction) Comment(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*CommentResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.Comment(ctx, in, opts...)
}

func (m *defaultAction) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.CommentList(ctx, in, opts...)
}

func (m *defaultAction) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.Follow(ctx, in, opts...)
}

func (m *defaultAction) FollowList(ctx context.Context, in *FollowListReq, opts ...grpc.CallOption) (*FollowListResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.FollowList(ctx, in, opts...)
}

func (m *defaultAction) FollowerLost(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.FollowerLost(ctx, in, opts...)
}

func (m *defaultAction) FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.FriendList(ctx, in, opts...)
}
