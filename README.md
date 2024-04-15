# Golang Marshal & Unmarshal Benchmark


## Results

### Marshal
```
$ go test -bench=. -benchtime=5000x benchmark/marshal/*

goos: darwin
goarch: arm64
BenchmarkEncodingJson-12                	    5000	     11308 ns/op
BenchmarkEncodingJsonIterDefault-12     	    5000	      8638 ns/op
BenchmarkEncodingJsonIterFastest-12     	    5000	      7848 ns/op
BenchmarkEncodingJsonIterStandard-12    	    5000	      8303 ns/op
BenchmarkGoJson-12                      	    5000	      4093 ns/op
PASS
ok  	command-line-arguments	0.659s
```


### Unmarshal
```
$ go test -bench=. -benchtime=5000x benchmark/unmarsal/*

goos: darwin
goarch: arm64
BenchmarkEncodingJson-12                	    5000	     44928 ns/op
BenchmarkEncodingJsonIterDefault-12     	    5000	     11352 ns/op
BenchmarkEncodingJsonIterFastest-12     	    5000	     11278 ns/op
BenchmarkEncodingJsonIterStandard-12    	    5000	     11229 ns/op
BenchmarkGoJson-12                      	    5000	      8264 ns/op
PASS
ok  	command-line-arguments	0.668s
```

### Marshal + Unmarshal
```
$ go test -bench=. -benchtime=5000x benchmark/marshalunmarsal/*

goos: darwin
goarch: arm64
BenchmarkEncodingJson-12                	    5000	     52505 ns/op
BenchmarkEncodingJsonIterDefault-12     	    5000	     19728 ns/op
BenchmarkEncodingJsonIterFastest-12     	    5000	     19771 ns/op
BenchmarkEncodingJsonIterStandard-12    	    5000	     20092 ns/op
BenchmarkGoJson-12                      	    5000	     12894 ns/op
PASS
ok  	command-line-arguments	0.852s
```