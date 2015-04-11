package main

import (
	"fmt"
	"github.com/KLIM8D/Lupine-Super-Prime/lib"
	"math/big"
	"os"
	"os/signal"
	"sync"
)

const (
	BUFFERSIZE  = 128
	NSCHEDULERS = 1
	NWORKERS    = 1
	NPRIMES     = 10000
)

func main() {
	base := &lib.Base{NWorkers: NWORKERS}
	base.Work = make(chan lib.PrimeCalc, BUFFERSIZE)
	base.Done = make(chan lib.PrimeCalc, BUFFERSIZE)

	base.Primes = make([]*big.Int, NPRIMES)
	base.Primes[0] = big.NewInt(2)
	base.Primes[1] = big.NewInt(3)
	base.Primes[2] = big.NewInt(5)

	base.LowestKey = 0
	base.KeyMutex = &sync.Mutex{}

	base.PrevEnd = big.NewInt(7)
	base.PrevEndMutex = &sync.Mutex{}

	for i := 0; i < NSCHEDULERS; i++ {
		s := &lib.Scheduler{Base: base}
		go s.Start()
	}

	base.SpawnWorkers()

	done := make(chan bool)
	// capture ctrl+c and perform clean-up
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println()
			fmt.Printf("captured %v, exiting...\n", sig)
			os.Exit(0)
		}
		done <- true
	}()
	<-done
}
