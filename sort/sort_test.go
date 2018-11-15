package sort

import (
	"fmt"
	"github.com/service-kit/algorithm/common"
	"math/rand"
	"testing"
	"time"
)

type ComparableItem struct {
	Val int
}

func (ci *ComparableItem) Less(other common.Comparable) bool {
	return ci.Val < other.(*ComparableItem).Val
}

func getNoOrderList(count, max int) []common.Comparable {
	var cl []common.Comparable = make([]common.Comparable, count)
	for i := 0; i < count; i++ {
		item := new(ComparableItem)
		item.Val = rand.Int() % max
		cl[i] = item
	}
	return cl
}

func getNoOrderIntList(count, max int) []int {
	cl := make([]int, count)
	for i := 0; i < count; i++ {
		cl[i] = rand.Int() % max
	}
	return cl
}

func getMilSec() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func testBubbleSort(count int, asc bool) {
	cl := getNoOrderList(count, count)
	BubbleSort(cl, asc)
	res := common.CheckListIsSorted(cl, asc)
	fmt.Println("check bubblesort res ", res)
}

func testInsertSort(count int, asc bool) {
	cl := getNoOrderList(count, count)
	InsertSort(cl, asc)
	res := common.CheckListIsSorted(cl, asc)
	fmt.Println("check insertsort res ", res)
}

func testShellSort(count int, asc bool) {
	cl := getNoOrderList(count, count)
	ShellSort(cl, asc)
	res := common.CheckListIsSorted(cl, asc)
	fmt.Println("check shellsort res ", res)
}

func testSelectSort(count int, asc bool) {
	cl := getNoOrderList(count, count)
	SelectSort(cl, asc)
	res := common.CheckListIsSorted(cl, asc)
	fmt.Println("check selectsort res ", res)
}

func testQuickSort(count int, asc bool) {
	cl := getNoOrderList(count, count)
	QuickSort(cl, asc)
	res := common.CheckListIsSorted(cl, asc)
	fmt.Println("check quicksort res ", res)
}

func testMergeSort(count int, asc bool) {
	cl := getNoOrderList(count, count)
	MergeSort(cl, asc)
	res := common.CheckListIsSorted(cl, asc)
	fmt.Println("check mergesort res ", res)
}

func testCountingSort(count int, asc bool) {
	cl := getNoOrderIntList(count, count)
	CountingSort(cl, asc)
	res := common.CheckIntListIsSorted(cl, asc)
	fmt.Println("check countingsort res ", res)
}

func testHeapSort(count int, asc bool) {
	cl := getNoOrderList(count, count)
	HeapSort(cl, asc)
	res := common.CheckListIsSorted(cl[1:], asc)
	fmt.Println("check heapsort res ", res)
}

func TestAllSort(t *testing.T) {
	count := 10000
	for _, val := range []bool{true, false} {
		beginTime := getMilSec()
		testBubbleSort(count, val)
		bubbleTime := getMilSec()
		testInsertSort(count, val)
		insertTime := getMilSec()
		testShellSort(count, val)
		shellTime := getMilSec()
		testSelectSort(count, val)
		selectTime := getMilSec()
		testQuickSort(count, val)
		quickTime := getMilSec()
		testMergeSort(count, val)
		mergeTime := getMilSec()
		testCountingSort(count, val)
		countingTime := getMilSec()
		testHeapSort(count, val)
		heapTime := getMilSec()
		fmt.Printf("sort time list:\n"+
			"bubblesort time:%vms\n"+
			"insertsort time:%vms\n"+
			"shellsort time:%vms\n"+
			"selectsort time:%vms\n"+
			"quicksort time:%vms\n"+
			"mergesort time:%vms\n"+
			"countingsort time:%vms\n"+
			"heapsort time:%vms\n",
			bubbleTime-beginTime,
			insertTime-bubbleTime,
			shellTime-insertTime,
			selectTime-shellTime,
			quickTime-selectTime,
			mergeTime-quickTime,
			countingTime-mergeTime,
			heapTime-countingTime)
	}
	t.Log("test sort pass")
}
