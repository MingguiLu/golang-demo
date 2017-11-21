package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("交换之前，a的值是%d\n",a)
	fmt.Printf("交换之前，b的值是%d\n",b)

	swap(a,b)

	fmt.Printf("交换之后，a的值是%d\n",a)
	fmt.Printf("交换之后，b的值是%d\n",b)

	/*
	传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。默认情况下，Go语言使用的是值传递，即在调用过程中不会影响到实际参数。
	*/
}

func swap(num1,num2 int) (int,int) {
	var temp int
	temp = num1
	num1 = num2
	num2 = temp

	return num1,num2
}
