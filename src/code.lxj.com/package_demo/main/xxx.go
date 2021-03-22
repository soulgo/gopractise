package main

import (
	"code.lxj.com/arithmetic"
	"fmt"
)

func main() {
	fmt.Println(arithmetic.Name)
	add := arithmetic.Add(10, 20)
	fmt.Println(add)
	sub := arithmetic.Sub(10, 20)
	fmt.Println(sub)
}
