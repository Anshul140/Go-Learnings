package arraysandslices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfInput := len(numbersToSum)

	// define empty array of size of input 2d array
	sums := make([]int, lengthOfInput)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

/*
Notes:
inbuilt range method
- range returns two values - the index and the value
- in above method, we are ignoring index by using _
- '_' -> underscore is blank identifier


variadic functions:
- functions that can be called with any number of trailing arguments.
*/
