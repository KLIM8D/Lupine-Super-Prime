package lib

import (
	"container/heap"
	. "github.com/KLIM8D/Lupine-Super-Prime/utils"
	"math/big"
)

const (
	WORKSIZE = 10
)

func (s *Scheduler) Start() {
	s.Queue = &MinHeap{}
	heap.Init(s.Queue)
	w := s.CreateWork()

	for {
		select {
		case r := <-s.Base.Done:
			heap.Push(s.Queue, r)
			debug("(sche) Recv done work\n --- Start: %v\n --- End: %v\n --- Result: %v\n", r.Start, r.End, r.Result)
		case s.Base.Work <- w:
			w = s.CreateWork()
			//debug("(sche) Adding work\n --- Start: %v\n --- End: %v\n", w.Start, w.End)
		}
	}
}

func (s *Scheduler) CreateWork() PrimeCalc {
	s.Base.PrevEndMutex.Lock()
	start := *s.Base.PrevEnd
	s.Base.PrevEnd = Add(s.Base.PrevEnd, big.NewInt(WORKSIZE))
	end := *s.Base.PrevEnd
	s.Base.PrevEndMutex.Unlock()

	w := PrimeCalc{Start: start, End: end}
	w.Result = make([]*big.Int, WORKSIZE)

	return w
}
