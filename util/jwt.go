package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//owner密钥
var jwtSecretOwner = []byte("owner密钥")

//manager密钥
var jwtSecretManager = []byte("manager密钥")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成ownertoken
func GenerateOwnerToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(20 * time.Minute)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecretOwner)

	return token, err
}

//生成ownertoken
func GenerateManagerToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(20 * time.Minute)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecretManager)

	return token, err
}

//解析ownertoken
func ParseOwnerToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretOwner, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

//解析managertoken
func ParseManagerToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretManager, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
