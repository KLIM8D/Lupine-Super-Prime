package scheduler

import (
	"math/big"
)

type Scheduler struct {
	Work      chan Worker
	Done      chan Worker
	LowestKey float64
}

type Worker struct {
	Start  float64
	End    float64
	Result []Prime
}

type Prime struct {
	Index float64
	Value big.Int
}
