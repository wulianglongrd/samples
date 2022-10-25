# samples

```shell
$ go test -bench=.
```

## benchmark
```code
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/wulianglongrd/samples
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
Benchmark100IstioMd5KeyFind-4        	    7149	    148424 ns/op
Benchmark1000IstioMd5KeyFind-4       	     807	   1474633 ns/op
Benchmark10000IstioMd5KeyFind-4      	      75	  15681955 ns/op
Benchmark100000IstioMd5KeyFind-4     	       7	 162218185 ns/op
Benchmark200000IstioMd5KeyFind-4     	       3	 341029533 ns/op
Benchmark100AnotherMd5KeyFind-4      	    9860	    106101 ns/op
Benchmark1000AnotherMd5KeyFind-4     	    1018	   1058252 ns/op
Benchmark10000AnotherMd5KeyFind-4    	      85	  12099213 ns/op
Benchmark100000AnotherMd5KeyFind-4   	       9	 117997682 ns/op
Benchmark200000AnotherMd5KeyFind-4   	       4	 284960255 ns/op
Benchmark100HashKeyFind-4            	   14815	     94756 ns/op
Benchmark1000HashKeyFind-4           	    2353	    978508 ns/op
Benchmark10000HashKeyFind-4          	      81	  15593229 ns/op
Benchmark100000HashKeyFind-4         	       8	 147479565 ns/op
Benchmark200000HashKeyFind-4         	       8	 134735181 ns/op
Benchmark100StringKeyFind-4          	   33698	     34066 ns/op
Benchmark1000StringKeyFind-4         	    3073	    356549 ns/op
Benchmark10000StringKeyFind-4        	     264	   4437221 ns/op
Benchmark100000StringKeyFind-4       	      24	  50649196 ns/op
Benchmark200000StringKeyFind-4       	       9	 119251025 ns/op
PASS
ok  	github.com/wulianglongrd/samples	58.493s
```

## string key vs md5 key
```shell
$ go test -bench="IstioMd5KeyFind$" -test.benchmem -count=5 . > md5.txt
$ go test -bench="StringKeyFind$" -test.benchmem -count=5 . > str.txt
$ benchstat md5.txt str.txt
```
note the output first column is different, update to same before execute `benchstat`.

**output is:** 
```code
$ benchstat md5.txt str.txt
name      old time/op    new time/op    delta
100-4        387µs ±24%     116µs ±15%  -69.98%  (p=0.008 n=5+5)
1000-4      4.22ms ±18%    0.92ms ±54%  -78.21%  (p=0.008 n=5+5)
10000-4     32.5ms ±23%     6.1ms ±45%  -81.16%  (p=0.008 n=5+5)
100000-4     289ms ±17%      83ms ±67%  -71.21%  (p=0.008 n=5+5)
200000-4    773ms ±109%     169ms ±46%  -78.20%  (p=0.008 n=5+5)

name      old alloc/op   new alloc/op   delta
100-4       51.1kB ± 0%    47.2kB ± 0%   -7.63%  (p=0.008 n=5+5)
1000-4       529kB ± 0%     500kB ± 0%   -5.52%  (p=0.008 n=5+5)
10000-4     5.15MB ± 0%    4.81MB ± 0%   -6.63%  (p=0.008 n=5+5)
100000-4    50.8MB ± 0%    46.8MB ± 0%   -7.83%  (p=0.008 n=5+5)
200000-4     102MB ± 0%      94MB ± 0%   -7.83%  (p=0.008 n=5+5)

name      old allocs/op  new allocs/op  delta
100-4          802 ± 0%       402 ± 0%  -49.88%  (p=0.008 n=5+5)
1000-4       8.00k ± 0%     4.00k ± 0%  -49.99%  (p=0.008 n=5+5)
10000-4      80.0k ± 0%     40.0k ± 0%  -50.00%  (p=0.008 n=5+5)
100000-4      800k ± 0%      400k ± 0%  -50.00%  (p=0.008 n=5+5)
200000-4     1.60M ± 0%     0.80M ± 0%  -50.00%  (p=0.008 n=5+5)
```