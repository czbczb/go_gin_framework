package main

import (
	"fmt"
	"syscall/js"
)



func main () {

	// obj := js.New()

	

	global := js.Global()

	global.Set("obj", js.TypeObject)

	obj := global.Get("obj")
	obj.Set("name", "我是张三")
	obj.Set("sayName", js.FuncOf(SayName))

	result := global.Get("obj").Call("sayName", "Cui Zongbao")

	fmt.Println(result)
}

func SayName(name string) string{
	return `my name is : ` + name
}


