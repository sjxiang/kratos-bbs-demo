package service

import (
	"github.com/google/wire"
	"github.com/go-kratos/kratos/v2/log"

	"kratos-bbs-demo/internal/biz"
	v1 "kratos-bbs-demo/api/xiaohongshu/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewXiaohongshuService)


type XiaohongshuService struct {
	v1.UnimplementedXiaohongshuServer
	uc  *biz.UserUsecase
	sc  *biz.SocialUsecase

	log *log.Helper
}

func NewXiaohongshuService(uc *biz.UserUsecase, logger log.Logger) *XiaohongshuService {
	return &XiaohongshuService{
		uc: uc,
		log: log.NewHelper(logger),
	}
}

