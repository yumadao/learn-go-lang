package main

import (
	"fmt"
	"learn-go/calculator"
	"os"
	"unsafe"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Multiply(1, 2))

	var ui1 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1)
	var ui2 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui2)
	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("value of ui1(dereference): %v\n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)

	ok, result := true, "A"
	if ok {
		result := "B"
		println(result)
	} else {
		result := "C"
		println(result)
	}
	println(result)
}
