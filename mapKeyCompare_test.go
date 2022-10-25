package samples

import (
	"math/rand"
	"testing"
)

func Benchmark100Md5KeyFind(b *testing.B) {
	benchmarkMd5KeyFind(b, 100)
}
func Benchmark1000Md5KeyFind(b *testing.B) {
	benchmarkMd5KeyFind(b, 1000)
}
func Benchmark10000Md5KeyFind(b *testing.B) {
	benchmarkMd5KeyFind(b, 10000)
}
func Benchmark100000Md5KeyFind(b *testing.B) {
	benchmarkMd5KeyFind(b, 100000)
}
func Benchmark200000Md5KeyFind(b *testing.B) {
	benchmarkMd5KeyFind(b, 200000)
}
func benchmarkMd5KeyFind(b *testing.B, max int) {
	samples := buildTestSamples(max)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Md5KeyFind(samples)
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
