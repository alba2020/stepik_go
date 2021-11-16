package main

import "testing"

func BenchmarkConvertFormat(b *testing.B) {
	for n := 0; n < b.N; n++ {
			ConvertFormat(10)
	}
}

func BenchmarkConvertString(b *testing.B) {
	for n := 0; n < b.N; n++ {
			ConvertString(10)
	}
}
