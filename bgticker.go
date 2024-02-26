package main

import (
	"fmt"
	"time"
)

func bgTicker() {
	nt := time.NewTicker(1 * time.Second)
	for _ = range nt.C {
		fmt.Println("Go")
	}
}

func main() {
	fmt.Println("Go Tickers")

	fmt.Println("Uygulamamın geri kalanı devam edebilir...")

	go bgTicker()

	select {}

}
