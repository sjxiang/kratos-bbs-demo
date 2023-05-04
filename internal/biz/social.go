package biz

import (
	"context"

	// v1 "kratos-bbs-demo/api/xiaohongshu/v1"

	// "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)


type ArticleRepo interface {
	
}

type CommentRepo interface {}

type TagRepo interface {}

type SocialUsecase struct {
	ar ArticleRepo
	cr CommentRepo
	tr TagRepo
	
	log *log.Helper
}


func NewSocialUsecase(
	ar ArticleRepo, 
	cr CommentRepo, 
	tr TagRepo, 
	logger log.Logger) *SocialUsecase {
	
	return &SocialUsecase{ar:ar, cr:cr, tr:tr, log: log.NewHelper(logger)}
}


// func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
// 	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
// 	return uc.repo.Save(ctx, g)
// }

func (uc *SocialUsecase) CreateArticle(ctx context.Context) {

}