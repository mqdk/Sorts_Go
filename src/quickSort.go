package src

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		middle := getMiddle(arr, low, high)
		QuickSort(arr, low, middle-1)  //递归排序左边
		QuickSort(arr, middle+1, high) //递归排序右边
	}
	return arr
}

func getMiddle(arr []int, low, high int) int {
	key := arr[high]
	for j := low; j < high; j++ { //从前往后搜索
		if arr[j] <= key { //找到比基准值小的元素，swap到前面
			swap(arr, low, j)
			low++
		}
	}
	swap(arr, low, high) //一轮下来 基准值置于中间（并非正中那种概念）
	return low
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
