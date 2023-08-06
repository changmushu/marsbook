package mytypes

// 公共结构体
type Reply[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type LoginReq struct {
	Userid   string `json:"userId"`
	Password string `json:"password"`
}

type LoginReplyData struct {
	AccessToken  string `json:"token"`
	AccessExpire int64  `json:"accessExpire"`
}

type UserInfoReq struct {
	Userid string `form:"userId"`
}

type UserInfoReplyData struct {
	Id       int64  `json:"id"`
	Userid   string `json:"userId"`
	Username string `json:"name"`
	Role     string `json:"role"`   //用户角色
	Avatar   string `json:"avatar"` //用户图像
}

type RoleMenuReq struct {
	Role string `form:"role"`
}

// type RoleMenuReply struct {
// 	Code    string      `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"` //空接口，接收任意类型
// }

type UseridModel struct {
	Id     int64
	Userid string
}
