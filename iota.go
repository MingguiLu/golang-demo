package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a,b,c)

	const (
		d = iota
		e
		f
	)
	fmt.Println(d,e,f)

	const (
		g = iota
		h
		i
		j = "ha"
		k
		l = 100
		m
		n = iota
		o
	)
	fmt.Println(g,h,i,j,k,l,m,n,o)

	const (
		w=1<<iota
		x=3<<iota
		y
		z
	)
	fmt.Println("w=",w)
	fmt.Println("x=",x)
	fmt.Println("y=",y)
	fmt.Println("z=",z)
}
