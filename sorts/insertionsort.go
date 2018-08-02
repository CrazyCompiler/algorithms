package sorts

func Sort(givenNumbers []int) []int {
	for index, number := range givenNumbers[1:]{
		j := index
		for j >= 0 && number < givenNumbers[j] {
			givenNumbers[j + 1] = givenNumbers[j]
			j = j - 1
		}
		givenNumbers[j+1] = number
	}
	return givenNumbers
}
