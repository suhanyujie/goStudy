package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 7},
		{5, 12, 17},
		{8, 15, 23},
		{30000, 40000, 70000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d);"+"got %d;expected %d", tt.a, tt.b, tt.c,actual)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	b.ResetTimer()
	ans := 70000
	for i := 0; i < b.N; i++ {
		actual := calcTriangle(30000,40000)
		if actual != ans {
			b.Errorf("got %d for calcTriangle(%d,%d),except %d", actual,30000,40000,ans)
		}
	}
}
