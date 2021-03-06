package api

import (
	"encoding/json"
	"fmt"
	"github.com/appleboy/gin-jwt"
	"go-crud/conf"
	"go-crud/model"
	"go-crud/serializer"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "Pong",
	})
}

func AppPing(c *gin.Context){
	c.JSON(200,serializer.Response{
		Status: 0,
		Data:   nil,
		Msg:    "pong",
		Error:  "",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {

	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}

	// token分析
	claims:=jwt.ExtractClaims(c)
	tmp:=(claims["id"])
	if tmp!=nil{
		// interface{} 不能直接转成int64
		uid,ok:=tmp.(float64)
		if ok {
			// float64传进去会找不到数据
			user,err:=model.GetUser(int64(uid))
			if err==nil{
				return &user
			}
		}
	}

	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Response{
				Status: 40001,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
