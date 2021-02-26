package linq_golang_string

import "fmt"

type WhereIterator struct {
	value func() IteratorValue
	predicate func(string) bool
}

func (o *WhereIterator) Next() IteratorValue {
	for {
		value := o.value()
		switch v := value.(type) {
		case Empty:
			return value
		case Value:
			if o.predicate(value.Value()) {
				return value
			}
		default:
			panic(fmt.Errorf("unknown type: %s", v))
		}
	}
	panic("loop broken in unexpected way")
}

func (e Enumerable) Where(predicate func(string) bool) Enumerable {
	return Enumerable {
		iterator: &WhereIterator {
			predicate: predicate,
			value: func() IteratorValue { return e.iterator.Next()},
		},
	}
}
