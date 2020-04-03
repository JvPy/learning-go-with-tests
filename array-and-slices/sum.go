package slices

//Sum - Sum and returns the value
func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}

//SumAll = Sum slices and return a new slice with the sums
func SumAll(numbersToSum ...[]int) (sumSlice []int) {

	for _, numbers := range numbersToSum {
		sumSlice = append(sumSlice, Sum(numbers))
	}

	return
}

//SumAllTails = Sum the last element of slices and return a new slice with the sums
func SumAllTails(numbersToSum ...[]int) (sumSliceTails []int) {

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumSliceTails = append(sumSliceTails, 0)
		} else {
			tail := numbers[1:]
			sumSliceTails = append(sumSliceTails, Sum(tail))
		}
	}

	return
}
