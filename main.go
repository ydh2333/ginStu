package main

import (
	jwtstu "ginStu/jwtStu"
)

func main() {
	//router := gin.Default()
	//
	//ginstu1.GetStu1(router)
	//ginstu1.PostStu1(router)
	//
	//err := router.Run()
	//if err != nil {
	//	panic(err)
	//}

	jwtstu.JWTStu()
}
