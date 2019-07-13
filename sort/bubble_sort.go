package sort

//BubbleSort sort the input array using bubble sort algorithm
func BubbleSort(input []int64) []int64 {
	if len(input) < 2 {
		return input
	}
	for i := 0; i < len(input)-1; i++ {
		for index := 0; index < len(input)-1; index++ {
			if input[index] > input[index+1] {
				input[index], input[index+1] = input[index+1], input[index]
			}
		}
	}
	return input
}
