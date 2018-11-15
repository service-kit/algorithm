package sort

import (
	"github.com/service-kit/algorithm/common"
)

func MergeSort(list []common.Comparable, asc bool) {
	list_len := len(list)
	if 2 > list_len {
		return
	}
	copy(list, mergeSort(list, asc))
}

func mergeSort(list []common.Comparable, asc bool) []common.Comparable {
	list_len := len(list)
	if 2 > list_len {
		return list
	}
	mid := list_len / 2
	left := list[:mid]
	right := list[mid:]
	return merge(mergeSort(left, asc), mergeSort(right, asc), asc)
}

func merge(left, right []common.Comparable, asc bool) []common.Comparable {
	leftLen := len(left)
	rightLen := len(right)
	resLen := leftLen + rightLen
	resIndex := 0
	leftIndex := 0
	rightIndex := 0
	res := make([]common.Comparable, resLen)
	for leftLen > leftIndex && rightLen > rightIndex {
		if asc {
			if !right[rightIndex].Less(left[leftIndex]) {
				res[resIndex] = left[leftIndex]
				leftIndex++
			} else {
				res[resIndex] = right[rightIndex]
				rightIndex++
			}
		} else {
			if !left[leftIndex].Less(right[rightIndex]) {
				res[resIndex] = left[leftIndex]
				leftIndex++
			} else {
				res[resIndex] = right[rightIndex]
				rightIndex++
			}
		}
		resIndex++
	}
	for leftIndex < leftLen {
		res[resIndex] = left[leftIndex]
		resIndex++
		leftIndex++
	}
	for rightIndex < rightLen {
		res[resIndex] = right[rightIndex]
		resIndex++
		rightIndex++
	}
	return res
}
