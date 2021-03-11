package Pasc

import "math/big"

func PascelsTri(x int) []int {
	pt := make([]int, x)
	if x == 0 {
		pt[0] = 1
		return pt
	}
	for i := 1; i <= x; i++ {
		if i == 1 {
			pt[0] = 1
			pt[1] = 1
		} else {
			DpT := make([]int, x+1)
			for j := 0; j < i+1; j++ {
				//first number of new line
				if j == 0 {
					DpT[j] = 1
				} else if j == 1 { //second
					DpT[j] = i
				} else if j == i-1 { //second to last
					DpT[j] = i
				} else if j == i { //last number of new line
					DpT[j] = 1
				} else { //otherwise calculate regularly
					DpT[j] = pt[j] + pt[j-1]
				}
			}
			pt = DpT
		}
	}
	return pt
}

//This function take in an exponant and returns a list of Big Integers that are the coefficients in Pascel's Triangle.
func PascelsTriBig(x int) []*big.Int {
	pt := make([]*big.Int, x)
	if x == 0 {
		pt[0] = big.NewInt(1)
		return pt
	}
	for i := 1; i <= x; i++ {
		if i == 1 {
			pt[0] = big.NewInt(1)
			pt[1] = big.NewInt(1)
		} else {
			DpT := make([]*big.Int, x+1)
			for j := 0; j < i+1; j++ {
				//first number of new line
				if j == 0 {
					DpT[j] = big.NewInt(1)
				} else if j == 1 { //second
					DpT[j] = big.NewInt(int64(i))
				} else if j == i-1 { //second to last
					DpT[j] = big.NewInt(int64(i))
				} else if j == i { //last number of new line
					DpT[j] = big.NewInt(1)
				} else { //otherwise calculate regularly
					var c big.Int
					DpT[j] = c.Add(pt[j], pt[j-1])
				}
			}
			pt = DpT
		}
	}
	return pt
}
