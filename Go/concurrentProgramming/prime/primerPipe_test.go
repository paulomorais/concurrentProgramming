package prime

import (
	"testing"
)

func benchmarkPrimePipe(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunTest(n)
	}
}

func BenchmarkPrimePipe500(b *testing.B)  { benchmarkPrimePipe(500, b) }
func BenchmarkPrimePipe1e3(b *testing.B)  { benchmarkPrimePipe(1e3, b) }
func BenchmarkPrimePipe2e3(b *testing.B)  { benchmarkPrimePipe(2e3, b) }
func BenchmarkPrimePipe3e3(b *testing.B)  { benchmarkPrimePipe(3e3, b) }
func BenchmarkPrimePipe4e3(b *testing.B)  { benchmarkPrimePipe(4e3, b) }
func BenchmarkPrimePipe5e3(b *testing.B)  { benchmarkPrimePipe(5e3, b) }
func BenchmarkPrimePipe1e4(b *testing.B)  { benchmarkPrimePipe(1e4, b) }
func BenchmarkPrimePipe5e4(b *testing.B)  { benchmarkPrimePipe(5e4, b) }
func BenchmarkPrimePipe1e5(b *testing.B)  { benchmarkPrimePipe(1e5, b) }

/*

3000	      414,357 ns/op 500
2000	    1,155,566 ns/op 1e3
500	    3,914,224 ns/op 2e3
200	    7,860,449 ns/op 3e3
100	   12,750,729 ns/op 4e3
100	   18,901,081 ns/op 5e3
20	   62,203,560 ns/op 1e4
1	1,006,057,500 ns/op 5e4
1	3,656,209,100 ns/op 1e5
*/