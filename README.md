This is a simple attempt at creating a syntax inspired by linq in C#.
It's not exact, it's not trying to be. I just wanted to see if it could be done.

code example:

```go
linq_golang_string.Enumerate([]string{"a","b"}).
    Where(func(s string) bool {return s=="a"}).
    Map(func(s string) string {return s+"Amap"}).
    Map(func(s string) string {return s+"Bmap"}).
    Aggregate("seed", func(agg string, item string) string {
        return fmt.Sprintf("%s, %s", agg, item)
    })
```