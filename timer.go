package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(5 * time.Second)

	done := make(chan bool)

	go func() {
		<-timer.C
		fmt.Println("Zaman doldu!")
		done <- true
	}()

	<-done
}
