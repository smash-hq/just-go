package tools

import log "github.com/corgi-kx/logcustom"

func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func CheckErr(err error) {
	if err != nil {
		log.Errorf("错误：%s", err)
	}
}
