package utils

// Sum function sums together an array of integers
func Sum(array []int) int {
	sum := 0
	for _, i := range array {
		sum += i
	}
	return sum
}

func Contains(array []string, i string) bool {
	for _, a := range array {
		if a == i {
			return true
		}
	}
	return false
}
