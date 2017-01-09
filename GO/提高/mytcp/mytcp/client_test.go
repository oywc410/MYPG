package mytcp

import "testing"

func BenchmarkLineServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go ClientStart()
	}
}

//go test -v -bench=".*" -benchmem