package main

import (
	"fmt"
	"time"
)

func main() {

	nt := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-nt.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	nt.Stop()
	done <- true
	fmt.Println("Ticker durdu!")
}
