package math

import (
	"math"
	"testing"
)

var (
	i64res int64
	f64res float64
)

func Benchmark_Builtin_Div_AllFloat64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f64res = 1.0 / 2.0
	}
}

func Benchmark_Generic_Div_AllFloat64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f64res = Div[float64](1.0, 2.0)
	}
}

func Benchmark_Builtin_Div_Int64ToFloat64(b *testing.B) {
	va := 1
	vb := 2

	for n := 0; n < b.N; n++ {
		f64res = float64(va) / float64(vb)
	}
}

func Benchmark_Generic_Div_Int64ToFloat64(b *testing.B) {
	va := 1
	vb := 2

	for n := 0; n < b.N; n++ {
		f64res = Div[float64](va, vb)
	}
}

func Benchmark_Builtin_Div_Int64ToInt64(b *testing.B) {
	va := 1
	vb := 2

	for n := 0; n < b.N; n++ {
		i64res = int64(float64(va) / float64(vb))
	}
}

func Benchmark_Generic_Div_Int64ToInt64(b *testing.B) {
	va := 1
	vb := 2

	for n := 0; n < b.N; n++ {
		i64res = Div[int64](va, vb)
	}
}

func Benchmark_Builtin_Remainder(b *testing.B) {
	va := 1.0
	vb := 2.0

	for n := 0; n < b.N; n++ {
		f64res = math.Remainder(va, vb)
	}
}

func Benchmark_Generic_Remainder(b *testing.B) {
	va := 1
	vb := 2

	for n := 0; n < b.N; n++ {
		f64res = Remainder[float64](va, vb)
	}
}
