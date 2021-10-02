package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//自定义一个字符串
var jwtKey = []byte("customer.xiaoduoai.com")
var NotHasToken = errors.New("生成token失败")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

//颁发token
func setting(UserId string) (string, error) {
	expireTime := time.Now().Add(3 * 24 * time.Hour)
	claims := &Claims{
		UserId: UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "xiaoduoai", // 签名颁发者
			Subject:   "jwt_token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", NotHasToken
	}
	return tokenString, nil
}

//解析token
func getting(tokenString string) (string, bool) {
	token, _, err := ParseToken(tokenString)
	if err != nil {
		return err.Error(), false
	}
	if !token.Valid {
		return "token校验失败", false
	}
	return token.Raw, true
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, Claims, err
}

func MakeToken(UserId string) (string, error) {
	return setting(UserId)
}

func CheckToken(token string) (bool, error) {
	t, tokenIsOk := getting(token)
	if tokenIsOk {
		return true, nil
	}
	return false, errors.New(t)
}
