package logic

import (
	"context"
	"strings"
	"time"

	"marsbook/common/pubmodel"
	"marsbook/service/marsbook/api/internal/mytypes"
	"marsbook/service/marsbook/api/internal/svc"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *mytypes.LoginReq) (resp *mytypes.Reply[mytypes.LoginReplyData], err error) {
	// todo: add your logic here and delete this line

	resp = &mytypes.Reply[mytypes.LoginReplyData]{
		Code:    "200",
		Message: "登陆成功",
	}

	//第一步判断异常值
	if len(strings.TrimSpace(req.Userid)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		resp.Code = "400"
		resp.Message = "用户id或密码为空"
		return resp, nil
	}

	//第二步查询数据库
	query := `SELECT id,userid FROM sys_user t WHERE t.userid = ? AND t.password = ?`
	model, err := pubmodel.GenericQuerySql[mytypes.UseridModel](l.svcCtx.MysqlConn, query, req.Userid, req.Password)
	if err != nil || len(model) == 0 {
		resp.Code = "400"
		resp.Message = "用户id或密码错误"
		return resp, nil
	}

	//第三步生成JWT
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, model[0].Id)
	if err != nil {
		resp.Code = "400"
		resp.Message = "token生成错误"
		return resp, nil
	}

	resp.Data = mytypes.LoginReplyData{
		AccessToken:  jwtToken,
		AccessExpire: accessExpire,
	}

	return resp, nil
}

func(l *LoginLogic) getJwtToken(secretKey string,iat,seconds,userId int64)(string,error){
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodES256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
