// Package result
// Description todo
// Author smash_hq
// Created 2022/12/8 11:39
package result

import (
	"encoding/json"
	log "github.com/corgi-kx/logcustom"
)

// SingleResult
// @Description: 单个
type SingleResult struct {
	Result
	Data interface{} `json:"data"`
}

func (receiver *SingleResult) SetData(data interface{}) {
	receiver.Data = data
}

func (receiver *SingleResult) ToJSON() string {
	marshal, err := json.Marshal(receiver)
	if err != nil {
		log.Info("序列化基础返回失败，", err)
	}
	return string(marshal)
}

func (receiver *SingleResult) BuildSuccess() *SingleResult {
	receiver.Success = true
	return receiver
}

func (receiver *SingleResult) BuildFailure(code int, message string) *SingleResult {
	receiver.Code = code
	receiver.Message = message
	receiver.Success = false
	return receiver
}
