package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"fmt"

	"github.com/dgrijalva/jwt-go"

	"github.com/hesen/blog/pkg/setting"
)

// Md5 带size的需要明确转化的编码，如md5.Sum返回值里的byte推迟是按照16进制存储的
// 所以像 string([:]) 这样转化出来会乱码，需要特定格式转化lib来转化
func Md5(content string) string {
	degist := md5.Sum([]byte(content))
	return hex.EncodeToString(degist[:])
}

// JwtClaims 数据结构
/*
	jwt payload

	内置字段:
	(issuer)：签发人
	(expiration time)：过期时间
	(subject)：主题
	(audience)：受众
	(Not Before)：生效时间
	(Issued At)：签发时间
	(JWT ID)：编号
*/
type JwtClaims struct {
	Phone string `json:"phone"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// BuildJwt 生成jwt
// @description 生成jwt
// @param payload JwtClaims jwt负载数据
// @param ExpiresAt int64 过期时间，单位小时
func BuildJwt(payload JwtClaims, ExpiresAt int) (string, error) {
	nowTime := time.Now()

	payload.StandardClaims.Issuer = "hs"
	payload.StandardClaims.IssuedAt = nowTime.Unix()
	payload.StandardClaims.NotBefore = nowTime.Unix()
	payload.StandardClaims.ExpiresAt = nowTime.Add(time.Duration(ExpiresAt) * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(setting.ServerSetting.JwtSecret))

	return tokenString, err
}

// ParseJwt 解析jwt 信息
// @description 解析jwt 信息
// @param tokenString string token字符串
func ParseJwt(tokenString string) (map[string]interface{}, error) {
	// lib会自动判断token是否过期
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }

    // hmacSampleSecret is a []byte containing your secret
    return []byte(setting.ServerSetting.JwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
