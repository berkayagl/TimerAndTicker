package main

import (
	"fmt"
	"time"
)

func bgTicker(stop chan struct{}) {
	nt := time.NewTicker(1 * time.Second)
	count := 0

	for {
		select {
		case <-nt.C:
			count++
			fmt.Println("Go")

			if count >= 5 {
				close(stop)
				return
			}
		case <-stop:
			nt.Stop()
			return
		}
	}
}

func main() {
	fmt.Println("Go Tickers")

	fmt.Println("Uygulamamın geri kalanı devam edebilir...")

	stopTicker := make(chan struct{})
	go bgTicker(stopTicker)

	<-stopTicker
}
