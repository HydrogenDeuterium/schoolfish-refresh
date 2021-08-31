package controller

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
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
func returnGood(c *gin.Context, data interface{}) {
	returnJson(c, 200, data, "请求成功")
}

//坏的返回值不用给数据，告诉前端哪里有问题
func returnError(c *gin.Context, msg string) {
	returnJson(c, 400, nil, msg)
}

//内部错误啥都不用给
func returnInternal(c *gin.Context) {
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

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		//for k,v :=range c.Request.Header {
		//	fmt.Println(k,v)
		//}
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			returnError(c, "请求头中auth为空")
			c.Abort()
			return
		}
		//别问我为啥他要这么处理，非要在前面加个“Bearer ”
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			returnError(c, "请求头中auth格式有误")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		claim, err := ParseToken(parts[1])
		if err != nil {
			returnError(c, "无效的Token")
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("email", claim.Email)
		c.Set("uid", claim.Uid)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
