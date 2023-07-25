package gotk

import "sync"

type ArrayPredicate[T comparable] func(value T) bool
type ArrayCallback[T comparable] func(value T, index int)

type Array[T comparable] interface {
	First() T
	Last() T
	At(i int) T
	Get(i int) T
	Take(i int) T
	Push(v ...T) int
	Pop() T
	Slice(start int, stop ...int) Array[T]
	Splice(offset int, length int, v ...T) Array[T]
	Contains(v T) bool
	IndexOf(v T) int
	Length() int
	Find(fn ArrayPredicate[T]) T
	Filter(fn ArrayPredicate[T]) []T
	ForEach(fn ArrayCallback[T])
}

// ============================

type array[T comparable] struct {
	mutex sync.Mutex
	data  []T
}

func (a *array[T]) First() T {
	return a.data[0]
}

func (a *array[T]) Last() T {
	return a.data[a.Length()-1]
}

func (a *array[T]) At(i int) T {
	return a.data[i]
}

func (a *array[T]) Get(i int) T {
	return a.At(i)
}

func (a *array[T]) Take(i int) T {
	value := a.At(i)
	a.Splice(i, 1)
	return value
}

func (a *array[T]) Push(v ...T) int {
	a.data = append(a.data, v...)
	return a.Length() - 1
}

func (a *array[T]) Pop() T {
	var value T

	lastIndex := len(a.data)

	if lastIndex > 0 {
		lastIndex -= 1

		value = a.data[lastIndex]
		a.data = a.data[:lastIndex]
	}

	return value
}

func (a *array[T]) Slice(start int, stop ...int) Array[T] {
	var value []T

	if len(stop) == 0 {
		value = a.data[start:]
	} else {
		value = a.data[start:stop[0]]
	}

	return NewArray(value...)
}

func (a *array[T]) Splice(offset int, length int, v ...T) Array[T] {
	a.mutex.Lock()

	defer a.mutex.Unlock()

	endOffset := offset + length

	start := a.data[:offset]
	value := a.data[offset:endOffset]
	end := a.data[endOffset:]

	a.data = append(start, v...)
	a.data = append(a.data, end...)

	return NewArray(value...)
}

func (a *array[T]) Contains(v T) bool {
	return a.IndexOf(v) >= 0
}

func (a *array[T]) IndexOf(v T) int {
	for i, data := range a.data {
		if data == v {
			return i
		}
	}

	return -1
}

func (a *array[T]) Length() int {
	return len(a.data)
}

func (a *array[T]) Find(fn ArrayPredicate[T]) T {
	var value T

	for _, entry := range a.data {
		if ok := fn(entry); ok {
			value = entry
			break
		}
	}

	return value
}

func (a *array[T]) Filter(fn ArrayPredicate[T]) []T {
	value := make([]T, 0)

	for _, entry := range a.data {
		if ok := fn(entry); ok {
			value = append(value, entry)
		}
	}

	return value
}

func (a *array[T]) ForEach(fn ArrayCallback[T]) {
	for i, entry := range a.data {
		fn(entry, i)
	}
}

// ============================

func NewArray[T comparable](data ...T) Array[T] {
	return &array[T]{
		data: data,
	}
}
