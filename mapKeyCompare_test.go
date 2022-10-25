package samples

import (
	"math/rand"
	"testing"
)

// istio current used md5 key
func Benchmark100IstioMd5KeyFind(b *testing.B) {
	benchmarkIstioMd5KeyFind(b, 100)
}
func Benchmark1000IstioMd5KeyFind(b *testing.B) {
	benchmarkIstioMd5KeyFind(b, 1000)
}
func Benchmark10000IstioMd5KeyFind(b *testing.B) {
	benchmarkIstioMd5KeyFind(b, 10000)
}
func Benchmark100000IstioMd5KeyFind(b *testing.B) {
	benchmarkIstioMd5KeyFind(b, 100000)
}
func Benchmark200000IstioMd5KeyFind(b *testing.B) {
	benchmarkIstioMd5KeyFind(b, 200000)
}
func benchmarkIstioMd5KeyFind(b *testing.B, max int) {
	samples := buildTestSamples(max)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IstioMd5KeyFind(samples)
	}
}

func Benchmark100AnotherMd5KeyFind(b *testing.B) {
	benchmarkAnotherMd5KeyFind(b, 100)
}
func Benchmark1000AnotherMd5KeyFind(b *testing.B) {
	benchmarkAnotherMd5KeyFind(b, 1000)
}
func Benchmark10000AnotherMd5KeyFind(b *testing.B) {
	benchmarkAnotherMd5KeyFind(b, 10000)
}
func Benchmark100000AnotherMd5KeyFind(b *testing.B) {
	benchmarkAnotherMd5KeyFind(b, 100000)
}
func Benchmark200000AnotherMd5KeyFind(b *testing.B) {
	benchmarkAnotherMd5KeyFind(b, 200000)
}
func benchmarkAnotherMd5KeyFind(b *testing.B, max int) {
	samples := buildTestSamples(max)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		AnotherMd5KeyFind(samples)
	}
}

func Benchmark100HashKeyFind(b *testing.B) {
	benchmarkHashKeyFind(b, 100)
}
func Benchmark1000HashKeyFind(b *testing.B) {
	benchmarkHashKeyFind(b, 1000)
}
func Benchmark10000HashKeyFind(b *testing.B) {
	benchmarkHashKeyFind(b, 10000)
}
func Benchmark100000HashKeyFind(b *testing.B) {
	benchmarkHashKeyFind(b, 100000)
}
func Benchmark200000HashKeyFind(b *testing.B) {
	benchmarkHashKeyFind(b, 200000)
}
func benchmarkHashKeyFind(b *testing.B, max int) {
	samples := buildTestSamples(max)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		HashKeyFind(samples)
	}
}

func Benchmark100StringKeyFind(b *testing.B) {
	benchmarkStringKeyFind(b, 100)
}
func Benchmark1000StringKeyFind(b *testing.B) {
	benchmarkStringKeyFind(b, 1000)
}
func Benchmark10000StringKeyFind(b *testing.B) {
	benchmarkStringKeyFind(b, 10000)
}
func Benchmark100000StringKeyFind(b *testing.B) {
	benchmarkStringKeyFind(b, 100000)
}
func Benchmark200000StringKeyFind(b *testing.B) {
	benchmarkStringKeyFind(b, 200000)
}
func benchmarkStringKeyFind(b *testing.B, max int) {
	samples := buildTestSamples(max)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StringKeyFind(samples)
	}
}

func buildTestSamples(max int) []ConfigKey {
	randStr := func(n int) string {
		const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.-"
		b := make([]byte, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	var s = make([]ConfigKey, max)
	for i := 0; i < max; i++ {
		k := ConfigKey{
			Kind:      100,
			Name:      randStr(130),
			Namespace: randStr(64),
		}
		s = append(s, k)
	}
	return s
}
