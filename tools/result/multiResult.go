// Package result
// @Description: 一次性列表
package result

import (
	"encoding/json"
	log "github.com/corgi-kx/logcustom"
)

// MultiResult
// @Description: 一次性列表
type MultiResult struct {
	Result
	Data []interface{} `json:"data"`
}

func (receiver *MultiResult) SetData(data []interface{}) {
	receiver.Data = data
}

func (receiver *MultiResult) SetSuccess() *MultiResult {
	receiver.Success = true
	return receiver
}

func (receiver *MultiResult) BuildFailure(code int, message string) *MultiResult {
	receiver.Code = code
	receiver.Message = message
	receiver.Success = false
	return receiver
}

func (receiver *MultiResult) ToJSON() string {
	marshal, err := json.Marshal(receiver)
	if err != nil {
		log.Info("序列化一次性列表返回失败，", err)
	}
	return string(marshal)
}
