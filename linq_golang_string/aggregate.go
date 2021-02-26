package linq_golang_string

import (
	"fmt"
)

func (e Enumerable) Aggregate(seed string, enumeration func(string, string) string) string {
	result := seed
	for {
		value := e.iterator.Next()
		switch v := value.(type) {
		case Empty:
			return result
		case Value:
			result = enumeration(result, value.Value())
		default:
			panic(fmt.Errorf("unknown type: %s", v))
		}
	}
	panic("loop broken in unexpected way")
}

