package utils

import (
    "fmt"
	"github.com/go-kratos/bingfood-client-micro/app/user/service/global"
    "log"
    "time"

    jwt "github.com/dgrijalva/jwt-go"
    "github.com/pkg/errors"
)

var (
	TokenIsExpired     = errors.New("Token is expired")
	TokenIsMalformed   = errors.New("Token is malformed")
	TokenIsNotActivity = errors.New("Token is not activity yet")
	TokenIsInvalid     = errors.New("Token is invalid")
)

var mySecret = global.JWT_SECRET // 密钥, 不同的算法需要传入密钥数据类型是不一样的!!SigningMethodHS256算法是要传入字节数组作为密钥

func CreateToken(userMobile string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24).Unix() //token过期时间
	claims := &UserClaims{
		UserMobile: userMobile,
		StandardClaims: jwt.StandardClaims{ // TODO 这些也是配置，后面需要用yaml管理
			ExpiresAt: expireTime, //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  //签名颁发机构
			Subject:   "user token", //签名主题
		},
	}
	fmt.Println(userMobile)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 把Claims组装进去,此时的token有两部分了，准备签名
	fmt.Printf("token:%v\n", token)
	tokenString, err := token.SignedString(mySecret) // 传入密钥进行签名，完成token的组装
	if err != nil {
		log.Printf("token签名失败 : %v", err)
		return "", err
	}
	log.Printf("token生成完成,token is  : %v", tokenString) // 组装完的token
	return tokenString, nil
}

type UserClaims struct {
	UserMobile string
	jwt.StandardClaims
}

func ParseToken(tokenString string) (*UserClaims, error) { //解析tokenString，tokenString是三部分的含签名的token
	claims := &UserClaims{}
	t, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil // 传入MySecret密钥，方便拆分出token的前两部分
	}) //ParseWithClaims解析出token中的前面两个部分，得到的t是两部分的token
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return claims, TokenIsMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return claims, TokenIsExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return claims, TokenIsNotActivity
			} else {
				return claims, TokenIsInvalid
			}
		}
	}
	if t != nil {
		if claims, ok := t.Claims.(*UserClaims); ok && t.Valid {
			return claims, nil
		}
	}
	return nil, TokenIsInvalid
}
