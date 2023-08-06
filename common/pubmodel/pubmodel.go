package pubmodel
 
import "github.com/zeromicro/go-zero/core/stores/sqlx"
 
// 泛型实现动态执行sql
// conn:对应数据库连接
// query:对应sql语句
// args：对应sql的参数
func GenericQuerySql[T any](conn sqlx.SqlConn, query string, args ...interface{}) ([]*T, error) {
	//1、获取数据库连接
	//conn := sqlx.NewMysql(dsn)
 
	//2、初始化切片
	resp := make([]*T, 0, 10)
 
	//3、执行sql
	err := conn.QueryRows(&resp, query, args...)
 
	if err != nil {
		return nil, err
	}
 
	return resp, nil
 
}