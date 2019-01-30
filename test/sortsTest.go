package main

import (
	"Sorts_go/src"
	"fmt"
	"log"
)

var arr = []int{2, 34, 1, 56, 3, 24, 78, 12, 32, 9, 0, 7, 4, 11, 14, 32, 22, 21, 8, 23, 44, 42, 15, 18, 46, 22, 66, 75, 24, 21, 37, 13, 43, 68, 12, 16, 25, 33, 38, 21, 0}

var sortedArr = []int{0, 0, 1, 2, 3, 4, 7, 8, 9, 11, 12, 12, 13, 14, 15, 16, 18, 21, 21, 21, 22, 22, 23, 24, 24, 25, 32, 32, 33, 34, 37, 38, 42, 43, 44, 46, 56, 66, 68, 75, 78}

func main() {
	sortedArrStr := fmt.Sprintf("%v", sortedArr)

	sortStr := fmt.Sprintf("%v", src.BubbleSort(arr))
	assertEqual(sortStr, sortedArrStr, "BubbleSort")

	sortStr = fmt.Sprintf("%v", src.BucketSort(arr))
	assertEqual(sortStr, sortedArrStr, "BucketSort")

	sortStr = fmt.Sprintf("%v", src.CountingSort(arr, 78))
	assertEqual(sortStr, sortedArrStr, "CountingSort")

	sortStr = fmt.Sprintf("%v", src.HeapSort(arr))
	assertEqual(sortStr, sortedArrStr, "HeapSort")

	sortStr = fmt.Sprintf("%v", src.InsertionSort(arr))
	assertEqual(sortStr, sortedArrStr, "InsertionSort")

	sortStr = fmt.Sprintf("%v", src.QuickSort(arr, 0, len(arr)-1))
	assertEqual(sortStr, sortedArrStr, "QuickSort")

	sortStr = fmt.Sprintf("%v", src.RadixSort(arr, 100))
	assertEqual(sortStr, sortedArrStr, "RadixSort")

	sortStr = fmt.Sprintf("%v", src.SelectSort(arr))
	assertEqual(sortStr, sortedArrStr, "SelectSort")
}

func assertEqual(a interface{}, b interface{}, msg string) {
	if a == b {
		fmt.Printf("%s test passed!\n", msg)
		return
	}
	//fmt.Printf("%s test failed!", msg)
	log.Panicf("%s test failed!", msg)
}
