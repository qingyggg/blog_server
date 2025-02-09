package socket

import (
	"encoding/json"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/qingyggg/blog_server/pkg/errno"
)

// ConvertResToJson 将响应体转换为json
func ConvertResToJson(res *Response) string {
	// 将结构体转换为 JSON
	jsonData, err := json.Marshal(res)
	if err != nil {
		hlog.Fatal("Error marshalling to JSON:", err)
	}
	return string(jsonData)
}

// GetBaseResponse 获取基本相应
func GetBaseResponse() *Response {
	return &Response{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	}
}

// GetResponse 获取基本相应
func GetResponse(data interface{}) *Response {
	return &Response{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Data:       data,
	}
}

// GetErrResponse 获取错误响应
func GetErrResponse(err error) *Response {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return &Response{
			StatusCode: e.ErrCode,
			StatusMsg:  e.ErrMsg,
		}
	}
	return &Response{
		StatusCode: errno.ServiceErrCode,
		StatusMsg:  err.Error(),
	}
}
