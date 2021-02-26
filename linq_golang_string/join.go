package linq_golang_string

import (
	"fmt"
)

func (e Enumerable) Join(separator string) string {
	next := e.iterator.Next()
	result := ""
	switch prevType := next.(type) {
		case Empty:
			return result
		case Value:
			result=next.Value()
			for {
				next := e.iterator.Next()
				switch v := next.(type) {
				case Empty:
					return fmt.Sprintf("%s", result)
				case Value:
					result = fmt.Sprintf("%s%s%s", result, separator, next.Value())
				default:
					panic(fmt.Errorf("unknown type: %s", v))
				}
			}
			panic("loop broken in unexpected way")
		default:
			panic(fmt.Errorf("unknown type: %s", prevType))
	}
}

