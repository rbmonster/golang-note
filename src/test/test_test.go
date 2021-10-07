package main

import "testing"

func TestWheel_Write(t *testing.T) {

	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 1, Y: 2},
			Redius: 10,
		},
		Spokes: 12,
		owner:  "12123",
	}

	_, err := w2.Write([]byte("asdfa"))
	if err != nil {
		t.Errorf("error")
	}
}

func BenchmarkWheel_Write(b *testing.B) {
	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 1, Y: 2},
			Redius: 10,
		},
		Spokes: 12,
		owner:  "12123",
	}
	// b.N 表示调用N次
	for i := 0; i < b.N; i++ {
		_, err := w2.Write([]byte("asdfa"))
		if err != nil {
			b.Errorf("error")
		}
	}
}
