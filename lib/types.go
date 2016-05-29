package lib

import (
	"github.com/KLIM8D/Lupine-Super-Prime/utils"
	"math/big"
	"sync"
)

type Base struct {
	Work     chan PrimeCalc
	Done     chan PrimeCalc
	SockRecv chan bool
	Primes   []*big.Int
	NWorkers uint32

	LowestKey *big.Int
	KeyMutex  *sync.Mutex

	PrevEnd      *big.Int
	PrevEndMutex *sync.Mutex

	Factory *utils.RedisConf
}

type Scheduler struct {
	Base      *Base
	Queue     *MinHeap
	NewPrimes chan bool
	Key       int64
	IsMaster  bool
}

type PrimeCalc struct {
	Id     float64
	Start  big.Int
	End    big.Int
	Result []*big.Int
}

type NetworkSync struct {
	masters []SyncHost
	slaves  []SyncHost
}

type SyncHost struct {
	ip   string
	port int16
}
