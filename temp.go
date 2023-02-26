package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(primeNumber(13))
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "akuma", "steve", "geese"}))
}

func primeNumber(a int) bool {
	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			isPrime = false
		}
	}
	return isPrime
}

func pow(x, n int) int {
	if n == 1 {
		return x
	} else if n == 0 {
		return 1
	} else if n%2 == 0 {
		return pow(x, n/2) * pow(x, n/2)
	} else {
		return x * pow(x, n/2) * pow(x, n/2)
	}
}

func ArrayMerge(arrayA, arrayB []string) []string {
	arrayC := make([]string, len(arrayA))
	copy(arrayC, arrayA)
	for _, arr := range arrayB {
		isSame := false
		for _, arr1 := range arrayA {
			if arr == arr1 {
				isSame = true
			}
		}
		if isSame == false {
			arrayC = append(arrayC, arr)
		}
	}
	return arrayC
}
