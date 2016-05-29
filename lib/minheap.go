package lib

import (
	"github.com/KLIM8D/Lupine-Super-Prime/utils"
)

type MinHeap []PrimeCalc

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return utils.Sub(&h[j].Start, &h[i].Start).Sign() == 1 }

//func (h MinHeap) Less(i, j int) bool { return uti h[i].Id < h[j].Id }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() *PrimeCalc {
	if h.Len() > 0 {
		return &h[0]
	} else {
		return nil
	}
}

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(PrimeCalc))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
