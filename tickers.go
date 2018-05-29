package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(400 * time.Millisecond)
	go func() {
		for t := range ticker.C{
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1601 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
