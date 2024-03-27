package main

import (
	"fmt"
	"strings"
)

func main() {

	x := [5]int{1, 2, 3, 4, 5}
	total := 0
	for _, value := range x {
		total += value
	}

	u := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(u[2:5]) // slice

	m := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	minim := m[1]
	position := 0
	for i, value := range m {
		if value < minim {
			position = i
			minim = value
		}
	}
	fmt.Printf("minimul se afla pe pozitia: %d si este: %d\n", position, minim)

	var helloWorld2 = "hello" + " " + "world"
	helloWorld2 += "!"

	fmt.Println(helloWorld2)

	fmt.Println(string(helloWorld2[0]))

	type Trouble struct {
		primaValoare, aDouaValoare int
	}
	mapa := make(map[string]Trouble)
	mapa["zero"] = Trouble{
		1, 2,
	}
	fmt.Println(mapa["zero"])

	mapa2 := make(map[string]string)
	mapa2["primul"] = "primul"
	mapa2["alDoilea"] = "doilea"

	fmt.Println(mapa2["primul"])
	fmt.Println(mapa2)

	sCheck := "hello world my name is gabi , i like apples and apples and apples my name is gabi"
	wordsCount := wordCount(sCheck)
	for word, count := range wordsCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func wordCount(s string) map[string]int {

	words := strings.Fields(s)
	//fmt.Println(words)
	//fmt.Println(len(words))
	m := make(map[string]int)

	for _, value := range words {
		if m[value] == 0 {
			m[value] = 0
		}
		m[value] += 1
	}
	return m
}

func swap(x, y string) (string, string) {
	return y, x
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func average(xs []float64) float64 {
	total := 0.
	for _, value := range xs {
		total += value
	}
	return total
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func add(x int, y int) int {
	return x + y
}
