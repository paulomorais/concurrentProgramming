package threads

import (
	"testing"
)

func benchmarkFibonacciThread(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(n)
	}
}

func BenchmarkFibonacciThread0(b *testing.B)  { benchmarkFibonacciThread(0, b) }
func BenchmarkFibonacciThread1(b *testing.B)  { benchmarkFibonacciThread(1, b) }
func BenchmarkFibonacciThread2(b *testing.B)  { benchmarkFibonacciThread(2, b) }
func BenchmarkFibonacciThread3(b *testing.B)  { benchmarkFibonacciThread(3, b) }
func BenchmarkFibonacciThread4(b *testing.B)  { benchmarkFibonacciThread(4, b) }
func BenchmarkFibonacciThread5(b *testing.B)  { benchmarkFibonacciThread(5, b) }
func BenchmarkFibonacciThread6(b *testing.B)  { benchmarkFibonacciThread(6, b) }
func BenchmarkFibonacciThread7(b *testing.B)  { benchmarkFibonacciThread(7, b) }
func BenchmarkFibonacciThread8(b *testing.B)  { benchmarkFibonacciThread(8, b) }
func BenchmarkFibonacciThread9(b *testing.B)  { benchmarkFibonacciThread(9, b) }
func BenchmarkFibonacciThread10(b *testing.B)  { benchmarkFibonacciThread(10, b) }
func BenchmarkFibonacciThread11(b *testing.B)  { benchmarkFibonacciThread(11, b) }
func BenchmarkFibonacciThread12(b *testing.B)  { benchmarkFibonacciThread(12, b) }
func BenchmarkFibonacciThread13(b *testing.B)  { benchmarkFibonacciThread(13, b) }
func BenchmarkFibonacciThread14(b *testing.B)  { benchmarkFibonacciThread(14, b) }
func BenchmarkFibonacciThread15(b *testing.B)  { benchmarkFibonacciThread(15, b) }
func BenchmarkFibonacciThread16(b *testing.B)  { benchmarkFibonacciThread(16, b) }
func BenchmarkFibonacciThread17(b *testing.B)  { benchmarkFibonacciThread(17, b) }
func BenchmarkFibonacciThread18(b *testing.B)  { benchmarkFibonacciThread(18, b) }
func BenchmarkFibonacciThread19(b *testing.B)  { benchmarkFibonacciThread(19, b) }
func BenchmarkFibonacciThread20(b *testing.B)  { benchmarkFibonacciThread(20, b) }
func BenchmarkFibonacciThread21(b *testing.B)  { benchmarkFibonacciThread(21, b) }
func BenchmarkFibonacciThread22(b *testing.B)  { benchmarkFibonacciThread(22, b) }

/*
3000000	       600 ns/op	0	Go
3000000	       600 ns/op	1	Go
1000000	      2244 ns/op	2	Go
300000	      4093 ns/op	3	Go
300000	      4876 ns/op	4	Go
200000	      7050 ns/op	5	Go
200000	     11030 ns/op	6	Go
100000	     15410 ns/op	7	Go
100000	     19951 ns/op	8	Go
50000	     26801 ns/op	9	Go
50000	     37622 ns/op	10	Go
30000	     53569 ns/op	11	Go
20000	     78604 ns/op	12	Go
10000	    118306 ns/op	13	Go
10000	    182510 ns/op	14	Go
5000	    298617 ns/op	15	Go
3000	    495361 ns/op	16	Go
2000	    813046 ns/op	17	Go
1000	   1490085 ns/op	18	Go
500	   2756157 ns/op	19	Go
300	   4653599 ns/op	20	Go

2000000	       610 ns/op
3000000	       604 ns/op
1000000	      2310 ns/op
300000	      4580 ns/op
300000	      5796 ns/op
200000	      8390 ns/op
200000	     11695 ns/op
100000	     16070 ns/op
100000	     21931 ns/op
50000	     30201 ns/op
30000	     41635 ns/op
30000	     60170 ns/op
20000	     85404 ns/op
10000	    129207 ns/op
10000	    198711 ns/op
5000	    318018 ns/op
3000	    535697 ns/op
2000	   1014058 ns/op
1000	   1658094 ns/op
500	   3196182 ns/op
300	   5023620 ns/op
200	   8505486 ns/op
100	  14600835 ns/op
*/
