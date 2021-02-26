package linq_golang_string

import (
	"fmt"
)

func (e Enumerable) Join(separator string) string {
	next := e.iterator.Next()
	result := ""
	switch nextType := next.(type) {
		case Empty:
			return result
		case Value:
			return e.Aggregate(next.Value(), func(a,b string) string {return fmt.Sprintf("%s%s%s", a, separator, b)})
		default:
			panic(fmt.Errorf("unknown type: %s", nextType))
	}
}

