package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	errLoginJSON    []byte
	invalidUserName = "123"
	invalidSecret   = "456"
)

func init() {
	errLoginJSON, _ = json.Marshal(gin.H{
		"code": http.StatusBadRequest,
		"msg":  "用户不存在",
		"data": gin.H{},
	})
}

func TestNewJwt(t *testing.T) {
	j := NewJwt()
	ty := reflect.TypeOf(j)
	t1 := ty.String()
	if t1 != "*middlewares.Jwt" {
		t.Fatal("错误的类型")
	}
}

// 测试登录
func TestLogin(t *testing.T) {
	var err error
	r := gin.New()
	loginFunc := func(c *gin.Context) {
		j := NewJwt()
		j.Login(c)
	}
	r.GET("/login", loginFunc)

	userName := "testName"
	secret := "testSecret"

	url := fmt.Sprintf("/login?username=%s&secret=%s", userName, secret)
	req := httptest.NewRequest("GET", url, &bytes.Buffer{})

	// 测试 系统用户名|密码为空
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("请求错误")
	}
	eqs := strings.EqualFold(string(body), string(errLoginJSON))
	if !eqs {
		t.Fatal("返回值错误")
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("状态码错误")
	}

	// 用户名密码错误情况
	// 忽略设置环境变量错误情况
	os.Setenv("GOBLOGUSERNAME", invalidUserName)
	os.Setenv("GOBLOGSECRET", invalidSecret)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("请求错误")
	}
	eqs = strings.EqualFold(string(body), string(errLoginJSON))
	if !eqs {
		t.Fatal("返回值错误")
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("状态码错误")
	}

	os.Setenv("GOBLOGUSERNAME", userName)
	os.Setenv("GOBLOGSECRET", secret)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resp = w.Result()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("请求错误")
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("状态码错误")
	}
	var m SuccessResp
	json.Unmarshal(body, &m)
	if reflect.TypeOf(m.Data.Token).String() != "string" {
		t.Fatal("token类型错误")

	}

}

func TestNewToken(t *testing.T) {
	j := NewJwt()
	token := NewToken(j, "foo", "bar")
	expect := j.Tokens[token.content]
	if equal := reflect.DeepEqual(token, expect); !equal {
		t.Fatal("生成token错误")
	}
}
