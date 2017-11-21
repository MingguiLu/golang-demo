package main

import "fmt"

func main() {
	var a int = 30
	if a < 20 {
		fmt.Println("a 小于 20 !")
	} else {
		fmt.Println("a 不小于 20 !")
	}
	fmt.Printf("a = %d\n",a)
}
