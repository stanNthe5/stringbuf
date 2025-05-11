# stringbuf

A Go string concatenation library that is more efficient than strings.Builder.

## Install
```
go get github.com/stanNthe5/stringbuf
```

## Usage
```
sb := stringbuf.New("Hello ", "world,")
sb.Append("I am ", "StringBuf")
sb.Prepend("StringbBuf ", "testing: ")
str := sb.String()
```

## Benchmark

### Compare with strings.Builder
```
// stringbuf_bench_test.go
go test -bench=. -benchmem
```

```
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
BenchmarkStringBuf_Append-6                        10000            209105 ns/op          479674 B/op            16 allocs/op
BenchmarkStringsBuilder_Append-6                    1087           1063544 ns/op         1979769 B/op            24 allocs/op
BenchmarkStringBuf_Prepend-6                       10000            207060 ns/op          479674 B/op            16 allocs/op
BenchmarkStringsBuilder_PrependSimulated-6            14          81451154 ns/op        407881812 B/op         2014 allocs/op
```

## Why is stringbuf faster?

It defers the actual string concatenation (copying data) until `String()` or `Bytes()` is called. During Append or Prepend, it primarily stores references to the input strings in internal `[][]string` slices, chunking them to reduce reallocations compared to strings.Builder which might repeatedly reallocate and copy the growing byte buffer during every append. Prepending is also handled efficiently using a separate buffer, avoiding costly data shifting. (Inspired by the Node.js v8 engine)
