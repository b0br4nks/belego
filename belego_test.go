package belego_test

import (
    "testing"
    
    "github.com/b0br4nks/belego"
)

var (
    big = belego.Big
    little = belego.Little
    want = belego.Endian()
    got = belego.Endian()
)

// TestEndian tests the Endian function.
// It runs the function and compares the result with the expected value.
func TestEndian(t *testing.T) {
    if want != got {
        t.Errorf("want %v, got %v", want, got)
    }
}

// BenchmarkEndian benchmarks the Endian function.
// It runs the function b.N times, and reports the time it took to run the function.
// go test -bench=. -benchmem
func BenchmarkEndian(b *testing.B) {
    for i := 0; i < b.N; i++ {
        belego.Endian()
    }
}

