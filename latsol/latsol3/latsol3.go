package main

import "fmt"

func main() {
	var x, y, o int
	fmt.Scan(&x, &y)
	o = 0
	for x >= y {
		x = x - y
		o = o + 1
	}
	fmt.Print(o)
}
