# Json valid benchmark
This repository contains simple code for validating JSON strings to showcase the speed difference between `json.Valid` and `json.Unmarshal(...) == nil`

## Run benchmarks
``` bash
make bench
```

### Results

Hardware: cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

| Benchmark name | Iterations | Time per operation | Bits allocated per operation | Allocations per operation |
|-----|-----|----|----|----|
| BenchmarkJsonValid-12 | 2416996 | 486.0 ns/op | 144 B/op | 1 allocs/op |
| BenchmarkJsonUnmarshalValid-12 | 645478 | 1844 ns/op | 1152 B/op | 19 allocs/op |
