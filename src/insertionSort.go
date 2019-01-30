package src

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		currentVal := arr[i]
		for preIndex >= 0 && arr[preIndex] > currentVal { //比较元素currentVal比当前元素要小
			arr[preIndex+1] = arr[preIndex] //当前元素后移一位
			preIndex--                      //从比较元素位置往前遍历
		}
		arr[preIndex+1] = currentVal //比较元素更大 则插入
	}
	return arr
}
