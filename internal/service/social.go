package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "kratos-bbs-demo/api/xiaohongshu/v1"
)


func (s *XiaohongshuService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}

func (s *XiaohongshuService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}

func (s *XiaohongshuService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (*v1.SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}

func (s *XiaohongshuService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (*v1.SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}

func (s *XiaohongshuService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (*v1.SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}

func (s *XiaohongshuService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}

func (s *XiaohongshuService) AddComments(ctx context.Context, req *v1.AddCommentsRequest) (*v1.SingleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComments not implemented")
}

func (s *XiaohongshuService) GetComments(ctx context.Context, req *v1.GetCommentsRequest) (*v1.MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}

func (s *XiaohongshuService) DeleteComments(ctx context.Context, req *v1.DeleteCommentsRequest) (*v1.DeleteCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComments not implemented")
}

func (s *XiaohongshuService) FavouriteArticle(ctx context.Context, req *v1.FavouriteArticleRequest) (*v1.SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavouriteArticle not implemented")
}

func (s *XiaohongshuService) UnFavouriteArticle(ctx context.Context, req *v1.UnFavouriteArticleRequest) (*v1.SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavouriteArticle not implemented")
}

func (s *XiaohongshuService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (*v1.TagListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法 GetTags 待实现")
}