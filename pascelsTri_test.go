package Pasc

import (
	"testing"

	"math/big"

	"reflect"
)

func TestPascelsTriBig(t *testing.T) {
	var in int = 4
	var want = []*big.Int{big.NewInt(1), big.NewInt(4), big.NewInt(6), big.NewInt(4), big.NewInt(1)}
	got := PascelsTriBig(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PascelsTriBig(%d) == %d, want %d.\n", in, got, want)
	}
}

func TestPascelsTri(t *testing.T) {
	var in int = 4
	want := []int{1, 4, 6, 4, 1}
	got := PascelsTri(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PascelsTri(%d) == %d, want %d.\n", in, got, want)
	}
}
