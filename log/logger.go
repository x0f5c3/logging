package log

import (

	"fmt"
	"sync"

)

type (
	Marshaler[T any] func(T) (string, error)
	FieldType[T any] interface {
		fmt.Stringer
		Interface() (interface{}, error)
	}
	Field[T any] struct {
		FieldType[T]
		val         T
		marshalFunc Marshaler[T]
	}
)

func (f *Field[T]) Val() T {
	return f.val
}

func (f *Field[T]) SetVal(val T) {
	f.val = val
}

func (f *Field[T]) MarshalFunc() Marshaler[T] {
	return f.marshalFunc
}

func (f *Field[T]) SetMarshalFunc(marshalFunc Marshaler[T]) {
	f.marshalFunc = marshalFunc
}

type Logger struct {
	mu     *sync.Mutex
	fields []Field[any]
}
