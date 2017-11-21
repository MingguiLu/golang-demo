package main

import "fmt"

func main() {

/*
	var i int = 1
	for i <= 100 {
		fmt.Println(i)
		i++
	}
*/

	var b int = 15 
	var a int

	for a := 0; a < 10; a++ {
		fmt.Printf("a的值为：%d\n",a)
	}

	for a < b {
		a++
		fmt.Printf("a的值为：%d\n",a)
	}

	numbers := [6]int{1,2,3,5}

	for i,x := range numbers {
		fmt.Printf("第%d位x的值 = %d\n",i,x)
	}
}
