package Pasc

import (
	"testing"

	"math/big"

	"reflect"
)

func TestPascelsTri(t *testing.T) {
	var in int = 4
	var want = []*big.Int{big.NewInt(1), big.NewInt(4), big.NewInt(6), big.NewInt(4), big.NewInt(1)}
	got := PascelsTri(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PascelsTri(%d) == %d, want %d.\n", in, got, want)
	}
}
