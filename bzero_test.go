package bzero

import (
	"fmt"
	"testing"
)

func TestBZero(t *testing.T) {
	type testRow struct {
		size  uint
		align uint
	}

	testData := [...]testRow{
		{size: 0},
		{size: 1},
		{size: 15},
		{size: 16},
		{size: 17},
		{size: 63},
		{size: 64},
		{size: 65},
		{size: 127},
		{size: 128},
		{size: 129},
		{size: 255},
		{size: 256},
		{size: 257},
		{size: 1023},
		{size: 1024},
		{size: 1025},
		{size: 0, align: 11},
		{size: 1, align: 11},
		{size: 15, align: 11},
		{size: 16, align: 11},
		{size: 17, align: 11},
		{size: 63, align: 11},
		{size: 64, align: 11},
		{size: 65, align: 11},
		{size: 127, align: 11},
		{size: 128, align: 11},
		{size: 129, align: 11},
		{size: 255, align: 11},
		{size: 256, align: 11},
		{size: 257, align: 11},
		{size: 1023, align: 11},
		{size: 1024, align: 11},
		{size: 1025, align: 11},
	}

	for _, row := range testData {
		name := fmt.Sprintf("%d/%d", row.size, row.align)
		t.Run(name, func(t *testing.T) {
			data := make([]byte, 64+row.align+row.size+64)
			j0 := uint(0)
			j1 := uint(64)
			j2 := j1 + row.align
			j3 := j2 + row.size
			j4 := j3 + 64

			for i := j0; i < j1; i++ {
				data[i] = 'A'
			}
			for i := j1; i < j2; i++ {
				data[i] = 'B'
			}
			for i := j2; i < j3; i++ {
				data[i] = 'C'
			}
			for i := j3; i < j4; i++ {
				data[i] = 'D'
			}

			Uint8(data[j2:j3])

			var overwrites, underwrites uint
			sawNotA := false
			sawNotB := false
			sawNotZ := false
			sawNotD := false
			for i := j0; i < j1; i++ {
				if data[i] == 'A' {
					continue
				}
				if !sawNotA {
					sawNotA = true
					t.Logf("data[%d] is %q", i, data[i])
				}
				overwrites++
			}
			for i := j1; i < j2; i++ {
				if data[i] == 'B' {
					continue
				}
				if !sawNotB {
					sawNotB = true
					t.Logf("data[%d] is %q", i, data[i])
				}
				overwrites++
			}
			for i := j2; i < j3; i++ {
				if data[i] == 0 {
					continue
				}
				if !sawNotZ {
					sawNotZ = true
					t.Logf("data[%d] is %q", i, data[i])
				}
				underwrites++
			}
			for i := j3; i < j4; i++ {
				if data[i] == 'D' {
					continue
				}
				if !sawNotD {
					sawNotD = true
					t.Logf("data[%d] is %q", i, data[i])
				}
				overwrites++
			}
			if underwrites != 0 {
				t.Errorf("saw %d underwrites", underwrites)
			}
			if overwrites != 0 {
				t.Errorf("saw %d overwrites", overwrites)
			}
		})
	}
}

func BenchmarkBZero256(b *testing.B) {
	data := make([]byte, 256)
	for n := 0; n < b.N; n++ {
		Uint8(data)
	}
}

func BenchmarkBZero1K(b *testing.B) {
	data := make([]byte, 1024)
	for n := 0; n < b.N; n++ {
		Uint8(data)
	}
}

func BenchmarkBZero4K(b *testing.B) {
	data := make([]byte, 4096)
	for n := 0; n < b.N; n++ {
		Uint8(data)
	}
}
