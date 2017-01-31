package parallel

import (
	"testing"
)

func benchmarkParallelMatrix(n int, p int, b *testing.B) {
	Setup(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Run(n, p)
	}
}
func BenchmarkParallelMatrixN4T1(b *testing.B)  { benchmarkParallelMatrix(4, 1, b) }
func BenchmarkParallelMatrixN4T2(b *testing.B)  { benchmarkParallelMatrix(4, 2, b) }
func BenchmarkParallelMatrixN4T4(b *testing.B)  { benchmarkParallelMatrix(4, 4, b) }
func BenchmarkParallelMatrixN8T1(b *testing.B)  { benchmarkParallelMatrix(8, 1, b) }
func BenchmarkParallelMatrixN8T2(b *testing.B)  { benchmarkParallelMatrix(8, 2, b) }
func BenchmarkParallelMatrixN8T4(b *testing.B)  { benchmarkParallelMatrix(8, 4, b) }
func BenchmarkParallelMatrixN8T8(b *testing.B)  { benchmarkParallelMatrix(8, 8, b) }
func BenchmarkParallelMatrixN16T1(b *testing.B)  { benchmarkParallelMatrix(16, 1, b) }
func BenchmarkParallelMatrixN16T2(b *testing.B)  { benchmarkParallelMatrix(16, 2, b) }
func BenchmarkParallelMatrixN16T4(b *testing.B)  { benchmarkParallelMatrix(16, 4, b) }
func BenchmarkParallelMatrixN16T8(b *testing.B)  { benchmarkParallelMatrix(16, 8, b) }
func BenchmarkParallelMatrixN16T16(b *testing.B)  { benchmarkParallelMatrix(16, 16, b) }
func BenchmarkParallelMatrixN32T1(b *testing.B)  { benchmarkParallelMatrix(32, 1, b) }
func BenchmarkParallelMatrixN32T2(b *testing.B)  { benchmarkParallelMatrix(32, 2, b) }
func BenchmarkParallelMatrixN32T4(b *testing.B)  { benchmarkParallelMatrix(32, 4, b) }
func BenchmarkParallelMatrixN32T8(b *testing.B)  { benchmarkParallelMatrix(32, 8, b) }
func BenchmarkParallelMatrixN32T16(b *testing.B)  { benchmarkParallelMatrix(32, 16, b) }
func BenchmarkParallelMatrixN32T32(b *testing.B)  { benchmarkParallelMatrix(32, 32, b) }
func BenchmarkParallelMatrixN64T1(b *testing.B)  { benchmarkParallelMatrix(64, 1, b) }
func BenchmarkParallelMatrixN64T2(b *testing.B)  { benchmarkParallelMatrix(64, 2, b) }
func BenchmarkParallelMatrixN64T4(b *testing.B)  { benchmarkParallelMatrix(64, 4, b) }
func BenchmarkParallelMatrixN64T8(b *testing.B)  { benchmarkParallelMatrix(64, 8, b) }
func BenchmarkParallelMatrixN64T16(b *testing.B)  { benchmarkParallelMatrix(64, 16, b) }
func BenchmarkParallelMatrixN64T32(b *testing.B)  { benchmarkParallelMatrix(64, 32, b) }
func BenchmarkParallelMatrixN64T64(b *testing.B)  { benchmarkParallelMatrix(64, 64, b) }
func BenchmarkParallelMatrixN128T1(b *testing.B)	{ benchmarkParallelMatrix(128, 1, b) }
func BenchmarkParallelMatrixN128T2(b *testing.B)	{ benchmarkParallelMatrix(128, 2, b) }
func BenchmarkParallelMatrixN128T4(b *testing.B) 	{ benchmarkParallelMatrix(128, 4, b) }
func BenchmarkParallelMatrixN128T8(b *testing.B) 	{ benchmarkParallelMatrix(128, 8, b) }
func BenchmarkParallelMatrixN128T16(b *testing.B) 	{ benchmarkParallelMatrix(128, 16, b) }
func BenchmarkParallelMatrixN128T32(b *testing.B) 	{ benchmarkParallelMatrix(128, 32, b) }
func BenchmarkParallelMatrixN128T64(b *testing.B) 	{ benchmarkParallelMatrix(128, 64, b) }
func BenchmarkParallelMatrixN128T128(b *testing.B) 	{ benchmarkParallelMatrix(128, 128, b) }
func BenchmarkParallelMatrixN256T1(b *testing.B)	{ benchmarkParallelMatrix(256, 1, b) }
func BenchmarkParallelMatrixN256T2(b *testing.B)	{ benchmarkParallelMatrix(256, 2, b) }
func BenchmarkParallelMatrixN256T4(b *testing.B) 	{ benchmarkParallelMatrix(256, 4, b) }
func BenchmarkParallelMatrixN256T8(b *testing.B) 	{ benchmarkParallelMatrix(256, 8, b) }
func BenchmarkParallelMatrixN256T16(b *testing.B) 	{ benchmarkParallelMatrix(256, 16, b) }
func BenchmarkParallelMatrixN256T32(b *testing.B) 	{ benchmarkParallelMatrix(256, 32, b) }
func BenchmarkParallelMatrixN256T64(b *testing.B) 	{ benchmarkParallelMatrix(256, 64, b) }
func BenchmarkParallelMatrixN256T128(b *testing.B) 	{ benchmarkParallelMatrix(256, 128, b) }
func BenchmarkParallelMatrixN256T256(b *testing.B) 	{ benchmarkParallelMatrix(256, 256, b) }
func BenchmarkParallelMatrixN512T1(b *testing.B)	{ benchmarkParallelMatrix(512, 1, b) }
func BenchmarkParallelMatrixN512T2(b *testing.B)	{ benchmarkParallelMatrix(512, 2, b) }
func BenchmarkParallelMatrixN512T4(b *testing.B) 	{ benchmarkParallelMatrix(512, 4, b) }
func BenchmarkParallelMatrixN512T8(b *testing.B) 	{ benchmarkParallelMatrix(512, 8, b) }
func BenchmarkParallelMatrixN512T16(b *testing.B) 	{ benchmarkParallelMatrix(512, 16, b) }
func BenchmarkParallelMatrixN512T32(b *testing.B) 	{ benchmarkParallelMatrix(512, 32, b) }
func BenchmarkParallelMatrixN512T64(b *testing.B) 	{ benchmarkParallelMatrix(512, 64, b) }
func BenchmarkParallelMatrixN512T128(b *testing.B) 	{ benchmarkParallelMatrix(512, 128, b) }
func BenchmarkParallelMatrixN512T256(b *testing.B) 	{ benchmarkParallelMatrix(512, 256, b) }
func BenchmarkParallelMatrixN512T512(b *testing.B) 	{ benchmarkParallelMatrix(512, 512, b) }
func BenchmarkParallelMatrixN1024T1(b *testing.B)	{ benchmarkParallelMatrix(1024, 1, b) }
func BenchmarkParallelMatrixN1024T2(b *testing.B)	{ benchmarkParallelMatrix(1024, 2, b) }
func BenchmarkParallelMatrixN1024T4(b *testing.B) 	{ benchmarkParallelMatrix(1024, 4, b) }
func BenchmarkParallelMatrixN1024T8(b *testing.B) 	{ benchmarkParallelMatrix(1024, 8, b) }
func BenchmarkParallelMatrixN1024T16(b *testing.B) 	{ benchmarkParallelMatrix(1024, 16, b) }
func BenchmarkParallelMatrixN1024T32(b *testing.B) 	{ benchmarkParallelMatrix(1024, 32, b) }
func BenchmarkParallelMatrixN1024T64(b *testing.B) 	{ benchmarkParallelMatrix(1024, 64, b) }
func BenchmarkParallelMatrixN1024T128(b *testing.B)	{ benchmarkParallelMatrix(1024, 128, b) }
func BenchmarkParallelMatrixN1024T256(b *testing.B)	{ benchmarkParallelMatrix(1024, 256, b) }
func BenchmarkParallelMatrixN1024T512(b *testing.B) 	{ benchmarkParallelMatrix(1024, 512, b) }
func BenchmarkParallelMatrixN1024T1024(b *testing.B) 	{ benchmarkParallelMatrix(1024, 1024, b) }

