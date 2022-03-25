package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(arrNumbers ...[]int) []int {
	res := make([]int, len(arrNumbers))
	for i, arrNumber := range arrNumbers {
		res[i] = Sum(arrNumber)
	}
	return res
}

func SumAllTail(arrNumbers ...[]int) []int {
	res := make([]int, len(arrNumbers))
	for i, arrNumber := range arrNumbers {
		if len(arrNumber) == 0 {
			res[i] = 0
			continue
		}
		res[i] = Sum(arrNumber[len(arrNumber)-1:])
	}
	return res
}
