go test -bench=. -benchmem -cpu=4
goos: linux
goarch: amd64
pkg: swiss_table
cpu: 11th Gen Intel(R) Core(TM) i9-11900H @ 2.50GHz
BenchmarkPut/100/std-4         	 1293060	       997.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100/mapsutil-4    	 1259182	      1213 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100/mapsutil-OrderedMap-4         	  644725	      1813 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100/mapsutil-SyncLockMap-4        	  356325	      4617 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100/concurrent-swiss-map-4        	  327807	      3402 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100/cockroachdb-4                 	 1612507	      1040 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100/dolthub-4                     	  691689	      1712 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000/std-4                        	   94831	     12691 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000/mapsutil-4                   	   90744	     17033 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000/mapsutil-OrderedMap-4        	   48037	     27461 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000/mapsutil-SyncLockMap-4       	   41912	     28738 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000/concurrent-swiss-map-4       	   24250	     56350 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000/cockroachdb-4                	  147075	     10366 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000/dolthub-4                    	   83358	     15680 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/10000/std-4                       	    4461	    312105 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/10000/mapsutil-4                  	    5487	    203618 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/10000/mapsutil-OrderedMap-4       	    3831	    480605 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/10000/mapsutil-SyncLockMap-4      	    2496	    489567 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/10000/concurrent-swiss-map-4      	    2084	    714911 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/10000/cockroachdb-4               	   11256	    107462 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/10000/dolthub-4                   	    7779	    164512 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100000/std-4                      	     292	   4221363 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100000/mapsutil-4                 	     447	   3721313 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100000/mapsutil-OrderedMap-4      	     297	   4053657 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100000/mapsutil-SyncLockMap-4     	     288	   6487710 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100000/concurrent-swiss-map-4     	     194	   5671646 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100000/cockroachdb-4              	     770	   1504240 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/100000/dolthub-4                  	     602	   2055049 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/500000/std-4                      	      38	  29082303 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/500000/mapsutil-4                 	      52	  27754995 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/500000/mapsutil-OrderedMap-4      	      36	  38517948 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/500000/mapsutil-SyncLockMap-4     	      32	  43839097 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/500000/concurrent-swiss-map-4     	      22	  54117190 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/500000/cockroachdb-4              	      44	  24957631 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/500000/dolthub-4                  	      54	  19688221 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000000/std-4                     	      18	  63966812 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000000/mapsutil-4                	      20	  71459566 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000000/mapsutil-OrderedMap-4     	      13	 100489689 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000000/mapsutil-SyncLockMap-4    	      12	 100637893 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000000/concurrent-swiss-map-4    	       4	 271244211 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000000/cockroachdb-4             	      20	  71811020 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/1000000/dolthub-4                 	      30	  58357690 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/2500000/std-4                     	       7	 167286472 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/2500000/mapsutil-4                	       6	 194432724 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/2500000/mapsutil-OrderedMap-4     	       3	 343186153 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/2500000/mapsutil-SyncLockMap-4    	       3	 335323087 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/2500000/concurrent-swiss-map-4    	       2	 724138862 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/2500000/cockroachdb-4             	       7	 209065446 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/2500000/dolthub-4                 	       7	 251871224 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/5000000/std-4                     	       1	1147961747 ns/op	495454744 B/op	  153707 allocs/op
BenchmarkPut/5000000/mapsutil-4                	       1	1107358047 ns/op	495403992 B/op	  153463 allocs/op
BenchmarkPut/5000000/mapsutil-OrderedMap-4     	       1	1501888739 ns/op	918366808 B/op	  153977 allocs/op
BenchmarkPut/5000000/mapsutil-SyncLockMap-4    	       1	1210824109 ns/op	495319336 B/op	  153056 allocs/op
BenchmarkPut/5000000/concurrent-swiss-map-4    	       1	1643820761 ns/op	169811968 B/op	      38 allocs/op
BenchmarkPut/5000000/cockroachdb-4             	       3	 448747728 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/5000000/dolthub-4                 	       3	 594676105 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/10000000/std-4                    	       1	2946459167 ns/op	990815272 B/op	  307025 allocs/op
BenchmarkPut/10000000/mapsutil-4               	       1	2390261866 ns/op	990773256 B/op	  306823 allocs/op
BenchmarkPut/10000000/mapsutil-OrderedMap-4    	       1	3005836401 ns/op	1817720008 B/op	  307200 allocs/op
BenchmarkPut/10000000/mapsutil-SyncLockMap-4   	       1	2889094703 ns/op	990946936 B/op	  307658 allocs/op
BenchmarkPut/10000000/concurrent-swiss-map-4   	       1	4130893714 ns/op	285868032 B/op	      32 allocs/op
BenchmarkPut/10000000/cockroachdb-4            	       2	 999217354 ns/op	       0 B/op	       0 allocs/op
BenchmarkPut/10000000/dolthub-4                	       2	1306210176 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100/std-4                         	137221442	        11.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100/mapsutil-4                    	100000000	        11.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100/mapsutil-OrderedMap-4         	138008322	        11.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100/mapsutil-SyncLockMap-4        	52885318	        21.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100/concurrent-swiss-map-4        	40911187	        28.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100/cockroachdb-4                 	100000000	        11.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100/dolthub-4                     	64622610	        18.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000/std-4                        	158834455	        10.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000/mapsutil-4                   	100000000	        11.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000/mapsutil-OrderedMap-4        	100000000	        10.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000/mapsutil-SyncLockMap-4       	59887330	        21.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000/concurrent-swiss-map-4       	31235960	        49.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000/cockroachdb-4                	100000000	        12.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000/dolthub-4                    	74935575	        14.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000/std-4                       	54695334	        35.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000/mapsutil-4                  	53050880	        23.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000/mapsutil-OrderedMap-4       	46322793	        31.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000/mapsutil-SyncLockMap-4      	39388628	        26.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000/concurrent-swiss-map-4      	26125441	        61.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000/cockroachdb-4               	91851310	        15.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000/dolthub-4                   	59506518	        20.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100000/std-4                      	29013787	        42.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100000/mapsutil-4                 	40098703	        49.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100000/mapsutil-OrderedMap-4      	36277192	        48.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100000/mapsutil-SyncLockMap-4     	30177457	        52.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100000/concurrent-swiss-map-4     	19558730	        66.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100000/cockroachdb-4              	77947516	        22.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/100000/dolthub-4                  	53503814	        29.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/500000/std-4                      	22397205	        63.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/500000/mapsutil-4                 	23988702	        57.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/500000/mapsutil-OrderedMap-4      	21484336	        46.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/500000/mapsutil-SyncLockMap-4     	17233070	        80.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/500000/concurrent-swiss-map-4     	10812498	       111.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/500000/cockroachdb-4              	22483209	        47.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/500000/dolthub-4                  	32059609	        36.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000000/std-4                     	16962226	        72.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000000/mapsutil-4                	17481681	        70.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000000/mapsutil-OrderedMap-4     	17082783	        68.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000000/mapsutil-SyncLockMap-4    	13146884	        95.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000000/concurrent-swiss-map-4    	 5861941	       187.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000000/cockroachdb-4             	15026182	        76.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/1000000/dolthub-4                 	21819390	        66.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/2500000/std-4                     	17517018	        66.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/2500000/mapsutil-4                	16819360	        64.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/2500000/mapsutil-OrderedMap-4     	16769350	        88.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/2500000/mapsutil-SyncLockMap-4    	14836004	        92.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/2500000/concurrent-swiss-map-4    	 7343095	       234.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/2500000/cockroachdb-4             	14252509	        99.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/2500000/dolthub-4                 	15178856	        99.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/5000000/std-4                     	18854079	        68.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/5000000/mapsutil-4                	20944143	        85.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/5000000/mapsutil-OrderedMap-4     	20529351	        77.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/5000000/mapsutil-SyncLockMap-4    	12341049	        94.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/5000000/concurrent-swiss-map-4    	 7334277	       258.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/5000000/cockroachdb-4             	13587732	        99.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/5000000/dolthub-4                 	11597472	       126.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000000/std-4                    	19163136	        68.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000000/mapsutil-4               	23107812	        71.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000000/mapsutil-OrderedMap-4    	19276416	        84.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000000/mapsutil-SyncLockMap-4   	12163428	       112.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000000/concurrent-swiss-map-4   	 5203172	       255.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000000/cockroachdb-4            	14585926	        91.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet/10000000/dolthub-4                	10143468	       118.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100/std-4                      	388240143	         3.573 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100/mapsutil/no-op-4           	1000000000	         0.2901 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100/mapsutil-OrderedMap-4      	289610062	         3.882 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100/mapsutil-SyncLockMap-4     	39636910	        29.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100/concurrent-swiss-map-4     	32289637	        31.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100/cockroachdb-4              	145677853	        10.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100/dolthub-4                  	66875745	        18.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000/std-4                     	389300572	         2.718 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000/mapsutil/no-op-4          	1000000000	         0.2125 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000/mapsutil-OrderedMap-4     	426547341	         2.840 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000/mapsutil-SyncLockMap-4    	55239030	        28.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000/concurrent-swiss-map-4    	37994768	        30.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000/cockroachdb-4             	145390003	         8.725 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000/dolthub-4                 	44546529	        34.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000/std-4                    	512519494	         2.836 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000/mapsutil/no-op-4         	1000000000	         0.2743 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000/mapsutil-OrderedMap-4    	291977106	         4.090 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000/mapsutil-SyncLockMap-4   	37411615	        27.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000/concurrent-swiss-map-4   	31456320	        55.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000/cockroachdb-4            	63749499	        21.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000/dolthub-4                	32669562	        33.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100000/std-4                   	480873357	         2.488 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100000/mapsutil/no-op-4        	1000000000	         0.2608 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100000/mapsutil-OrderedMap-4   	   13027	     93844 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100000/mapsutil-SyncLockMap-4  	48871603	        28.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100000/concurrent-swiss-map-4  	21582734	        56.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100000/cockroachdb-4           	39460038	        29.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/100000/dolthub-4               	34307203	        37.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/500000/std-4                   	491703831	         3.425 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/500000/mapsutil/no-op-4        	1000000000	         0.3002 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/500000/mapsutil-OrderedMap-4   	    1180	   1015864 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/500000/mapsutil-SyncLockMap-4  	52937606	        24.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/500000/concurrent-swiss-map-4  	19033789	        67.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/500000/cockroachdb-4           	46645526	        28.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/500000/dolthub-4               	21339711	        52.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000000/std-4                  	295072033	         3.598 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000000/mapsutil/no-op-4       	1000000000	         0.2264 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000000/mapsutil-OrderedMap-4  	     669	   2147872 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000000/mapsutil-SyncLockMap-4 	40946436	        30.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000000/concurrent-swiss-map-4 	20691993	        56.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000000/cockroachdb-4          	17863585	        58.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/1000000/dolthub-4              	30721501	        57.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/2500000/std-4                  	358924080	         3.748 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/2500000/mapsutil/no-op-4       	1000000000	         0.2265 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/2500000/mapsutil-OrderedMap-4  	     322	   3351941 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/2500000/mapsutil-SyncLockMap-4 	52610647	        26.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/2500000/concurrent-swiss-map-4 	16128831	        81.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/2500000/cockroachdb-4          	16056295	        67.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/2500000/dolthub-4              	27717394	        61.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/5000000/std-4                  	453804004	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/5000000/mapsutil/no-op-4       	1000000000	         0.2242 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/5000000/mapsutil-OrderedMap-4  	     100	  10708694 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/5000000/mapsutil-SyncLockMap-4 	54711370	        24.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/5000000/concurrent-swiss-map-4 	18250561	        62.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/5000000/cockroachdb-4          	13376004	        84.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/5000000/dolthub-4              	18583477	        69.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000000/std-4                 	13458648	        93.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000000/mapsutil/no-op-4      	1000000000	         0.2775 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000000/mapsutil-OrderedMap-4 	      49	  23727232 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000000/mapsutil-SyncLockMap-4         	12911248	       131.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000000/concurrent-swiss-map-4         	 4827019	       251.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000000/cockroachdb-4                  	11359152	       100.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete/10000000/dolthub-4                      	 9482443	       127.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	swiss_table	474.344s