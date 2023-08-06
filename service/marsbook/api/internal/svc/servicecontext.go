package svc
 
import (
	"marsbook/service/marsbook/api/internal/config"
	"marsbook/service/marsbook/api/internal/middleware"
 
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)
 
type ServiceContext struct {
	Config    config.Config
	Cors      rest.Middleware
	MysqlConn sqlx.SqlConn //mysql连接对象
}
 
func NewServiceContext(c config.Config) (*ServiceContext, error) {
	// 测试连接mysql数据库是否成功
	conn, err := MysqlTestConn(c)
	if err != nil {
		return &ServiceContext{}, err
	}
 
	return &ServiceContext{
		Config:    c,
		Cors:      middleware.NewCorsMiddleware().Handle,
		MysqlConn: conn, //返回mysql连接对象
	}, err
}
 
type MysqlTestConnStruct struct {
	Sysdate string
}
 
// 测试连接mysql数据库是否成功
func MysqlTestConn(c config.Config) (sqlx.SqlConn, error) {
	//1、返回一个连接池，没有真正发起mysql连接
	conn := sqlx.NewMysql(c.Mysql.DataSource) //返回一个连接池
 
	//2、mysql数据库初始化连接
	var mysqlTestConn MysqlTestConnStruct
	var err error = nil
	err = conn.QueryRow(&mysqlTestConn, " SELECT  SYSDATE();")
	if err != nil {
		logx.Error("mysql初始化失败！", err)
	} else {
		logx.Info("mysql数据库初始化连接成功")
	}
 
	return conn, err
 
}