package lib

import (
	"container/heap"
	. "github.com/KLIM8D/Lupine-Super-Prime/utils"
	"math/big"
	//"sync/atomic"
)

const (
	WORKSIZE = 10
)

var (
	newPrimes chan bool
)

func (s *Scheduler) Start() {
	s.Queue = &MinHeap{}
	heap.Init(s.Queue)
	w := s.CreateWork()

	for {
		select {
		case r := <-s.Base.Done:
			heap.Push(s.Queue, r)
			newPrimes <- true
			debug("(sche) Recv done work\n --- Start: %v\n --- End: %v\n --- Result: %v\n", r.Start, r.End, r.Result)
		case s.Base.Work <- w:
			w = s.CreateWork()
			debug("(sche) Adding work\n --- Start: %v\n --- End: %v\n", w.Start, w.End)
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

/*
Ideas:
Store primes in lists in redis. Each list contains 1000 primes.
The key should just be a int64 or float64, incremented from 1 to N
Problem with concurrency?
*/
func (s *Scheduler) StoreWork() {
	for {
		<-newPrimes
		p := Sub(&s.Queue.Peek().Start, s.Base.LowestKey).Sign()
		for p == 1 {
			//for _, v := range r.Result {
			//	if v != nil {
			//		i = atomic.AddUint64(&s.Base.PIndex, 1)
			//		s.Base.Primes = append(s.Base.Primes, v)
			//		go s.Base.Factory.Add(&RedisItem{Key: i, Value: v})
			//	}
			//}
			p = Sub(&s.Queue.Peek().Start, s.Base.LowestKey).Sign()
		}
	}
}
