package utils

import (
	"ApiGuide/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 生成token
func CreateToken(user *models.User, ip string) (tokenss string, err error) {
	claim := jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Unix() + 1000,
		"iss":      "jh",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	str := beego.AppConfig.String("SecretKey") + ip
	tokenss, err = token.SignedString([]byte(str))
	return
}

func secret(ip string) jwt.Keyfunc {
	return func(token *jwt.Token) (i interface{}, e error) {
		return []byte(beego.AppConfig.String("SecretKey") + ip), nil
	}
}

// 验证token
func ParseToken(tokenss string, ip string) (user *models.User, err error) {
	user = &models.User{}
	secret_Data := secret(ip)
	token, err := jwt.Parse(tokenss, secret_Data)
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claim)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	user.Id = int(claim["id"].(float64))
	user.Username = claim["username"].(string)
	return
}

func TrueReturn(msg string, data interface{}) map[string]interface{} {
	tartget := make(map[string]interface{})
	tartget["code"] = 200
	tartget["msg"] = msg
	tartget["data"] = data
	return tartget
}

func FalseReturn(code int, msg string, data interface{}) map[string]interface{} {
	target := make(map[string]interface{})
	target["code"] = code
	target["msg"] = msg
	target["data"] = data
	return target
}
