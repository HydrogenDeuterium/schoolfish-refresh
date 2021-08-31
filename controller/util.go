package controller

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
)

func returnJson(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  message,
	})
}

//好的返回值不用给信息，能得到数据就行
func ReturnGood(c *gin.Context, data interface{}) {
	returnJson(c, 200, data, "请求成功")
}

//坏的返回值不用给数据，告诉前端哪里有问题
func ReturnError(c *gin.Context, msg string) {
	returnJson(c, 400, nil, msg)
}

//内部错误啥都不用给
func ReturnInternal(c *gin.Context) {
	returnJson(c, 500, nil, "")
}

type Claims struct {
	Uid   uint   `json:"uid"`
	Email string `json:"email"`
	Pwd   string `json:"password"`
	jwt.StandardClaims
}

func getToken(uid uint, email string, pwd string) (string, error) {
	//设置token有效时间
	claims := Claims{
		Uid:   uid,
		Email: email,
		Pwd:   pwd,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
			// 指定token发行人
			Issuer: "back-end",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	jwtSecret := []byte("123456")
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	//去掉数据库读写时间等信息
	delete(data, "CreatedAt")
	delete(data, "UpdatedAt")
	delete(data, "DeletedAt")
	return data
}

func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	jwtSecret := []byte("123456")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
