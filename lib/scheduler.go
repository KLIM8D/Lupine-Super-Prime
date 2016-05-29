package lib

import (
	"container/heap"
	. "github.com/KLIM8D/Lupine-Super-Prime/utils"
	"math/big"
	//"sync/atomic"
)

const (
	WORKSIZE = 10
	LISTSIZE = 1000
)

var (
	nelements int16
)

func (s *Scheduler) Start() {
	s.Queue = &MinHeap{}
	heap.Init(s.Queue)
	w := s.CreateWork()

	go s.StoreWork()

	for {
		select {
		case r := <-s.Base.Done:
			s.Queue.Push(r)
			s.NewPrimes <- true
			debug("(sche) Recv done work\n --- Start: %v\n --- End: %v\n --- Result: %v\n", r.Start, r.End, r.Result)
		case s.Base.Work <- w:
			if s.IsMaster {
				w = s.CreateWork()
				debug("(sche) Adding work\n --- Start: %v\n --- End: %v\n", w.Start, w.End)
			} else {
				//Sync with master
			}
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
The s.Key should just be a int64 or float64, incremented from 1 to N
Problem with concurrency?
*/
func (s *Scheduler) StoreWork() {
	for {
		<-s.NewPrimes
		p := Sub(&s.Queue.Peek().Start, s.Base.LowestKey).Sign()
		for p == 1 {
			if nelements > LISTSIZE {
				s.Key++
				nelements = 0
			}

			prime := s.Queue.Pop().(PrimeCalc)
			for _, v := range prime.Result {
				if v != nil {
					s.Base.Factory.LPush(s.Key, v.String())
					s.Base.Primes = append(s.Base.Primes, v)
					nelements++
				}
			}

			if qelement := s.Queue.Peek(); qelement != nil {
				p = Sub(&qelement.Start, s.Base.LowestKey).Sign()
			} else {
				p = 0
			}
		}
	}
}
