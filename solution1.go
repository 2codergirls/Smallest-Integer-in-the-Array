package kata

func SmallestIntegerFinder(numbers []int) int {
	smallest := numbers[0]
	for _, number := range numbers {
		if smallest > number {
			smallest = number
		}
	}
	return smallest
}
