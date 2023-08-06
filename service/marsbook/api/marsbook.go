package main
 
import (
	"flag"
	"fmt"
 
	"marsbook/service/marsbook/api/internal/config"
	"marsbook/service/marsbook/api/internal/handler"
	"marsbook/service/marsbook/api/internal/svc"
 
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)
 
var configFile = flag.String("f", "etc/marsbook-api.yaml", "the config file")
 
func main() {
	flag.Parse()
 
	var c config.Config
	conf.MustLoad(*configFile, &c)
 
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
 
	// ctx := svc.NewServiceContext(c)
	ctx, err := svc.NewServiceContext(c)
	//mysql,redis等资源连接不了,退出启动
	if err != nil {
		return
	}
	handler.RegisterHandlers(server, ctx)
 
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}