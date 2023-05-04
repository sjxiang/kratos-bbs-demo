package service

import (
	// "context"

	v1 "kratos-bbs-demo/api/xiaohongshu/v1"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-bbs-demo/internal/biz"
)

type XiaohongshuService struct {
	v1.UnimplementedXiaohongshuServer


	log *log.Helper
	uc *biz.GreeterUsecase

}

func NewXiaohongshuService(logger log.Logger) *XiaohongshuService {
	return &XiaohongshuService{
		log: log.NewHelper(logger),
	}
}

