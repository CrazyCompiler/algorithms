package peak_finding

type peak struct {
	index int
	value int
}

func FindAPeak(numbers []int) (int, int) {
	for index, number := range numbers{
		if len(numbers) == 1{
			return index, number
		}
 		if index < 1 && numbers[index + 1] <= number{
			return index, number
		}
		if index == len(numbers) - 1 && numbers[index - 1] <= number{
			return index, number
		}
		if index > 0 && index < len(numbers) && numbers[index - 1] <= number && numbers[index + 1] <= number{
			return index, number
		}
	}
	return -1, -1
}

func FindAllPeaks(numbers []int) []peak {
	var allPeaks []peak
	for index, number := range numbers{
		if len(numbers) == 1{
			allPeaks = append(allPeaks,peak{index, number})
		}else if index < 1 && numbers[index + 1] <= number{
			allPeaks = append(allPeaks,peak{index, number})
		}else if index == len(numbers) - 1 && numbers[index - 1] <= number{
			allPeaks = append(allPeaks,peak{index, number})
		}else if index > 0 && index < len(numbers) && numbers[index - 1] <= number && numbers[index + 1] <= number{
			allPeaks = append(allPeaks,peak{index, number})
		}
	}
	return allPeaks
}
