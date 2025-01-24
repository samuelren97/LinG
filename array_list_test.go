package ling_test

import (
	"testing"

	ling "github.com/samuelren97/LinG"
)

func createNewList(nbElements int) *ling.ArrayList[int] {
	list := ling.NewArrayList[int](nbElements)
	for i := 0; i < nbElements; i++ {
		list.Push(i)
	}
	return list
}

func createIdenticalListAndArray(nbElements int) (*ling.ArrayList[int], []int) {
	list := ling.NewArrayList[int](nbElements)
	arr := make([]int, nbElements)

	for i := 0; i < nbElements; i++ {
		list.Push(i)
		arr[i] = i
	}

	return list, arr
}

func validateList(expected []int, list *ling.ArrayList[int], t *testing.T) {
	if len(expected) != list.Count() {
		t.Errorf("unmatching list lengths")
		return
	}

	for i, v := range expected {
		val := list.Get(i)
		if val != v {
			t.Errorf("list is incorrect [index: %d], wanted: %d, got: %d", i, v, val)
		}
	}
}

func TestGet_ThreeValueList(t *testing.T) {
	list := createNewList(3)
	expected := 1

	got := list.Get(1)
	if got != expected {
		t.Errorf("wanted: %d, got: %d", expected, got)
	}
}

func TestPush_EmptyListCapacityOnePushOne(t *testing.T) {
	list := ling.NewArrayList[int](1)
	expected := [1]int{1}

	list.Push(1)

	if list.Count() != len(expected) {
		t.Errorf("list is supposed to have a length of %d", len(expected))
	}

	got := list.Get(0)
	if got != expected[0] {
		t.Errorf("wanted: %d, got: %d", expected[0], got)
	}
}

func TestPush_EmptyListCapacityOnePushTwo(t *testing.T) {
	list := ling.NewArrayList[int](1)
	expected := [2]int{0, 1}

	list.Push(0)
	list.Push(1)

	if list.Count() != len(expected) {
		t.Errorf("list is supposed to have a length of %d", len(expected))
	}

	for i, v := range expected {
		val := list.Get(i)
		if val != v {
			t.Errorf("list is incorrect [index: %d], wanted: %d, got: %d", i, v, val)
		}
	}
}

func TestPush_EmptyListCapacityTwoPushTwo(t *testing.T) {
	list := ling.NewArrayList[int](2)
	expected := [2]int{0, 1}

	list.Push(0)
	list.Push(1)

	validateList(expected[:], list, t)
}

func TestPop_OneValue(t *testing.T) {
	list := createNewList(1)
	expected := [0]int{}
	expectedOut := 0

	got := list.Pop()
	if got != expectedOut {
		t.Errorf("wanted: %d, got: %d", expectedOut, got)
	}

	validateList(expected[:], list, t)
}

func TestPop_ThreeValues(t *testing.T) {
	list := createNewList(3)
	expected := [2]int{0, 1}
	expectedOut := 2

	got := list.Pop()
	if got != expectedOut {
		t.Errorf("wanted: %d, got: %d", expectedOut, got)
	}

	validateList(expected[:], list, t)
}

func TestRemoveAt_FirstElement(t *testing.T) {
	list := createNewList(3)
	expected := [2]int{1, 2}

	list.RemoveAt(0)

	validateList(expected[:], list, t)
}

func TestRemoveAt_MiddleElement(t *testing.T) {
	list := createNewList(3)
	expected := [2]int{0, 2}

	list.RemoveAt(1)

	validateList(expected[:], list, t)
}

func TestRemoveAt_LastElement(t *testing.T) {
	list := createNewList(3)
	expected := [2]int{0, 1}

	list.RemoveAt(2)

	validateList(expected[:], list, t)
}

func TestForEach_ThreeValues(t *testing.T) {
	list, expected := createIdenticalListAndArray(3)

	index := 0
	list.ForEach(func(i int) {
		if i != expected[index] {
			t.Errorf("list is incorrect [index: %d], wanted: %d, got: %d", index, expected[index], i)
		}
		index++
	})
}

func TestRange_ThreeValues(t *testing.T) {
	list, expected := createIdenticalListAndArray(3)

	index := 0
	for v := range list.Range() {
		if v != expected[index] {
			t.Errorf("list is incorrect [index: %d], wanted: %d, got: %d", index, expected[index], v)
		}
		index++
	}
}

func TestShift_OneValue(t *testing.T) {
	list := createNewList(1)
	expectedLength := 0
	wanted := 0

	if got := list.Shift(); got != wanted {
		t.Errorf("wanted: %d, got: %d", wanted, got)
		return
	}

	if list.Count() != expectedLength {
		t.Errorf("incorrect list length, wanted: %d, got: %d", expectedLength, list.Count())
	}
}

func TestShift_ThreeValues(t *testing.T) {
	list := createNewList(3)
	expectedList := [2]int{1, 2}
	wanted := 0

	if got := list.Shift(); got != wanted {
		t.Errorf("wanted: %d, got: %d", wanted, got)
		return
	}

	validateList(expectedList[:], list, t)
}

func TestSort_Ints(t *testing.T) {
	list := createNewList(3)
	wanted := [3]int{2, 1, 0}

	list.Sort(func(a, b int) bool {
		return a < b
	})

	validateList(wanted[:], list, t)
}

func TestToSorted_Ints(t *testing.T) {
	list := createNewList(3)
	wantedOriginList := [3]int{0, 1, 2}
	wanted := [3]int{2, 1, 0}

	sortedList := list.ToSorted(func(a, b int) bool {
		return a < b
	})

	validateList(wantedOriginList[:], list, t)
	validateList(wanted[:], sortedList, t)
}

func TestMap(t *testing.T) {
	list := createNewList(3)
	expected := [3]int{1, 2, 3}

	resultList := ling.Map(list, func(i int) int {
		return i + 1
	})

	validateList(expected[:], resultList, t)
}

func TestReduce_IntListToSum(t *testing.T) {
	list := createNewList(3)
	expected := 3

	got := ling.Reduce(list, func(acc int, value int, index int) int {
		return acc + value
	}, 0)

	if got != expected {
		t.Errorf("wanted: %d, got: %d", expected, got)
	}
}

func TestReduce_ConcatStringWithNotEmptyInitializer(t *testing.T) {
	list := ling.NewArrayList[string](2)
	list.Push("Hello ")
	list.Push("World!")
	prefix := "Wow"
	wanted := prefix + "Hello World!"

	got := ling.Reduce(list, func(acc string, value string, index int) string {
		return acc + value
	}, prefix)

	if got != wanted {
		t.Errorf("wanted: %s, got: %s", wanted, got)
	}
}

func TestWhere_OneResult(t *testing.T) {
	list := createNewList(3)
	wanted := [1]int{1}

	got := list.Where(func(i int) bool {
		return i == 1
	})

	validateList(wanted[:], got, t)
}
