package main

import "fmt"

func main() {
	result  := max(5,9)
	fmt.Println(result)

	a,b := swap("mingguilu","didadandan")
	fmt.Printf("%s love %s\n", a,b)
}

func max(num1,num2 int) int {

		if num1>num2 {
				return num1
		} else {
				return num2
		}
}

func swap(str1,str2 string) (string,string) {
	return str1,str2
}
