package lib

import (
	"fmt"
	. "github.com/KLIM8D/Lupine-Super-Prime/utils"
	"math/big"
	"os"
	"time"
)

const (
	DEBUG = true
)

func (b *Base) SpawnWorkers() {
	for i := 0.0; i < b.NWorkers; i++ {
		go b.CalculatePrime()
	}
}

func (b *Base) CalculatePrime() {
	for {
		w := <-b.Work
		done := false
		k := big.NewInt(1)
		n := &w.Start
		i := 0

		for !done {
			debug("(calc) n: %v\n", n)
			debug("(calc) k: %v\n", k)

			k = findK(n, k)
			j := 0

			for {
				t := b.Primes[j]
				debug("(calc) j: %d\n", j)
				debug("(calc) t: %v\n", t)
				debug("(calc) k: %v\n", k)
				if t == nil {
					debug("(calc) breaking loop t was %v\n", nil)
					break
				}

				isGreater := Sub(t, k).Sign()
				debug("(calc) isGreater: %v\n", isGreater)
				if isGreater == 1 {
					i++
					w.Result[i] = n
					break
				}

				mod := Mod(n, t)
				debug("(calc) mod: %v\n", mod)

				if mod.Sign() == 0 {
					debug("(calc) mod was ZERO", mod)
					break
				}
				j++
			}
			n = Add(n, big.NewInt(1))
			done = Sub(&w.End, n).Sign() != 1

			if DEBUG {
				fmt.Printf("(calc) w.End: %v\n", w.End)
				fmt.Printf("(calc) n: %v\n", n)
				fmt.Printf("(calc) done: %v\n", done)
				time.Sleep(1 * time.Second)
			}
		}
		b.Done <- w
		time.Sleep(1 * time.Second)
		break
	}
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
