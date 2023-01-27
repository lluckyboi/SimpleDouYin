// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: video.proto

package pb

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

// 视频流
type FeedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastTime string `protobuf:"bytes,1,opt,name=LastTime,proto3" json:"LastTime,omitempty"`
}

func (x *FeedReq) Reset() {
	*x = FeedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedReq) ProtoMessage() {}

func (x *FeedReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedReq.ProtoReflect.Descriptor instead.
func (*FeedReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *FeedReq) GetLastTime() string {
	if x != nil {
		return x.LastTime
	}
	return ""
}

type FeedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	NextTime   string   `protobuf:"bytes,3,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
	VideoList  []*Video `protobuf:"bytes,4,rep,name=VideoList,proto3" json:"VideoList,omitempty"`
}

func (x *FeedResp) Reset() {
	*x = FeedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResp) ProtoMessage() {}

func (x *FeedResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResp.ProtoReflect.Descriptor instead.
func (*FeedResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

func (x *FeedResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FeedResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *FeedResp) GetNextTime() string {
	if x != nil {
		return x.NextTime
	}
	return ""
}

func (x *FeedResp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Author        *Author `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
	PlayUrl       string  `protobuf:"bytes,3,opt,name=PlayUrl,proto3" json:"PlayUrl,omitempty"`
	CoverUrl      string  `protobuf:"bytes,4,opt,name=CoverUrl,proto3" json:"CoverUrl,omitempty"`
	FavoriteCount int64   `protobuf:"varint,5,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`
	CommentCount  int64   `protobuf:"varint,6,opt,name=CommentCount,proto3" json:"CommentCount,omitempty"`
	IsFavorite    bool    `protobuf:"varint,7,opt,name=IsFavorite,proto3" json:"IsFavorite,omitempty"`
	Title         string  `protobuf:"bytes,8,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	FollowCount   int64  `protobuf:"varint,3,opt,name=FollowCount,proto3" json:"FollowCount,omitempty"`
	FollowerCount int64  `protobuf:"varint,4,opt,name=FollowerCount,proto3" json:"FollowerCount,omitempty"`
	IsFollow      bool   `protobuf:"varint,5,opt,name=IsFollow,proto3" json:"IsFollow,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

func (x *Author) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *Author) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *Author) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

// 投稿
type PublishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	VideoUrl string `protobuf:"bytes,3,opt,name=VideoUrl,proto3" json:"VideoUrl,omitempty"`
	VideoId  int64  `protobuf:"varint,4,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
	Hash     string `protobuf:"bytes,5,opt,name=Hash,proto3" json:"Hash,omitempty"`
	CoverUrl string `protobuf:"bytes,6,opt,name=CoverUrl,proto3" json:"CoverUrl,omitempty"`
}

func (x *PublishReq) Reset() {
	*x = PublishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReq) ProtoMessage() {}

func (x *PublishReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReq.ProtoReflect.Descriptor instead.
func (*PublishReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{4}
}

func (x *PublishReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublishReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublishReq) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *PublishReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *PublishReq) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *PublishReq) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

type PublishResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
}

func (x *PublishResp) Reset() {
	*x = PublishResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResp) ProtoMessage() {}

func (x *PublishResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResp.ProtoReflect.Descriptor instead.
func (*PublishResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{5}
}

func (x *PublishResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PublishResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

// 发布列表
type PublishListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *PublishListReq) Reset() {
	*x = PublishListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListReq) ProtoMessage() {}

func (x *PublishListReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListReq.ProtoReflect.Descriptor instead.
func (*PublishListReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{6}
}

func (x *PublishListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PublishListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	VideoList  []*Video `protobuf:"bytes,3,rep,name=VideoList,proto3" json:"VideoList,omitempty"`
}

func (x *PublishListResp) Reset() {
	*x = PublishListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListResp) ProtoMessage() {}

func (x *PublishListResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListResp.ProtoReflect.Descriptor instead.
func (*PublishListResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{7}
}

func (x *PublishListResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PublishListResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *PublishListResp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x22, 0x25, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x08,
	0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xf4,
	0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0xa0, 0x01, 0x0a, 0x0a, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x4b, 0x0a, 0x0b, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x28, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x7b, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x32,
	0xa2, 0x01, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x76, 0x12, 0x27, 0x0a, 0x04, 0x46,
	0x65, 0x65, 0x64, 0x12, 0x0e, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12,
	0x11, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x1a, 0x12, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData = file_video_proto_rawDesc
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_proto_rawDescData)
	})
	return file_video_proto_rawDescData
}

var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_video_proto_goTypes = []interface{}{
	(*FeedReq)(nil),         // 0: video.FeedReq
	(*FeedResp)(nil),        // 1: video.FeedResp
	(*Video)(nil),           // 2: video.Video
	(*Author)(nil),          // 3: video.Author
	(*PublishReq)(nil),      // 4: video.PublishReq
	(*PublishResp)(nil),     // 5: video.PublishResp
	(*PublishListReq)(nil),  // 6: video.PublishListReq
	(*PublishListResp)(nil), // 7: video.PublishListResp
}
var file_video_proto_depIdxs = []int32{
	2, // 0: video.FeedResp.VideoList:type_name -> video.Video
	3, // 1: video.Video.Author:type_name -> video.Author
	2, // 2: video.PublishListResp.VideoList:type_name -> video.Video
	0, // 3: video.videoSv.Feed:input_type -> video.FeedReq
	4, // 4: video.videoSv.Publish:input_type -> video.PublishReq
	6, // 5: video.videoSv.PublishList:input_type -> video.PublishListReq
	1, // 6: video.videoSv.Feed:output_type -> video.FeedResp
	5, // 7: video.videoSv.Publish:output_type -> video.PublishResp
	7, // 8: video.videoSv.PublishList:output_type -> video.PublishListResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedReq); i {
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
		file_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedResp); i {
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
		file_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReq); i {
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
		file_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResp); i {
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
		file_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListReq); i {
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
		file_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListResp); i {
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
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}
