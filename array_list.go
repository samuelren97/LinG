package ling

type ArrayList[T any] struct {
	count    int
	capacity int
	data     []T
}

func NewArrayList[T any](capacity int) *ArrayList[T] {
	return &ArrayList[T]{
		count:    0,
		capacity: capacity,
		data:     make([]T, capacity),
	}
}

func (l *ArrayList[T]) Count() int {
	return l.count
}

func (l *ArrayList[T]) Get(index int) T {
	return l.data[index]
}

func (l *ArrayList[T]) Push(element T) {
	if l.count == l.capacity {
		capacity := 1
		if l.capacity > 0 {
			capacity = l.capacity * 2
		}
		nData := make([]T, capacity)
		copy(nData, l.data)
		l.data = nData
	}

	l.data[l.count] = element
	l.count++
}

func (l *ArrayList[T]) Pop() T {
	lastElement := l.data[len(l.data)-1]
	l.RemoveAt(len(l.data) - 1)

	return lastElement
}

func (l *ArrayList[T]) RemoveAt(index int) {
	l.data = append(l.data[:index], l.data[index+1:]...)
	l.capacity = len(l.data)
	l.count = len(l.data)
}

func (l *ArrayList[T]) ForEach(f func(T)) {
	for _, v := range l.data {
		f(v)
	}
}

func (l *ArrayList[T]) Range() chan T {
	ch := make(chan T)
	go func() {
		for _, v := range l.data {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func (l *ArrayList[T]) Shift() T {
	firstElement := l.data[0]
	l.RemoveAt(0)
	return firstElement
}

func (l *ArrayList[T]) Sort(predicate func(T, T) bool) {
	if len(l.data) < 2 {
		return
	}

	for i := 0; i < l.count-1; i++ {
		for j := 0; j < l.count-i-1; j++ {
			if predicate(l.data[j], l.data[j+1]) {
				tmpItem := l.data[j]
				l.data[j] = l.data[j+1]
				l.data[j+1] = tmpItem
			}
		}
	}
}

func (l *ArrayList[T]) ToSorted(predicate func(T, T) bool) *ArrayList[T] {
	list := &ArrayList[T]{
		count:    len(l.data),
		capacity: len(l.data),
		data:     make([]T, len(l.data)),
	}
	copy(list.data, l.data)

	list.Sort(predicate)
	return list
}

func (l *ArrayList[T]) Where(predicate func(T) bool) *ArrayList[T] {
	list := &ArrayList[T]{
		count:    0,
		capacity: 0,
		data:     make([]T, 0),
	}

	for _, v := range l.data {
		if predicate(v) {
			list.Push(v)
		}
	}

	return list
}

func Map[T, U any](list *ArrayList[T], predicate func(T) U) *ArrayList[U] {
	nList := NewArrayList[U](len(list.data))
	nList.count = len(list.data)

	for i, v := range list.data {
		nList.data[i] = predicate(v)
	}

	return nList
}

func Reduce[T, U any](list *ArrayList[T], predicate func(U, T, int) U, init U) U {
	var acc U = init

	for i, v := range list.data {
		acc = predicate(acc, v, i)
	}

	return acc
}
