package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("whyy1")

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func GetToken(userid int64) string {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenString)
	return tokenString
}

func QueryToken(tokenString string) (int64, bool) {
	//如果token为空或者未找到
	if tokenString == "" {
		return 0, false
	}
	//fmt.Println(tokenString)
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return 0, false
	}
	return claims.UserId, true
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
