package basics

import (
	"encoding/json"
	"fmt"
)

// User 发现在进行method作用时无法进行跨包
type User struct {
	ID         int    `id主键`
	Name       string `名称`
	Age        int    `年龄`
	CheckValue bool   `校验值`
}

// AllInterface 实现AllInterface是需要实现内部所有方法，当实现了它内部所有的方法，那么这个struct就实现了这个接口下的所有内容
// AllInterface 当你满足AllInterface了所有条件，那么我就可以说你是一个我的人(实现)
type AllInterface interface {
	TellMeWhoYouAre() int
	HowOldAreYou()
}

func (receiver User) HowOldAreYou() {
	PrintMessage("你多少岁", receiver)

}
func (receiver User) TellMeWhoYouAre() int {
	PrintMessage("告诉我你是谁", receiver)
	return receiver.ID
}

func (receiver User) ToString() string {
	marshal, _ := json.Marshal(receiver)
	return string(marshal)
}

func PrintMessage(major string, user User) {
	fmt.Println("操作类型：", major)
	fmt.Println("结果数据：", user)
}
