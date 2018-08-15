package sorts

func AscSort(givenNumbers []int) []int {
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

func DescSort(givenNumbers []int) []int {
	length := len(givenNumbers) - 2
	for length >= 0 {
		number := givenNumbers[length]
		index := length + 1
		for index < len(givenNumbers) && number < givenNumbers[index] {
			givenNumbers[index - 1] = givenNumbers[index]
			givenNumbers[index] = number
			index = index + 1
		}
		length = length - 1
	}
	return givenNumbers
}