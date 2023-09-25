package main

import (
	"ginDemo/ctl"
	mystruct "ginDemo/myStruct"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {

	// 启动一个默认的路由
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", mystruct.BookableDate)
	}

	ck := cookie.NewStore([]byte("key1111"))
	// 设置session 中间件
	r.Use(sessions.Sessions("mysession", ck))

	// 中间件
	r.Use(ctl.CreateTime)
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	// 静态资源加载
	r.Static("/static", "./wwwroot")
	// 一个get请求，相当于.net 的控制器
	r.GET("/sayHellow", func(c *gin.Context) {
		c.JSON(200, gin.H{ // 相当于 mvc 的model
			"丐帮帮主": "乔峰",
		})
	})
	//r.Use(ctl.CookieMiddleware)  // 放在这里，后面的请求都需要经过它，
	r.GET("/index", ctl.CookieMiddleware, ctl.Index) // 只有本次请求才会经过cookie中间件处理
	r.Any("/login", ctl.LoginHandler)
	r.GET("/query", ctl.QueryHandler)
	r.POST("/form", ctl.FormHandler)
	r.GET("/book/:action", ctl.ParamHandler)
	r.GET("/books/:nian/:yue/:ri", ctl.TimeRHandler)
	//r.LoadHTMLFiles("./upload.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	//r.MaxMultipartMemory=8*1000  设置上传大文件
	// 上传文件
	r.POST("/upload", ctl.UploadHandler)
	// 实体绑定， 直接将传入的参数 弄成对象来处理
	r.POST("/dataBin", ctl.DataBin)
	// 启动webservice
	r.Run(":7000")

}
