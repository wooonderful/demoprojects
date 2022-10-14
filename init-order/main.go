package main

import (
	"fmt"
	"github.com/wooonderful/demoprojects/init-order/pkg/pkg1"
)

const (
	constA = 1
)

var (
	varB = 2
)

func init() {
	fmt.Println("main.go's init function")
}

func main() {
	fmt.Println("it's main function")
	pkg1.Add()
}