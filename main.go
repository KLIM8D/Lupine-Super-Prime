package main

import (
	"fmt"
	. "github.com/KLIM8D/Lupine-Super-Prime/utils"
	"math/big"
	"os"
	"time"
)

const (
	DEBUG = false
	COUNT = 10000000
)

var (
	primes []big.Int
)

func main() {
	i := 2
	n := big.NewInt(6)
	k := big.NewInt(1)
	primes := make([]*big.Int, COUNT)
	primes[0] = big.NewInt(2)
	primes[1] = big.NewInt(3)
	primes[2] = big.NewInt(5)

	start := time.Now()
	c := 6
	for c < COUNT {
		n = Add(n, big.NewInt(1))

		debug("(main) n: %v\n", n)
		debug("(main) k: %v\n", k)

		k = findK(n, k)
		j := 0

		for {
			t := primes[j]
			debug("(main) j: %d\n", j)
			debug("(main) t: %v\n", t)
			debug("(main) k: %v\n", k)

			if t == nil {
				debug("(main) breaking loop t was %v\n", nil)
				break
			}

			isGreater := Sub(t, k).Sign()
			debug("(main) isGreater: %v\n", isGreater)
			if isGreater == 1 {
				i++
				primes[i] = n
				break
			}

			mod := Mod(n, t)
			debug("(main) mod: %v\n", mod)

			if mod.Sign() == 0 {
				debug("(main) mod was ZERO", mod)
				break
			}
			j++
		}
		c++
		debug("\n", nil)
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s\n", elapsed)
	fmt.Printf("Number of primes: %d\n", i+1)

	//for _, v := range primes {
	//	if v != nil {
	//		fmt.Println(v)
	//	}
	//}
}

func findK(n, k *big.Int) *big.Int {
	isNegative := Sqrt(n, k)

	for isNegative != -1 {
		k = Add(k, big.NewInt(1))
		isNegative = Sqrt(n, k)
	}

	return k
}

func debug(format string, a ...interface{}) {
	if DEBUG {
		fmt.Fprintf(os.Stdout, format, a...)
	}
}
