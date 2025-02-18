package main

import (
	"fmt"
	"math"
)

func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main(){
	fmt.Println(Pow(3, 2, 10))
	fmt.Println(Pow(3, 3, 20))
}