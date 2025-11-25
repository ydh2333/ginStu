package jwtstu

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JWTStu() {
	mySigningKey := []byte("MyKey")
	c := MyClaims{
		Username: "zhangsan",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 5,
			ExpiresAt: time.Now().Unix() + 5,
			Issuer:    "zhangsan",
		},
	}

	// StandardClaims
	// MapClaims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Println(t)
	// 加密
	ss, err := t.SignedString(mySigningKey)
	fmt.Printf("%v %v\n", ss, err)
	fmt.Println("------------------------------------")
	time.Sleep(6 * time.Second)

	// 解析
	token, err := jwt.ParseWithClaims(ss, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(token)
	fmt.Println("token.claims:", token.Claims.(*MyClaims).Username)
}
