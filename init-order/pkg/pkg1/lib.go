package pkg1

import "fmt"

const (
	constC = 3
)

var (
	varD = 4
)

func init() {
	fmt.Println("pkg1's init function")
}

func Add() {
	fmt.Printf("it's pkg1 add function, result = %d \n", constC+varD)
}
