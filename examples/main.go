package main

import (
	"encoding/json"
	log "github.com/corgi-kx/logcustom"
	"h7/examples/basics"
)

const userInfo = "{\"ID\":2,\"Name\":\"zy0\",\"Age\":23,\"CheckValue\":false}"

var user basics.User

func main() {
	log.Info(8_000 >> 3)
}

func init() {
	bytes := []byte(userInfo)
	err := json.Unmarshal(bytes, &user)
	if err != nil {
		log.Errorf("序列化失败:%s", err)
	}
}
