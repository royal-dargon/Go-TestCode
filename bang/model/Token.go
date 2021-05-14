package model

import (
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
