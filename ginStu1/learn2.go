package ginstu1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserDetail struct {
	Gender bool   `json:"g" form:"g"`
	Age    uint   `json:"a" form:"a"`
	Email  string `json:"e" form:"e" binding:"required,email"`
}

func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("mw1 before")
		ctx.Next()
		fmt.Println("mw1 after")
	}
}
func MiddleWare2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("mw2 before")
		ctx.Next()
		fmt.Println("mw2 after")
	}

}

func PostStu1(r *gin.Engine) {
	// 这个中间件将用于所有请求
	// r.Use(MiddleWare())
	// 这两个中间件只用于/login5这一个请求
	r.POST("/login5", MiddleWare(), MiddleWare2(), func(ctx *gin.Context) {
		user := ctx.PostForm("u")
		pwd := ctx.PostForm("p")
		fmt.Println("接口本体")

		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})

	// 这种接收方式适合参数少的情况
	r.POST("/login", func(ctx *gin.Context) {
		user := ctx.PostForm("u")
		pwd := ctx.PostForm("p")

		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	// 数据多的时候使用绑定的方式
	r.POST("/login2", func(ctx *gin.Context) {
		user := User{}
		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})
	// 想多次绑定
	r.POST("/login3", func(ctx *gin.Context) {
		user := User{}
		userDetail := UserDetail{}

		if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err == nil {
			fmt.Println(user)
			ctx.String(http.StatusOK, `the body should be User!`)
		} else if err2 := ctx.ShouldBindBodyWith(&userDetail, binding.JSON); err2 == nil {
			fmt.Println(userDetail)
			ctx.String(http.StatusOK, `the body should be UserDetail!`)
		}
	})

	// 模型验证
	r.POST("/login4", func(ctx *gin.Context) {
		userDetail := UserDetail{}
		err := ctx.ShouldBind(&userDetail)
		if err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, userDetail)
	})
}
