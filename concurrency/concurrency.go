package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(i) // adaugam i numar de goroutines la wait group // iteratii: pt 0 adaugam 0 la goroutine pt 1 adaugam 1 si tot asa
		go func(n int) {
			defer wg.Done() // decrementeaza numarul din wait group cu cate unu
			somerandomFunction(i)
		}(i) // ptr fiecare iteratie lansam un nou goroutine si o apelam imediat (i).
	}

	s := []int{7, 2, 8, -9, 4, 0}
	fmt.Println(s[:len(s)/2])
	fmt.Println(s[len(s)/2:])

	cha := make(chan int, 2) // buffered channel
	go suma(s[:len(s)/2], cha)
	go suma(s[len(s)/2:], cha)
	wg.Wait() // we wait for somerandomFunctiont to be done

	xch, ych := <-cha, <-cha // receive from c
	fmt.Println(xch, ych, xch+ych)

}

func somerandomFunction(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func suma(s []int, c chan int) {
	suma := 0
	for _, v := range s {
		suma += v
	}
	c <- suma // send sum to c
}
