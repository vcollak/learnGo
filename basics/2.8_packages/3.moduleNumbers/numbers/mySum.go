package numbers

//Sum take a list of numbers and returns an addition
func Sum(numbers ...int) int {
	sum := 0
	for num := range numbers {
		sum = sum + num
	}
	return sum
}
