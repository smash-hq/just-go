// Package result 定义一个公共返回的包结构
package result

import (
	"encoding/json"
	log "github.com/corgi-kx/logcustom"
)

// Result
// @Description: 公共返回
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// SetCode
// @Description: 设置code
// @receiver receiver
// @param code
func (receiver *Result) SetCode(code int) {
	receiver.Code = code
}

// SetMessage
// @Description: 设置返回信息
// @receiver receiver
// @param message
func (receiver *Result) SetMessage(message string) {
	receiver.Message = message
}

// SetSuccess
// @Description: 设置是否成功
// @receiver receiver
// @param success
func (receiver *Result) SetSuccess(success bool) {
	receiver.Success = success
}

func (receiver *Result) ToJSON() string {
	marshal, err := json.Marshal(receiver)
	if err != nil {
		log.Info("序列化基础返回失败，", err)
	}
	return string(marshal)
}

func (receiver *Result) BuildSuccess() *Result {
	receiver.Success = true
	return receiver
}

func (receiver *Result) BuildFailure(code int, message string) *Result {
	receiver.Code = code
	receiver.Message = message
	receiver.Success = false
	return receiver
}
