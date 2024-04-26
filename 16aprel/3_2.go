////
package main

import (
	"fmt"
	"time"
)

func main() {
	// Hozirgi vaqt
	now := time.Now()
	fmt.Println("hozirgi vaqt :", now)

	// 2-sini yaratamiz
	date := time.Date(2022, time.May, 15, 12, 30, 0, 0, time.UTC)
	fmt.Println("men yaratgan vaqt vaqt :", date)

	dif := date.Sub(now)

	days := int(dif.Hours() / 24)
	hours := int(dif.Hours()) % 24
	minutes := int(dif.Minutes()) % 60
	seconds := int(dif.Seconds()) % 60

	fmt.Println("kun :", days)
	fmt.Println("soat :", hours)
	fmt.Println("daqiqa :", minutes)
	fmt.Println("sekund :", seconds)
}
