package main

import (
	"fmt"
)
func findMaxnum(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	max := slice[0] 
	for _, num := range slice {
		if num > max {
			max = num 
		}	
	}
	return max
}	
func main () {
	slc := []int{10, 24, 209, 6, 35,100, 14, 29}
	max := findMaxnum(slc)
	fmt.Println("eng katta slc:", max)
}


