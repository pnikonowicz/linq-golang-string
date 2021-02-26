package linq_golang_string

type Iterator interface {
	Next() IteratorValue
}

type IteratorValue interface {
	Value() string
}

type Empty struct {
}
func (o Empty) Value() string {
	panic("empty Value called")
}
type Value struct {
	V string
}
func (o Value) Value() string {
	return o.V
}
