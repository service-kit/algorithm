package sort

func CountingSort(list []int, asc bool) {
	list_len := len(list)
	counting_list := make([][]int, list_len)
	for _, val := range list {
		if asc {
			counting_list[val] = append(counting_list[val], val)
		} else {
			counting_list[list_len-val-1] = append(counting_list[list_len-val-1], val)
		}
	}
	list_index := 0
	for _, arr := range counting_list {
		for _, val := range arr {
			list[list_index] = val
			list_index++
		}
	}
}
