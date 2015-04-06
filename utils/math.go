package utils

import "math/big"

func Sub(x, y *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Sub(x, y)
	return r
}

func Mul(x, y *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mul(x, y)

	return r
}

func Mod(x, y *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mod(x, y)

	return r
}

func Add(x, y *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Add(x, y)

	return r
}

func Sqrt(x, y *big.Int) int {
	k2 := Mul(y, y)
	sub := Sub(x, k2)
	isNegative := sub.Sign()

	return isNegative
}
