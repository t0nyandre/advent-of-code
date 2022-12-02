package utils

// Sum function sums together an array of integers
func Sum(array []int) int {
	sum := 0
	for _, i := range array {
		sum += i
	}
	return sum
}
