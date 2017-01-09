package main

import (
	"log"
)

type User struct {
	Name  string
	Email string
}

/**
定义接口方法
*/
type Notifier interface {
	Notify() error
}

type Notifier2 interface {
	Notify2() error
}

/**
调用接口方法
*/
func SendNotification(notify Notifier) error {
	return notify.Notify()
}

/**
调用接口方法
*/
func SendNotification2(notify Notifier2) error {
	return notify.Notify2()
}

/**
实现接口方法
(由于定义在街口中 调用Notify必须为*User )
每个类型都有与之关联的⽅方法集，这会影响到接⼝口实现规则。
• 类型 T ⽅方法集包含全部 receiver T ⽅方法。
• 类型 *T ⽅方法集包含全部 receiver T + *T ⽅方法。
• 如类型 S 包含匿名字段 T，则 S ⽅方法集包含 T ⽅方法。
• 如类型 S 包含匿名字段 *T，则 S ⽅方法集包含 T + *T ⽅方法。
• 不管嵌⼊入 T 或 *T，*S ⽅方法集总是包含 T + *T ⽅方法。
*/
func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)
	return nil
}

func (u User) Notify2() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)
	return nil
}

func main() {
	user := User{
		Name:  "AriesDevil",
		Email: "ariesdevil@xxoo.com",
	}

	//SendNotification(user) err
	SendNotification(&user)
	SendNotification2(user)
	SendNotification2(&user)
}
