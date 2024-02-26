package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(time.Second * 5)

	go func() {
		<-timer.C
		fmt.Println("Zaman doldu!")
	}()

	stopped := timer.Stop()

	if stopped {
		fmt.Println("Zamanlayıcı durdu!")
	}

	time.Sleep(3 * time.Second)
}
