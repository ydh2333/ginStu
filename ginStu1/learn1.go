package ginstu1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `json:"u" uri:"u" form:"u"`
	Pwd      string `json:"p" uri:"p" form:"p"`
}

func GetStu1(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":  "pong",
			"code": 200,
		})
	})

	// 使用结构体输出...
	r.GET("/outany", func(ctx *gin.Context) {
		user := User{
			UserName: "zhangsan",
			Pwd:      "123456",
		}

		// ctx.XML(http.StatusOK, user)
		ctx.JSON(http.StatusOK, user)
	})

	// 请求一个网页
	// r.LoadHTMLGlob("templates/*")
	// r.GET("/html", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "hello html test~",
	// 	})
	// })
	// 请求一个网页，带有子目录
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/htmlsub1", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "sub1/index2.tmpl", gin.H{
			"title": "hello html test~",
		})
	})

	// 不管什么类型的请求都可以请求到这里来
	r.Any("/any", func(ctx *gin.Context) {
		ctx.String(200, "any request!")
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world!")
	})

	// 设置一个路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/one", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "v1版本的one请求！")
		})

		v1.GET("/two", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "v1版本的one请求！")
		})
	}

	// 内部重定向
	r.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/ping")
	})

	// 外部重定向
	r.GET("/redirectout", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	// 访问静态文件三种方式
	r.Static("/static1", "./static")           // http://127.0.0.1:8080/static1/1.txt
	r.StaticFS("/static2", http.Dir("static")) // http://127.0.0.1:8080/static2
	r.StaticFile("/static3", "./static/1.txt") // http://127.0.0.1:8080/static3

	// get方式传参,方式一, http://127.0.0.1:8080/user/1111/2222
	r.GET("/user/:u/:p", func(ctx *gin.Context) {
		user := ctx.Param("u")
		pwd := ctx.Param("p")

		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})

	// get方式传参,方式二, http://127.0.0.1:8080/user2?u=111&p=222
	r.GET("/user2", func(ctx *gin.Context) {
		user := ctx.Query("u")
		pwd := ctx.Query("p")

		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})

	// 通过绑定的方式传参,方式一, User结构体需要加uri，注意使用的是ctx.ShouldBindUri(&user)
	r.GET("/user3/:u/:p", func(ctx *gin.Context) {
		user := User{}
		err := ctx.ShouldBindUri(&user)
		if err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})
	// 通过绑定的方式传参,方式二, User结构体需要加form   http://127.0.0.1:8080/user4?u=111111&p=222222
	r.GET("/user4", func(ctx *gin.Context) {
		user := User{}
		err := ctx.ShouldBindQuery(&user)
		if err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

	// 用户认证
	userAuth := r.Group("/user", gin.BasicAuth(gin.Accounts{
		"user1": "123456",
	}))
	userAuth.GET("/name", func(ctx *gin.Context) {
		user := ctx.MustGet(gin.AuthUserKey).(string)
		ctx.String(http.StatusOK, user)
	})
}
