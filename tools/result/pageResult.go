// Package result
// @Description: 分页返回
package result

import (
	"encoding/json"
	log "github.com/corgi-kx/logcustom"
)

// PageResult
// @Description: 分页返回
type PageResult struct {
	Result
	PageSize    int           `json:"pageSize"`
	PageNum     int           `json:"pageNum"`
	CurrentPage int           `json:"currentPage"`
	Total       int           `json:"total"`
	Pages       int           `json:"pages"`
	Data        []interface{} `json:"data"`
}

func (receiver *PageResult) SetPageSize(size int) {
	receiver.PageSize = size
}

func (receiver *PageResult) SetPageNum(pageNum int) {
	receiver.PageNum = pageNum
}

func (receiver *PageResult) SetCurrentPage(current int) {
	receiver.CurrentPage = current
}

func (receiver *PageResult) SetTotal(total int) {
	receiver.Total = total
}

func (receiver *PageResult) SetPages(pages int) {
	receiver.Pages = pages
}

func (receiver *PageResult) SetData(data []interface{}) {
	receiver.Data = data
}

func (receiver *PageResult) ToJSON() string {
	marshal, err := json.Marshal(receiver)
	if err != nil {
		log.Info("序列化分页返回失败，", err)
	}
	return string(marshal)
}
