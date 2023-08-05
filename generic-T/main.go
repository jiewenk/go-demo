package main

import (
	"fmt"
	_case "generic-T/case"
)

func main() {
	fmt.Println("hello go generic")
	_case.PlainTypeCase()
	fmt.Println()
	_case.GenericFunctionCase()
	fmt.Println()
	_case.CustomGenericCase()
}
