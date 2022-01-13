package main

import (
	"sync"
	"testing"
)

var m sync.Mutex

func call() {
	m.Lock()
	m.Unlock()
}

func deferCall() {
	m.Lock()
	defer m.Unlock()
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferCall()
	}
}

/*
   $ go test -bench=. defer_test.go
		goos: darwin
		goarch: amd64
		cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
		BenchmarkCall-4         71156937                16.66 ns/op
		BenchmarkDefer-4        54839784                21.65 ns/op
		PASS
		ok      command-line-arguments  2.422s
*/
