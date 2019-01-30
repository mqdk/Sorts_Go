package src

var length int

func HeapSort(arr []int) []int {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		swap(arr, 0, i)
		length--
		heapify(arr, 0)
	}
	return arr
}

func buildMaxHeap(arr []int) {
	length = len(arr)
	for i := length/2 - 1; i >= 0; i-- {
		heapify(arr, i)
	}
}

func heapify(arr []int, i int) {
	left := i*2 + 1
	right := i*2 + 2
	parent := i

	if left < length && arr[left] > arr[parent] {
		parent = left
	}
	if right < length && arr[right] > arr[parent] {
		parent = right
	}

	if parent != i {
		swap(arr, i, parent)
		heapify(arr, parent)
	}
}

//func swap(arr []int, i, j int) {
//	temp := arr[i]
//	arr[i] = arr[j]
//	arr[j] = temp
//}
