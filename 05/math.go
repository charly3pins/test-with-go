package math

// Sum will add up all the numbers of a slice and will return the total sum.
func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Add(a, b int) int {
	return a + b
}
