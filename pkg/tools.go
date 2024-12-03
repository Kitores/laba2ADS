package pkg

import "math/rand"

func RandArray(size int, seed int64) []int {
	rand.Seed(seed)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10000)
	}
	return arr
}
func IncreasingArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i + 10
	}
	return arr
}
