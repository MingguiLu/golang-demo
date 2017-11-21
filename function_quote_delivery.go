package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("交换之前,a的值是%d\n",a)
	fmt.Printf("交换之前,b的值是%d\n",b)

	swap(&a,&b)

	fmt.Printf("交换之后,a的值是%d\n",a)
	fmt.Printf("交换之后,b的值是%d\n",b)

	/*
	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数	
	*/
}

func swap(num1 *int, num2 *int) {
	var temp int
	temp = *num1
	*num1 = *num2
	*num2 = temp
}
