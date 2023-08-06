package logic

import (
	"context"

	"marsbook/service/marsbook/api/internal/svc"
	"marsbook/service/marsbook/api/internal/mytypes"

	"github.com/zeromicro/go-zero/core/logx"
)

type RolemenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRolemenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolemenuLogic {
	return &RolemenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RolemenuLogic) Rolemenu(req *mytypes.RoleMenuReq) (resp *mytypes.Reply[interface{}], err error) {
	// todo: add your logic here and delete this line

	return
}
