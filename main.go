package main

import (
	"fmt"
	"github.com/KLIM8D/Lupine-Super-Prime/lib"
	"github.com/KLIM8D/Lupine-Super-Prime/utils"
	"math/big"
	"os"
	"os/signal"
	"runtime"
	"sync"
)

const (
	BUFFERSIZE = 128 //How large the pool is which holds the work
	//NSCHEDULERS = 1   //Number of schedulers to spawn
	NWORKERS   = 1 //Number of workers which should be spawned by the scheulder
	CONFIGFILE = "config.json"
)

func main() {
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	//Read config file
	conf := &utils.Configuration{ConfigPath: CONFIGFILE}
	conf = conf.Init()

	fmt.Println(conf.Redis)

	base := &lib.Base{NWorkers: NWORKERS, Factory: conf.Redis}
	base.Work = make(chan lib.PrimeCalc, BUFFERSIZE)
	base.Done = make(chan lib.PrimeCalc, BUFFERSIZE)

	//Remove later
	base.Primes = append(base.Primes, big.NewInt(2))
	base.Primes = append(base.Primes, big.NewInt(3))
	base.Primes = append(base.Primes, big.NewInt(5))

	base.LowestKey = big.NewInt(0)
	base.KeyMutex = &sync.Mutex{}

	base.PrevEnd = big.NewInt(7)
	base.PrevEndMutex = &sync.Mutex{}

	for i := 0; i < conf.Scheduler.Amount; i++ {
		s := &lib.Scheduler{Base: base, IsMaster: conf.Scheduler.Master}
		s.Key = 0
		s.NewPrimes = make(chan bool)
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
