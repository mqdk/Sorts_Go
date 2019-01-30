package src

import (
	"math"
	"sort"
)

func BucketSort(arr []int) []int {
	length := len(arr)
	if length == 0 {
		return arr
	}

	type bucket struct {
		list []int
	}

	maxVal := math.MinInt32
	minVal := math.MaxUint32
	for _, v := range arr {
		maxVal = Max(maxVal, v)
		minVal = Min(minVal, v)
	}
	bucketCount := (maxVal-minVal)/length + 1
	buckets := make([]bucket, bucketCount)

	for _, v := range arr {
		num := (v - minVal) / length
		buckets[num].list = append(buckets[num].list, v) //必须亲自用buckets[num].list进行操作，这样才能改变它的内部引用
	}

	var index = 0
	for _, b := range buckets {
		sort.Ints(b.list)
		for _, v := range b.list {
			arr[index] = v
			index++
		}
	}
	return arr
}

//go的Min Max不支持int，且不支持三元表达式
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
