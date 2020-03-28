package prime_test

import (
	"banmach/prime"
	"testing"
)

func benchmark_IsPrime(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		prime.IsPrime(i)
	}
}

func Benchmark_IsPrime2(b *testing.B)     { benchmark_IsPrime(2, b) }
func Benchmark_IsPrime200(b *testing.B)   { benchmark_IsPrime(200, b) }
func Benchmark_IsPrime20000(b *testing.B) { benchmark_IsPrime(20000, b) }
