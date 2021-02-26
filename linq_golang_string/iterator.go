package linq_golang_string

type Enumerable struct {
	iterator Iterator
}

type RootIterator struct {
	Index      int
	Collection []string
}
func (o* RootIterator) Next() IteratorValue {
	if o.Index <len(o.Collection) {
		value := Value{
			V: o.Collection[o.Index],
		}
		o.Index +=1
		return value
	} else {
		return Empty{}
	}
}

func Enumerate(collection []string) Enumerable {
	return Enumerable {
		iterator: &RootIterator{
			Index: 0,
			Collection: collection,
		},
	}
}

