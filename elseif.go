package main
import "fmt"

func main() {
	var a int = 100
	var b int = 200

	if a == 100 {
		if b == 200 {
			fmt.Println("a = 100, b = 200")
		}
	}
	fmt.Printf("a = %d,b = %d\n",a,b)
}
