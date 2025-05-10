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

# Benchmark
(source code: stringbuf_bench_test.go)
```
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
BenchmarkStringBuf_Append-6                        10000            240271 ns/op
BenchmarkStringsBuilder_Append-6                    1407            766408 ns/op
BenchmarkStringBuf_Prepend-6                        5534            247672 ns/op
BenchmarkStringsBuilder_PrependSimulated-6             6         183996911 ns/op
PASS
```