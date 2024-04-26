package main

import (
	"fmt"
)

func combeExtrasum(extrsum []string) int {
	x := 0 
	for _, expr := range extrsum {
		switch expr {
		case "++x", "x++":
			x++
		case "--x", "x--":
			x--
		}
	}
	return x
}
func main() {
	input := []string{"--x","x--", "++x","--x","++x"}
	result := combeExtrasum(input)
	fmt.Println("Natija:", result)
}



