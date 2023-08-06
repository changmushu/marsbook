package logic

import (
	"context"

	"marsbook/service/marsbook/api/internal/svc"
	"marsbook/service/marsbook/api/internal/mytypes"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *mytypes.UserInfoReq) (resp *mytypes.UserInfoReplyData, err error) {
	// todo: add your logic here and delete this line

	return
}
