package sort

import "github.com/service-kit/algorithm/common"

func sink(list []common.Comparable, k, N int, asc bool) {
	for {
		i := 2 * k
		if i > N {
			break
		}
		if asc {
			if i < N && list[i].Less(list[i+1]) {
				i++
			}
			if !list[k].Less(list[i]) {
				break
			}
		} else {
			if i < N && !list[i].Less(list[i+1]) {
				i++
			}
			if list[k].Less(list[i]) {
				break
			}
		}
		common.Swap(&list[k], &list[i])
		k = i
	}
}

func HeapSort(list []common.Comparable, asc bool) {
	N := len(list)
	fake_list := make([]common.Comparable, N+1)
	copy(fake_list[1:], list)
	for k := N / 2; k >= 1; k-- {
		sink(fake_list, k, N, asc)
	}
	for N > 1 {
		common.Swap(&fake_list[1], &fake_list[N])
		N--
		sink(fake_list, 1, N, asc)
	}
	copy(list, fake_list[1:])
}
