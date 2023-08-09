package main

import (
	"fmt"

	"github.com/DmitriyVTitov/size"
)

type example struct {
	a []int
	b bool
	c int32
	d string
}

func main() {
	ex := example{
		a: []int{1, 2, 3}, // 24(slice itself)+8+8+8=24+24=48
		b: true,           // 1
		d: "1234",         // 16(string itself)+4=20
	} // 48+1+20+4=73+3(padding)
	fmt.Println("Размер в байтах для ex:", size.Of(ex))
	// Как получается результат?

	ex1 := example{
		a: []int{1, 2, 3}, // 24(slice itself)+8+8+8=24+24=48
		b: true,           // 1
		d: "1234",         // 16(string itself)+4=20
		c: 100,            //4
	} // 48+20+1+4=68+1+4=73+3(padding)=76
	fmt.Println("Размер в байтах для ex1:", size.Of(ex1))
	// Как получается результат?
}
