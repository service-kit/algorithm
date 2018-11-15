package sort

import "github.com/service-kit/algorithm/common"

func QuickSort(list []common.Comparable, asc bool) {
	quickSort(list, 0, len(list)-1, asc)
}

func quickSort(list []common.Comparable, low, high int, asc bool) {
	if low >= high {
		return
	}
	left := low
	right := high
	key := list[left]
	for left != right {
		for left < right {
			if asc {
				if !list[right].Less(key) {
					right--
					continue
				}
			} else {
				if list[right].Less(key) {
					right--
					continue
				}
			}
			break
		}
		list[left] = list[right]
		for left < right {
			if asc {
				if !key.Less(list[left]) {
					left++
					continue
				}
			} else {
				if !list[left].Less(key) {
					left++
					continue
				}
			}
			break
		}
		list[right] = list[left]
	}
	list[left] = key
	quickSort(list, low, left-1, asc)
	quickSort(list, left+1, high, asc)
}
