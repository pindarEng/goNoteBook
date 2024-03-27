package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	// Some new additions in golang 1.22 i tried
	for _ = range 10 {
		duration := rand.N(time.Minute * 60)
		fmt.Println(duration)
		time.Sleep(500 * time.Millisecond)
	}
	total := 0
	for i := range 5 {
		total += i
	}
	fmt.Println(total)
}
