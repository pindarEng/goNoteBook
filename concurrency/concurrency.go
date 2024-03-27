package main

import "fmt"

func suma(s []int, c chan int) {
	suma := 0
	for _, v := range s {
		suma += v
	}
	c <- suma // send sum to c
}

func main() {

	s := []int{7, 2, 8, -9, 4, 0}
	cha := make(chan int)
	xch, ych := <-cha, <-cha // receive from c
	go suma(s[:len(s)/2], cha)
	fmt.Println(s[:len(s)/2])
	go suma(s[len(s)/2:], cha)
	fmt.Println(s[len(s)/2:])
	fmt.Println(xch, ych, xch+ych)

}
