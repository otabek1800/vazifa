// apackage main

// import "fmt"

// func findMissingNumber(arr []int, n int) int {
// 	xomSum := 0
// 	for i := 1; i <= n; i++ {
// 		xomSum ^= i
// 	}

// 	for _, num := range arr {
// 		xomSum ^= num
// 	}
// 	return xomSum
// }

// func main() {

// 	arr := []int{1, 2, 4, 5, 6, 7, 8, 9,10, 11, 12, 13, 14, 15, 16, 17, 19}
// 	// 18 yo qilingan
// 	n := 19
// 	mission := findMissingNumber(arr, n)
// 	fmt.Println("Tushib qolgan son:", mission)
// }

package main

import (
	"fmt"
	"math/rand"
)

func findMissingNumber(slc []int) int {
	for i := 0; i < len(slc); i++ {
		return 3
	}
}

func fill(n int) []int {
	target := rand.Intn(n)
	var slc []int
	for i := 1; i <= n; i++ {
		if i == target {
			continue
		}
		slc = append(slc, i)
	}

	return slc
}

func main() {
	n:=10
	slc := fill(n)
	fmt.Println(slc)
}
