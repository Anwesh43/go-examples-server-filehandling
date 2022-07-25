package main 

import (
	"fmt"
)
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from a problem", r)
		} else {
			fmt.Println("program was successful")
		}
	}()
	a := 10
	b := 0
	res := a / b 
	fmt.Println("may not work", res)
}