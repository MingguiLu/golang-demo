package main

import "fmt"

func main() {

	var price int = 120000

	switch  {
		case price >= 500000:
			fmt.Println("豪车")
		case price >= 200000:
			fmt.Println("中级车")
		case price >= 100000:
			fmt.Println("紧凑型车")
		default:
			fmt.Println("小型车")

	}
}
