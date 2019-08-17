package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	CurrentTime int64       `json:"current_ime"`
	Data        interface{} `json:"data"`
}

// 成功返回
func RespSuccess(c *gin.Context, data interface{}) {
	resp := &Response{
		Code:        200,
		Message:     "成功",
		CurrentTime: time.Now().Unix(),
		Data:        data,
	}
	c.JSON(http.StatusOK, resp)
}


func RespWithErr(c *gin.Context, err error) {

	resp := &Response{
		Code:        400,
		Message:     err.Error(),
		CurrentTime: time.Now().Unix(),
		Data:        gin.H{},
	}
	c.JSON(http.StatusOK, resp)
}
