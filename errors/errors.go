package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fsqrt, err := Sqrt(64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fsqrt)
	}
	gsqrt, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gsqrt)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("NU POTI FACE SQRT LA NEGATIVE")
	}
	return math.Sqrt(f), nil // return nil adica nu avem eroare
}
