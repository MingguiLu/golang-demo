package main

import "fmt"
import "unsafe"

func main() {
	const length int = 10
	const width int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = length * width
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)

	const (
		Unknow = 0
		Female = "a"
		Male = true
	)
	println(Unknow)
	println(Female)
	println(Male)

	const (
		d = "abc"
		e = len(d)
		f = unsafe.Sizeof(d)
	)
	println(d,e,f)
}
