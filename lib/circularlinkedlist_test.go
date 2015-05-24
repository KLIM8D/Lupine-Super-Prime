package lib

import (
	"log"
	"testing"
)

var (
	list = &CircularList{}
)

func TestInsert(t *testing.T) {
	list.Init()
	list.Insert(&Node{Value: 1})
	list.Insert(&Node{Value: 2})
	list.Insert(&Node{Value: 3})

	e := list.Tail()
	for i := list.Len(); i > 0; i-- {
		log.Printf("Node: %v\n", e)
		e = e.Prev()
	}

	log.Printf("Head: %v\n", list.Head())
	log.Printf("Tail: %v\n", list.Tail())
}
