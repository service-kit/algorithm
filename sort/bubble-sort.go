package sort

import "github.com/service-kit/algorithm/common"

func BubbleSort(list []common.Comparable, asc bool) {
	list_len := len(list)
	if 2 > list_len {
		return
	}
	for i := 0; i < list_len; i++ {
		for j := list_len - 1; j > i; j-- {
			if asc {
				if !list[j-1].Less(list[j]) {
					common.Swap(&list[j-1], &list[j])
				}
			} else {
				if list[j-1].Less(list[j]) {
					common.Swap(&list[j-1], &list[j])
				}
			}
		}
	}
}
