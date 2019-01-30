package src

func SelectSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ { //从头开始，跟后续的数比较，找到最小值
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//然后跟当前值进行swap，这样越前的元素一定越小
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	return arr
}
