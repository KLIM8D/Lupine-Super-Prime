package lib

import (
	"github.com/KLIM8D/Lupine-Super-Prime/utils"
	"math/big"
	"sync"
)

type Base struct {
	Work     chan PrimeCalc
	Done     chan PrimeCalc
	Primes   []*big.Int
	PIndex   uint64
	NWorkers uint32

	LowestKey *big.Int
	KeyMutex  *sync.Mutex

	PrevEnd      *big.Int
	PrevEndMutex *sync.Mutex

	Factory *utils.RedisConf
}

type Scheduler struct {
	Base  *Base
	Queue *MinHeap
}

type PrimeCalc struct {
	Id     float64
	Start  big.Int
	End    big.Int
	Result []*big.Int
}
