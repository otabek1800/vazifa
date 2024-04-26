package main

import (
	"fmt"
	"strconv"
)
func main () {
	var input string
	fmt.Print("Qiymat kiriting:")
	fmt.Scanln(&input)

	
	
	if num, err := strconv.ParseInt(input, 10, 64); err == nil {
		fmt.Printf("Kiritilgan qiymat int: %d\n", num)
	}else if boolean, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("Kiritilgan qiymat bool: %t\n", boolean)
	} else {
		fmt.Printf("Kiritilgan qiymat string: %s\n", input)
	}
}