/*
repeat	time/operation
3000	    488694 ns/op	N64T1
5000	    265015 ns/op	N64T2
10000	    165209 ns/op	N64T4
10000	    157308 ns/op	N64T8
10000	    176310 ns/op	N64T16
10000	    186110 ns/op	N64T32
10000	    193611 ns/op	N64T64
300	   4736937 ns/op	N128T1
1000	   2440139 ns/op	N128T2
1000	   2023115 ns/op	N128T4
1000	   1787102 ns/op	N128T8
1000	   1805103 ns/op	N128T16
1000	   1799102 ns/op	N128T32
1000	   1946111 ns/op	N128T64
1000	   1837105 ns/op	N128T128
20	  50302880 ns/op	N256T1
50	  27081548 ns/op	N256T2
100	  24071377 ns/op	N256T4
100	  20161154 ns/op	N256T8
50	  20121150 ns/op	N256T16
100	  20361164 ns/op	N256T32
100	  22531289 ns/op	N256T64
100	  21211213 ns/op	N256T128
100	  21401224 ns/op	N256T256
3	 519029700 ns/op	N512T1
5	 287016420 ns/op	N512T2
10	 237313580 ns/op	N512T4
5	 255014580 ns/op	N512T8
5	 259414820 ns/op	N512T16
10	 265215170 ns/op	N512T32
5	 264215120 ns/op	N512T64
10	 266015210 ns/op	N512T128
10	 268515360 ns/op	N512T256
10	 268915390 ns/op	N512T512
1	7868450000 ns/op	N1024T1
1	4409252200 ns/op	N1024T2
1	3097177100 ns/op	N1024T4
1	2572147100 ns/op	N1024T8
1	2630150500 ns/op	N1024T16
1	2638150900 ns/op	N1024T32
1	2749157200 ns/op	N1024T64
1	2764158100 ns/op	N1024T128
1	2727156000 ns/op	N1024T256
1	2810160800 ns/op	N1024T512
1	2756157700 ns/op	N1024T1024


2000	     778,044 ns/op	N64T1
3000	     585,366 ns/op	N64T2
3000	     500,695 ns/op	N64T4
3000	     469,360 ns/op	N64T8
3000	     459,359 ns/op	N64T16
3000	     474,693 ns/op	N64T32
3000	     474,360 ns/op	N64T64
300	    5,676,991 ns/op	N128T1
500	    3,492,199 ns/op	N128T2
1000	    2,782,159 ns/op	N128T4
500	    2,668,152 ns/op	N128T8
500	    2,908,166 ns/op	N128T16
500	    2,892,165 ns/op	N128T32
500	    2,866,164 ns/op	N128T64
20	   55,053,150 ns/op	N256T1
50	   33,761,932 ns/op	N256T2
50	   29,421,684 ns/op	N256T4
50	   27,581,576 ns/op	N256T8
50	   26,121,494 ns/op	N256T16
100	   25,651,467 ns/op	N256T32
50	   26,461,514 ns/op	N256T64
2	  599,034,250 ns/op	N512T1
3	  353,020,200 ns/op	N512T2
5	  266,215,220 ns/op	N512T4
5	  248,014,180 ns/op	N512T8
5	  263,215,060 ns/op	N512T16
5	  268,615,360 ns/op	N512T32
5	  250,214,320 ns/op	N512T64
1	9,029,516,500 ns/op	N1024T1
1	5,029,287,700 ns/op	N1024T2
1	3,034,173,600 ns/op	N1024T4
1	2,531,144,800 ns/op	N1024T8
1	2,519,144,000 ns/op	N1024T16
1	2,541,145,300 ns/op	N1024T32
1	2,541,145,300 ns/op	N1024T64
*/