package lib

import (
	"math/big"
	"sync"
)

type Base struct {
	Work     chan PrimeCalc
	Done     chan PrimeCalc
	Primes   []*big.Int
	NWorkers float64

	LowestKey float64
	KeyMutex  *sync.Mutex

	PrevEnd      *big.Int
	PrevEndMutex *sync.Mutex
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
