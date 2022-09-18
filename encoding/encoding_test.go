package encoding

import "testing"

func TestGolumb(t *testing.T) {

	data := []byte{2, 3, 0, 1, 1, 6, 2, 4, 4, 4, 1, 3, 5, 2}
	got := Golomb(4, data)
	want := []string{"110", "111", "100", "101", "101", "0110", "110", "0100", "0100", "0100", "101", "111", "0101", "110"}

	for y, x := range got {
		if x != want[y] {
			t.Errorf("got %s, wanted %s", x, want[y])
		}
	}
}

func TestEliasGamma(t *testing.T) {

	data := []byte{2, 3, 0, 1, 1, 6, 2, 4, 4, 4, 1, 3, 5, 2}
	got := EliasGamma(data)
	want := []string{"011", "00100", "1", "010", "010", "00111", "011", "00101", "00101", "00101", "010", "00100", "00110", "011"}

	for y, x := range got {
		if x != want[y] {
			t.Errorf("got %s, wanted %s", x, want[y])
		}
	}
}

func TestUnario(t *testing.T) {
	data := []byte{2, 3, 0, 1, 1, 6, 2, 4, 4, 4, 1, 3, 5, 2}
	got := Unario(data)
	want := []string{"001", "0001", "1", "01", "01", "0000001", "001", "00001", "00001", "00001", "01", "0001", "000001", "001"}

	for y, x := range got {
		if x != want[y] {
			t.Errorf("got %s, wanted %s", x, want[y])
		}
	}
}

func TestFibonacci(t *testing.T) {
	data := []byte{3, 4, 1, 2, 2, 7, 3, 5, 5, 5, 2, 4, 6, 3}
	got := Fibonacci(data)
	want := []string{"0011", "1011", "11", "011", "011", "01011", "0011", "00011", "00011", "00011", "011", "1011", "10011", "0011"}

	for y, x := range got {
		if x != want[y] {
			t.Errorf("got %s, wanted %s", x, want[y])
		}
	}
}
