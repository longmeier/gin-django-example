package eye

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode             = 200
	BadRequestCode          = 400
	UnAuthorizedCode        = 401
	NotFoundCode            = 404
	InternalServerErrorCode = 500
)

func StatusText(code int) string {
	switch code {
	case SuccessCode:
		return "操作成功！"
	case BadRequestCode:
		return "请求参数错误！"
	case UnAuthorizedCode:
		return "未认证"
	case NotFoundCode:
		return "404未找到"
	case InternalServerErrorCode:
		return "服务器错误"
	default:
		return "unknown error"
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EyeCode int

func HandleSuccess(ctx *gin.Context, httpCode int, httpMsg string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	message := fmt.Sprintf("%v %v", StatusText(httpCode), httpMsg)
	resp := Response{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode int, httpMsg string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	message := fmt.Sprintf("%v %v", StatusText(httpCode), httpMsg)
	resp := Response{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
	//sentry.Client.CaptureException(errors.New(message))
	ctx.JSON(httpCode, resp)
}
