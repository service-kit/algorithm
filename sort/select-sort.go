package sort

import "github.com/service-kit/algorithm/common"

func SelectSort(list []common.Comparable, asc bool) {
	list_len := len(list)
	if 2 > list_len {
		return
	}
	for i := 0; i < list_len; i++ {
		min := list[i]
		for j := i + 1; j < list_len; j++ {
			if asc {
				if list[j].Less(min) {
					common.Swap(&list[j], &min)
				}
			} else {
				if !list[j].Less(min) {
					common.Swap(&list[j], &min)
				}
			}
		}
		list[i] = min
	}
}
