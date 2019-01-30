package src

func CountingSort(arr []int, max int) []int {
	bucket := make([]int, max+1)
	sortedIndex := 0

	for _, v := range arr {
		bucket[v]++
	}

	for k, v := range bucket {
		times := v
		for times > 0 {
			arr[sortedIndex] = k
			sortedIndex++
			times--
		}
	}
	return arr
}
