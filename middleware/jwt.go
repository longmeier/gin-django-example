package middleware

import (
	"errors"
	jwt2 "gin-django-example/pkg/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JwtUser struct {
	User_id  int    `json:"user_id"`
	Username string `json:"username"`
	exp      int    `json:"exp"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(jwt2.JwtKey), //可以设置过期时间
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims JwtUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*JwtUser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtUser{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*JwtUser); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"errcode": "408",
				"errmsg":  "未登录",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"errcode": "403",
					"errmsg":  "授权已过期，请重新登陆",
				})
				c.Abort()
				return
			} else {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"errcode": "405",
					"errmsg":  "认证失败",
				})
				c.Abort()
				return
			}
		}
		c.Set("requestUserId", claims.User_id)
		c.Next()
	}
}

//// AuthGrpcToken 验证Token
//func AuthGrpcToken(ctx context.Context) (code int32, msg string) {
//	code, msg = 200, "token认权成功"
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		code, msg = 500, "metadata异常"
//	}
//	if val, ok := md["token"]; ok {
//		token := val[0]
//		if token == "" {
//			code, msg = 408, "无token认证信息"
//		}
//		j := NewJWT()
//		// parseToken 解析token包含的信息
//		_, err := j.ParseToken(token)
//		if err != nil {
//			(fmt.Sprintf("请求接口拦截器校验失败，token：%v", token)
//			if err == TokenExpired {
//				code, msg = 403, "token授权已过期，请重新登录。"
//			} else {
//				code, msg = 405, "token认证失败"
//			}
//		}
//	}
//	return
//}
