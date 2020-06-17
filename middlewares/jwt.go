package middlewares

// 简易版jwt
// 用户名和密钥使用系统环境变量
// 签名算法使用sha1
// 过期时间可配置 默认2h

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// TokenHeader token在header头的名称
	TokenHeader = "Authorization"
)

// SuccessResp 成功返回结构体
type SuccessResp struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data SuccessRespBody `json:"data"`
}

// SuccessRespBody 成功返回data结构体
type SuccessRespBody struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

// Jwt json token验证器
type Jwt struct {
	// 过期时间
	Timeout time.Duration
	// token存储
	Tokens map[string]Token
}

// NewJwt 生成jwt
func NewJwt() *Jwt {
	return &Jwt{
		Timeout: 2 * time.Hour,
		Tokens:  make(map[string]Token),
	}
}

// IsAuth 是否认证
func (j *Jwt) IsAuth(c *gin.Context) {
	tokenKey := c.Request.Header.Get(TokenHeader)
	if tokenKey == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token不存在",
			"data": gin.H{},
		})
	}
	token, ok := j.Tokens[tokenKey]
	// token不存在
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token失效",
			"data": gin.H{},
		})
	}

	// token过期
	if time.Now().After(token.expiredTime) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "token已过期",
			"data": gin.H{},
		})
	}
	c.Next()
}

// Login 授权登录
// username=USERNAME&secret=SECRET
func (j *Jwt) Login(c *gin.Context) {
	goBlogUsername := os.Getenv("GOBLOGUSERNAME")
	goBlogSecret := os.Getenv("GOBLOGSECRET")
	if goBlogUsername == "" || goBlogSecret == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户不存在",
			"data": gin.H{},
		})
		return
	}
	username := c.Query("username")
	secret := c.Query("secret")
	if username != goBlogUsername || secret != goBlogSecret {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户不存在",
			"data": gin.H{},
		})
		return
	}
	token := NewToken(j, username, secret)
	c.JSON(http.StatusOK, SuccessResp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: SuccessRespBody{
			Token:  token.content,
			Expire: token.expiredTime.Format("2006-01-02 15:04:05"),
		},
	})
}

// Token token结构
type Token struct {
	expiredTime time.Time
	content     string
}

// NewToken 生成token
func NewToken(j *Jwt, params ...string) Token {
	var content string
	h := sha1.New()
	for _, item := range params {
		io.WriteString(h, item)
	}
	io.WriteString(h, time.Now().String())
	sum := h.Sum(nil)
	content = fmt.Sprintf("%x", sum)
	expiredTime := time.Now().Add(j.Timeout)
	t := Token{}
	t.expiredTime = expiredTime
	t.content = string(content)
	j.Tokens[t.content] = t
	return t
}
