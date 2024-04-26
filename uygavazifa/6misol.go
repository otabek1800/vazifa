package main

import (
	"fmt"
	"time"
)
func dataAndTimeinpot(ch chan<- time.Time) {
	var date time.Time
	fmt.Println("Sana va vaqtni kiritish uchun (2004-12-04 12:01:01) formatda kiriting:")
	fmt.Scanln(&date)
	ch <- date 
}

func dataAndTimerecin(ch <-chan time.Time) time.Time {
	return <-ch 
}
func kalkulatTimeDiffer(start, end time.Time) time.Duration {
	return end.Sub(start) 
}
func main() {
	
	dateCh := make(chan time.Time)

	go dataAndTimeinpot(dateCh)
	start := time.Now() 

	end := dataAndTimerecin(dateCh)
	
	diff := kalkulatTimeDiffer(start, end)

	fmt.Println("Vaqt oraliqi:", diff)
}
