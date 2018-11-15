package sort

import (
	"github.com/service-kit/algorithm/common"
)

func InsertSort(list []common.Comparable, asc bool) {
	list_len := len(list)
	if 2 > list_len {
		return
	}
	for i := 1; i < list_len; i++ {
		for j := i - 1; j > -1; j-- {
			if asc {
				if !list[j].Less(list[j+1]) {
					common.Swap(&list[j], &list[j+1])
				}
			} else {
				if list[j].Less(list[j+1]) {
					common.Swap(&list[j], &list[j+1])
				}
			}
		}
	}
}
