package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "kratos-bbs-demo/api/xiaohongshu/v1"
)



func (s *XiaohongshuService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}

func (s *XiaohongshuService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (s *XiaohongshuService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}

func (s *XiaohongshuService) UpdatetUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatetUser not implemented")
}

func (s *XiaohongshuService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}

func (s *XiaohongshuService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")

}

func (s *XiaohongshuService) UnFollowUser(ctx context.Context, req *v1.UnFollowUserRequest) (*v1.ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowUser not implemented")
}
