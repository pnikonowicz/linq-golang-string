package linq_golang_string_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pnikonowicz/linq_golang_string/linq_golang_string"
)

var _ = Describe("Iterate", func() {
	It("enumerate", func() {
		enumerator := linq_golang_string.RootIterator{
			Index: 0,
			Collection: []string{"a","b","c"},
		}

		Expect(enumerator.Next().Value()).To(Equal("a"))
		Expect(enumerator.Next().Value()).To(Equal("b"))
		Expect(enumerator.Next().Value()).To(Equal("c"))
	})
	It("aggregate", func() {
		value := linq_golang_string.Enumerate([]string{"b","c"}).
			Aggregate("a", func(acc string, v string) string {return fmt.Sprintf("%s,%s", acc, v) })

		Expect(value).To(Equal("a,b,c"))
	})
	It("map", func() {
		result := linq_golang_string.Enumerate([]string{"a","b"}).
			Map(func(s string) string {return s+"Amap"}).
			Map(func(s string) string {return s+"Bmap"}).
			Aggregate("seed", func(agg string, item string) string {
				return fmt.Sprintf("%s, %s", agg, item)
			})
		Expect(result).To(Equal("seed, aAmapBmap, bAmapBmap"))
	})
	It("where", func() {
		result := linq_golang_string.Enumerate([]string{"a","b"}).
			Where(func(s string) bool {return s=="a"}).
			Aggregate("", func(agg string, item string) string {
				return fmt.Sprintf("%s%s", agg, item)
			})
		Expect(result).To(Equal("a"))
	})
	It("join", func() {
		result := linq_golang_string.Enumerate([]string{"a","b","c"}).Join(",")
		Expect(result).To(Equal("a,b,c"))

		result = linq_golang_string.Enumerate([]string{"a"}).Join(",")
		Expect(result).To(Equal("a"))

		result = linq_golang_string.Enumerate([]string{}).Join(",")
		Expect(result).To(Equal(""))
	})
})