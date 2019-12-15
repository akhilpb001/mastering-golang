package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe              = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var (
	num          = 100
	str          = "hello"
	isBool       = true
	decimal      = 10.2435
	char    byte = 'A'
	unicode rune = '\u2318'
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Var: num; Type: %T; Value: %v\n", num, num)
	fmt.Printf("Var: str; Type: %T; Value: %v\n", str, str)
	fmt.Printf("Var: isBool; Type: %T; Value: %v\n", isBool, isBool)
	fmt.Printf("Var: decimal; Type: %T; Value: %v\n", decimal, decimal)
	fmt.Printf("Var: char; Type: %T; Value: %v\n", char, char)
	fmt.Printf("Var: unicode; Type: %T; Value: %v\n", unicode, unicode)
}
