package common

import (
	"runtime"
	"testing"
)

func testParallelN(enabled bool, n, procs int) bool {
	data := make([]bool, n)
	before := runtime.GOMAXPROCS(0)
	runtime.GOMAXPROCS(procs)
	parallel(n, func(start, end int) {
		for i := start; i < end; i++ {
			data[i] = true
		}
	})
	for i := 0; i < n; i++ {
		if data[i] != true {
			return false
		}
	}
	runtime.GOMAXPROCS(before)
	return true
}

func TestParallel(t *testing.T) {
	for _, e := range []bool{true, false} {
		for _, n := range []int{1, 10, 100, 1000} {
			for _, p := range []int{1, 2, 4, 8, 16, 100} {
				if testParallelN(e, n, p) != true {
					t.Errorf("test [parallel %v %d %d] failed", e, n, p)
				}
			}
		}
	}
}
