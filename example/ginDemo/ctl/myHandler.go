package ctl

import (
	"fmt"
	mystruct "ginDemo/myStruct"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	fmt.Println("kkkkkkkkk", ctx.Keys["key1"]) // 获取中间件的key

	// username, err := ctx.Cookie("userName")
	// if err != nil {
	// 	ctx.Redirect(302, "/login")
	// }
	username := ctx.Keys["username"]
	if username == "" {
		ctx.Redirect(302, "/login")
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{"imgUrl": "/static/images/g.jpg", "username": username})
}

func DataBin(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		var u mystruct.UserInfo
		ctx.ShouldBind(&u) //  mvc 是直接传实体 默认的那种提交方式
		//ctx.ShouldBindJSON("")// json 绑定 等等很多的
		ctx.JSON(http.StatusOK, gin.H{"userName": u.UserName, "pwd": u.Password})
	}
}

func CookieMiddleware(ctx *gin.Context) {
	username, err := ctx.Cookie("userName")
	if err != nil {
		ctx.Redirect(302, "/login")
		return
	}
	ctx.Set("username", username)
	ctx.Next()

}

func LoginHandler(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		var u mystruct.UserInfo
		ctx.ShouldBind(&u)
		fmt.Println("---------------", u.UserName)
		if u.UserName == "jet" {

			// 10s 过期
			ctx.SetCookie("userName", u.UserName, 10, "/", "127.0.0.1", false, true)
			ctx.Redirect(302, "/index") // 页面跳转
		}

	} else {
		ctx.HTML(http.StatusOK, "login.html", nil)
	}
}

// 参数处理
func QueryHandler(ctx *gin.Context) {
	// 获取不到用默认的
	name := ctx.DefaultQuery("name", "Jet")
	city := ctx.Query("city")
	ctx.JSON(http.StatusOK, gin.H{"name": name, "city": city})
}

// form 参数处理
func FormHandler(ctx *gin.Context) {
	// 表单获取
	name := ctx.PostForm("name")
	city := ctx.PostForm("city")
	ctx.JSON(http.StatusOK, gin.H{"name": name, "city": city})
}

// url中获取参数 相当于mvc user/{id}
func ParamHandler(ctx *gin.Context) {
	urlParm := ctx.Param("action")

	ctx.JSON(http.StatusOK, gin.H{
		"action": urlParm,
	})

}

func UploadHandler(c *gin.Context) {

	// 提取用户上传的文件
	fileObj, err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	// fileobj：上传的文件对象
	// fileobj.filename // 拿到上传文件的文件名
	filePath := fmt.Sprintf("./%s", fileObj.Filename)
	// 保存文件到本地的路径
	c.SaveUploadedFile(fileObj, filePath)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

func CreateTime(ctx *gin.Context) {
	ctx.Set("key1", "12344") // 后续处理的中间件 都可以拿到
	start := time.Now()
	ctx.Next()
	end := time.Since(start)
	fmt.Println("时间", end)

}

func TimeRHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"nian": ctx.Param("nian"),
		"yue":  ctx.Param("yue"),
		"ri":   ctx.Param("ri"),
	})
}
