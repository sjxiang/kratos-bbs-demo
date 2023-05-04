// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: xiaohongshu/v1/xiaohongshu.proto

package v1

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
	Xiaohongshu_Login_FullMethodName              = "/xiaohongshu.v1.Xiaohongshu/Login"
	Xiaohongshu_Register_FullMethodName           = "/xiaohongshu.v1.Xiaohongshu/Register"
	Xiaohongshu_GetCurrentUser_FullMethodName     = "/xiaohongshu.v1.Xiaohongshu/GetCurrentUser"
	Xiaohongshu_UpdatetUser_FullMethodName        = "/xiaohongshu.v1.Xiaohongshu/UpdatetUser"
	Xiaohongshu_GetProfile_FullMethodName         = "/xiaohongshu.v1.Xiaohongshu/GetProfile"
	Xiaohongshu_FollowUser_FullMethodName         = "/xiaohongshu.v1.Xiaohongshu/FollowUser"
	Xiaohongshu_UnFollowUser_FullMethodName       = "/xiaohongshu.v1.Xiaohongshu/UnFollowUser"
	Xiaohongshu_ListArticles_FullMethodName       = "/xiaohongshu.v1.Xiaohongshu/ListArticles"
	Xiaohongshu_FeedArticles_FullMethodName       = "/xiaohongshu.v1.Xiaohongshu/FeedArticles"
	Xiaohongshu_GetArticle_FullMethodName         = "/xiaohongshu.v1.Xiaohongshu/GetArticle"
	Xiaohongshu_CreateArticle_FullMethodName      = "/xiaohongshu.v1.Xiaohongshu/CreateArticle"
	Xiaohongshu_UpdateArticle_FullMethodName      = "/xiaohongshu.v1.Xiaohongshu/UpdateArticle"
	Xiaohongshu_DeleteArticle_FullMethodName      = "/xiaohongshu.v1.Xiaohongshu/DeleteArticle"
	Xiaohongshu_AddComments_FullMethodName        = "/xiaohongshu.v1.Xiaohongshu/AddComments"
	Xiaohongshu_GetComments_FullMethodName        = "/xiaohongshu.v1.Xiaohongshu/GetComments"
	Xiaohongshu_DeleteComments_FullMethodName     = "/xiaohongshu.v1.Xiaohongshu/DeleteComments"
	Xiaohongshu_FavouriteArticle_FullMethodName   = "/xiaohongshu.v1.Xiaohongshu/FavouriteArticle"
	Xiaohongshu_UnFavouriteArticle_FullMethodName = "/xiaohongshu.v1.Xiaohongshu/UnFavouriteArticle"
	Xiaohongshu_GetTags_FullMethodName            = "/xiaohongshu.v1.Xiaohongshu/GetTags"
)

// XiaohongshuClient is the client API for Xiaohongshu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XiaohongshuClient interface {
	// 用户相关
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	UpdatetUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	// 个人简介相关
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	// 关注
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	// 帖子
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleReply, error)
	// 评论
	AddComments(ctx context.Context, in *AddCommentsRequest, opts ...grpc.CallOption) (*SingleCommentReply, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error)
	DeleteComments(ctx context.Context, in *DeleteCommentsRequest, opts ...grpc.CallOption) (*DeleteCommentsReply, error)
	// 帖子 赞同/不赞同
	FavouriteArticle(ctx context.Context, in *FavouriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UnFavouriteArticle(ctx context.Context, in *UnFavouriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	// 标签
	GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*TagListReply, error)
}

type xiaohongshuClient struct {
	cc grpc.ClientConnInterface
}

func NewXiaohongshuClient(cc grpc.ClientConnInterface) XiaohongshuClient {
	return &xiaohongshuClient{cc}
}

