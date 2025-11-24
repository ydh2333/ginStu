package main

import (
	ginstu1 "ginStu/ginStu1"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	ginstu1.GetStu1(router)
	ginstu1.PostStu1(router)

	err := router.Run()
	if err != nil {
		panic(err)
	}

}
