package main

import (
	"fmt"
	"math"
)

func main() {
	num := []int32{10, 20, 7}
	fmt.Println(SumMin(num, 4))
}

// SumMin
func SumMin(num []int32, k int32) int32 {
	var idx int32
	var div float64
	var idx2 int32
	var n int32

	if len(num) == 0 || len(num) > int(math.Pow(10, 5)) {
		return 0
	}
	if k < 0 || k > int32(math.Pow(10, 7)) {
		return 0
	}

	for idx = 0; idx <= k-1; idx++ {
		if int(idx) > len(num)-1 {
			idx2 = 0
		} else {
			idx2 = idx
		}
		div = float64(num[idx2] / 2)
		div = math.Ceil(div)
		num[idx2] = int32(div)

	}
	n = Sum(num)
	return n
}

// SumMin
func Sum(num []int32) int32 {
	var count int32
	for i := 0; i < len(num); i++ {
		count += num[i]
	}
	return count
}
