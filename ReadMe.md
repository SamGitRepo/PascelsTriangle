This is a small module that allows you to input an exponant and will return a Big Integer list of the corresponding coefficients in Pascel's Triangle.

example:
(a+b)^2 = a^2 + 2ab + b^2

code:
Pasc.PascelsTri(2) = []int{1, 1}

Pasc.PascelsTriBig(2) = []*big.Int{big.NewInt(1), big.NewInt(1)}

there are two functions: PascelsTri and PascelsTriBig. PascelsTri returns and []int while PascelsTriBig returns []*big.Int. PascelsTriBig is for numbers too high for an int to store.
