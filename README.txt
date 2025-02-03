$ go test -bench=. -benchmem -benchtime 5s
goos: linux
goarch: amd64
pkg: swiss_table
cpu: 11th Gen Intel(R) Core(TM) i9-11900H @ 2.50GHz
BenchmarkPut/concurrent-swiss-map-16         	   74023	     80565 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/cockroachdb-16                  	  412298	     14456 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/dolthub-16                      	  199984	     29748 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/concurrent-swiss-map-16         	86112866	        66.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/cockroachdb-16                  	396088783	        15.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/dolthub-16                      	239704557	        24.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/concurrent-swiss-map-16      	100000000	        59.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/cockroachdb-16               	405566287	        14.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/dolthub-16                   	154245849	        39.11 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	swiss_table	64.293s