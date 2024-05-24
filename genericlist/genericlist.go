package genericlist

type GenericList[T any] struct {
	data []T
}

func New[T any]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Insert(value T) {
	l.data = append(l.data, value)
}

func (l *GenericList[T]) Get(i int) T {
	if i < 0 || i >= len(l.data) {
		panic("Invalid index")
	}
	return l.data[i]
}
