package Slices

func Sum(numbers []int) int {

	var sum int
	sum = 0
	for _,number := range numbers{
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int{

	length := len(numbers)
	sums := make([]int,length)
	for i,number := range numbers{
		sums[i] = Sum(number)
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {

	length := len(numbers)
	sums := make([]int,length)
	for i,number := range numbers{
		if len(number) < 1 {
			sums[i] = 0
		}else {
			sums[i] = Sum(number[1:])
		}
	}
	return sums
}
