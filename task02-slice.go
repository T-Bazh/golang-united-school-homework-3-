package homework

func reverse(input []int64) (result []int64) {
	inputLen := len(input)
	result = make([]int64, inputLen)
	for index, item := range input {
		result[inputLen-index-1] = item
	}
	return result
}
