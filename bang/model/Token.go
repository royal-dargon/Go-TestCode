package model

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	jwt.StandardClaims
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

var (
	key        = "bang" //salt
	ExpireTime = 604800 //token expire time
)

//一个产生token的函数
func CreatToken(name string, id string) string {
	claims := &JwtClaims{
		UserId:   id,
		UserName: name,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims JwtClaims) (string, error) {
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//验证token的函数
func VerifyToken(token string) (string, error) {
	TempToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", errors.New("token解析失败")
	}
	claims, ok := TempToken.Claims.(*JwtClaims)
	if !ok {
		return "", errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return "", errors.New("发生错误")
	}
	return claims.UserId, nil
}
