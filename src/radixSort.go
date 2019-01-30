package src

func RadixSort(arr []int, maxDigit int) []int {
	d := 1
	k := 0
	length := len(arr)
	bucket := make([][]int, 10)
	for i := range bucket { //初始化二维数组
		bucket[i] = make([]int, length)
	}
	order := make([]int, length)
	for d < maxDigit {
		for _, v := range arr {
			digit := (v / d) % 10
			bucket[digit][order[digit]] = v
			order[digit]++
		}

		for i := 0; i < len(bucket); i++ {
			if order[i] != 0 {
				for j := 0; j < order[i]; j++ {
					arr[k] = bucket[i][j]
					k++
				}
			}
			order[i] = 0
		}
		d *= 10
		k = 0
	}
	return arr
}
