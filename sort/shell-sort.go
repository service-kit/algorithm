package sort

import "github.com/service-kit/algorithm/common"

func ShellSort(list []common.Comparable, asc bool) {
	list_len := len(list)
	if 2 > list_len {
		return
	}
	gap := list_len / 2
	for gap >= 1 {
		for i := gap; i < list_len; i++ {
			for j := i - gap; j > -1; j -= gap {
				if asc {
					if !list[j].Less(list[j+gap]) {
						common.Swap(&list[j], &list[j+gap])
					}
				} else {
					if list[j].Less(list[j+gap]) {
						common.Swap(&list[j], &list[j+gap])
					}
				}
			}
		}
		gap = gap / 2
	}
}
