package main

import (
	"math/rand"
	"time"
)

//Rounding function
func roundingFunction(roundingTax float64) float64 {
	//2nd Approach -  math.Ceil(roundingTax*20) / 20
	return (float64(int64(20*roundingTax+0.5)) / 20)
}

//Add function
func addFunction(number1 float64, number2 float64) float64 {
	return ((number1 * 100) + (number2 * 100)) / 100
}

//Random number generator function
func randomGenerator() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(len(itemTypes) - 1)
}
