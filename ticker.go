package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Go Tickers")
	ticker := time.NewTicker(1 * time.Second)

	for tc := range ticker.C {
		fmt.Println("Golang ", tc)
	}
}
