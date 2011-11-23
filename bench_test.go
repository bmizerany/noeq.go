package noeq

import (
	"testing"
)

func BenchmarkMultiGen(b *testing.B) {
	b.StopTimer()
	cn, _ := New("localhost:4444")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := cn.Gen(100)
		if err != nil {
			panic(err)
		}
	}
}
