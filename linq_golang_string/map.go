package linq_golang_string

import (
	"fmt"
)

type MapIterator struct {
	value func() IteratorValue
	enumeration func(string) string
}

func (o *MapIterator) Next() IteratorValue {
	value := o.value()
	switch v := value.(type) {
	case Empty:
		return v
	case Value:
		return Value {
			V: o.enumeration(value.Value()),
		}
	default:
		panic(fmt.Errorf("unknown type: %s", v))
	}
}

func (e Enumerable) Map(enumeration func(string) string) Enumerable {
	return Enumerable {
		iterator: &MapIterator {
			enumeration: enumeration,
			value: func() IteratorValue { return e.iterator.Next()},
		},
	}
}
