package zy0

import (
	"fmt"
	"h7/tools"
	"os"
)

func Precious() {

	var name string
	fmt.Println("请输入你的名字:")
	fmt.Scan(&name)

	zy0 := "钟妍灵"
	if name == zy0 {
		writeFile(name)
		return
	}
	fmt.Println("口令错误，不告诉你了")
}

func writeFile(name string) {
	zy0 := "zy0.txt"
	zy0File, err := os.Create(zy0)
	tools.CheckErr(err)
	defer func(zy0File *os.File) {
		err := zy0File.Close()
		tools.CheckErr(err)
	}(zy0File)
	zy0File.WriteString("呜呼，我叫" + name + "，我是瓜猫儿！！！")
}