func (c *xiaohongshuClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_GetCurrentUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) UpdatetUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_UpdatetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_GetProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_FollowUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_UnFollowUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_ListArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_FeedArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_GetArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_CreateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_UpdateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleReply, error) {
	out := new(DeleteArticleReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_DeleteArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) AddComments(ctx context.Context, in *AddCommentsRequest, opts ...grpc.CallOption) (*SingleCommentReply, error) {
	out := new(SingleCommentReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_AddComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error) {
	out := new(MultipleCommentsReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_GetComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) DeleteComments(ctx context.Context, in *DeleteCommentsRequest, opts ...grpc.CallOption) (*DeleteCommentsReply, error) {
	out := new(DeleteCommentsReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_DeleteComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) FavouriteArticle(ctx context.Context, in *FavouriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_FavouriteArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) UnFavouriteArticle(ctx context.Context, in *UnFavouriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_UnFavouriteArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xiaohongshuClient) GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*TagListReply, error) {
	out := new(TagListReply)
	err := c.cc.Invoke(ctx, Xiaohongshu_GetTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XiaohongshuServer is the server API for Xiaohongshu service.
// All implementations must embed UnimplementedXiaohongshuServer
// for forward compatibility
type XiaohongshuServer interface {
	// 用户相关
	Login(context.Context, *LoginRequest) (*UserReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error)
	UpdatetUser(context.Context, *UpdateUserRequest) (*UserReply, error)
	// 个人简介相关
	GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error)
	// 关注
	FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error)
	UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error)
	// 帖子
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error)
	// 评论
	AddComments(context.Context, *AddCommentsRequest) (*SingleCommentReply, error)
	GetComments(context.Context, *GetCommentsRequest) (*MultipleCommentsReply, error)
	DeleteComments(context.Context, *DeleteCommentsRequest) (*DeleteCommentsReply, error)
	// 帖子 赞同/不赞同
	FavouriteArticle(context.Context, *FavouriteArticleRequest) (*SingleArticleReply, error)
	UnFavouriteArticle(context.Context, *UnFavouriteArticleRequest) (*SingleArticleReply, error)
	// 标签
	GetTags(context.Context, *GetTagsRequest) (*TagListReply, error)
	mustEmbedUnimplementedXiaohongshuServer()
}

// UnimplementedXiaohongshuServer must be embedded to have forward compatible implementations.
type UnimplementedXiaohongshuServer struct {
}

func (UnimplementedXiaohongshuServer) Login(context.Context, *LoginRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedXiaohongshuServer) Register(context.Context, *RegisterRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedXiaohongshuServer) GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedXiaohongshuServer) UpdatetUser(context.Context, *UpdateUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatetUser not implemented")
}
func (UnimplementedXiaohongshuServer) GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedXiaohongshuServer) FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedXiaohongshuServer) UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowUser not implemented")
}
func (UnimplementedXiaohongshuServer) ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedXiaohongshuServer) FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}
func (UnimplementedXiaohongshuServer) GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedXiaohongshuServer) CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedXiaohongshuServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedXiaohongshuServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedXiaohongshuServer) AddComments(context.Context, *AddCommentsRequest) (*SingleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComments not implemented")
}
func (UnimplementedXiaohongshuServer) GetComments(context.Context, *GetCommentsRequest) (*MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedXiaohongshuServer) DeleteComments(context.Context, *DeleteCommentsRequest) (*DeleteCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComments not implemented")
}
func (UnimplementedXiaohongshuServer) FavouriteArticle(context.Context, *FavouriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavouriteArticle not implemented")
}
func (UnimplementedXiaohongshuServer) UnFavouriteArticle(context.Context, *UnFavouriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavouriteArticle not implemented")
}
func (UnimplementedXiaohongshuServer) GetTags(context.Context, *GetTagsRequest) (*TagListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedXiaohongshuServer) mustEmbedUnimplementedXiaohongshuServer() {}

// UnsafeXiaohongshuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XiaohongshuServer will
// result in compilation errors.
type UnsafeXiaohongshuServer interface {
	mustEmbedUnimplementedXiaohongshuServer()
}

func RegisterXiaohongshuServer(s grpc.ServiceRegistrar, srv XiaohongshuServer) {
	s.RegisterService(&Xiaohongshu_ServiceDesc, srv)
}

func _Xiaohongshu_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_GetCurrentUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_UpdatetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).UpdatetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_UpdatetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).UpdatetUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_FollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_UnFollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).UnFollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_UnFollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).UnFollowUser(ctx, req.(*UnFollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_ListArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_FeedArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).FeedArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_FeedArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).FeedArticles(ctx, req.(*FeedArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_GetArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_CreateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_UpdateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_DeleteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_AddComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).AddComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_AddComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).AddComments(ctx, req.(*AddCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_GetComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_DeleteComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).DeleteComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_DeleteComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).DeleteComments(ctx, req.(*DeleteCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_FavouriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavouriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).FavouriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_FavouriteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).FavouriteArticle(ctx, req.(*FavouriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_UnFavouriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFavouriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).UnFavouriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_UnFavouriteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).UnFavouriteArticle(ctx, req.(*UnFavouriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xiaohongshu_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XiaohongshuServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Xiaohongshu_GetTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XiaohongshuServer).GetTags(ctx, req.(*GetTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Xiaohongshu_ServiceDesc is the grpc.ServiceDesc for Xiaohongshu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Xiaohongshu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xiaohongshu.v1.Xiaohongshu",
	HandlerType: (*XiaohongshuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Xiaohongshu_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Xiaohongshu_Register_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _Xiaohongshu_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdatetUser",
			Handler:    _Xiaohongshu_UpdatetUser_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Xiaohongshu_GetProfile_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _Xiaohongshu_FollowUser_Handler,
		},
		{
			MethodName: "UnFollowUser",
			Handler:    _Xiaohongshu_UnFollowUser_Handler,
		},
		{
			MethodName: "ListArticles",
			Handler:    _Xiaohongshu_ListArticles_Handler,
		},
		{
			MethodName: "FeedArticles",
			Handler:    _Xiaohongshu_FeedArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _Xiaohongshu_GetArticle_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _Xiaohongshu_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _Xiaohongshu_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _Xiaohongshu_DeleteArticle_Handler,
		},
		{
			MethodName: "AddComments",
			Handler:    _Xiaohongshu_AddComments_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _Xiaohongshu_GetComments_Handler,
		},
		{
			MethodName: "DeleteComments",
			Handler:    _Xiaohongshu_DeleteComments_Handler,
		},
		{
			MethodName: "FavouriteArticle",
			Handler:    _Xiaohongshu_FavouriteArticle_Handler,
		},
		{
			MethodName: "UnFavouriteArticle",
			Handler:    _Xiaohongshu_UnFavouriteArticle_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _Xiaohongshu_GetTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xiaohongshu/v1/xiaohongshu.proto",
}