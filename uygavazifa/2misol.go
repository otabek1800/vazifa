package main

import (
	"fmt"
)

var uzbStr = map[int]string{
	0:  "nol",
	1:  "bir",
	2:  "ikki",
	3:  "uch",
	4:  "to'rt",
	5:  "besh",
	6:  "olti",
	7:  "yetti",
	8:  "sakkiz",
	9:  "to'qqiz",
}
var onliklar = map[int]string{
	1: "o'n",
	2: "yigirma",
	3: "o'ttiz",
	4: "qirq",
	5: "ellik",
	6: "oltmish",
	7: "yetmish",
	8: "sakson",
	9: "to'qson",
}
var mingliklar = map[int]string{
	1: "ming",
	2: "milyon",
	3: "milliard",
	4: "trillion",
}
func main() {
	var number int
	fmt.Print("son kiriting: ")
	fmt.Scanln(&number)

	var parts []string

	toonsend := number / 1000
	if toonsend > 0 {
		parts = append(parts, uzbStr[toonsend], mingliklar[1])
	}

	hundred, ten, one := (number/100)%10, (number/10)%10, number%10

	if hundred > 0 {
		parts = append(parts, uzbStr[hundred], "yuz")
	}
	if ten > 0 {
		parts = append(parts, onliklar[ten])
	}
	if one > 0 || (one == 0 && ten == 0 && hundred == 0 && toonsend == 0) {
		parts = append(parts, uzbStr[one])
	}
	result := ""
	for _, part := range parts {
		result += part + " "
	}
	fmt.Println("oqilgan holatdagi son:", result)
}
