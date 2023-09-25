package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

// go-micro 插件方式
func main() {
	//注册
	registry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"121.4.181.166:8500",
		}
	})
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"version": "1.0.0",
		})
	})

	server := web.NewService(
		web.Name("gin-web"),
		web.Registry(registry),
		web.Handler(r),
	)
	server.Init() ////加了这句就可以使用命令行的形式去设置我们一些启动的配置 --server_address=:8080
	server.Run()
}
