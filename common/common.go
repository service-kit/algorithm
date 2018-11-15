package common

type Comparable interface {
	Less(other Comparable) bool
}

type Comparator func(v1, v2 interface{}) bool

func Swap(v1, v2 *Comparable) {
	var tmp Comparable = nil
	tmp = *v1
	*v1 = *v2
	*v2 = tmp
}

func CheckListIsSorted(list []Comparable, asc bool) bool {
	list_len := len(list)
	for i := 0; i < list_len-1; i++ {
		if asc && list[i+1].Less(list[i]) {
			return false
		}
		if !asc && list[i].Less(list[i+1]) {
			return false
		}
	}
	return true
}

func CheckIntListIsSorted(list []int, asc bool) bool {
	list_len := len(list)
	for i := 0; i < list_len-1; i++ {
		if asc && list[i+1] < list[i] {
			return false
		}
		if !asc && list[i] < list[i+1] {
			return false
		}
	}
	return true
}
