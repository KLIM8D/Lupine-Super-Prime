package lib

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	list = &CircularList{}
)

const (
	VERBOSE = false
)

func TestInsert(t *testing.T) {
	const START = 1
	const LIMIT = 100

	list.Init()

	for i := START; i <= LIMIT; i++ {
		list.Insert(&Node{Value: i})
	}

	e := list.Tail()
	for i := list.Len(); i > 0; i-- {
		e = e.Prev()
	}

	if VERBOSE {
		log.Printf("mem: %p -   -1: %v\n", list.Head().Prev(), list.Head().Prev())
		log.Printf("mem: %p - Head: %v\n", list.Head(), list.Head())
		log.Printf("mem: %p -   +1: %v\n", list.Head().Next(), list.Head().Next())
		log.Printf("mem: %p -   -1: %v\n", list.Tail().Prev(), list.Tail().Prev())
		log.Printf("mem: %p - Tail: %v\n", list.Tail(), list.Tail())
		log.Printf("mem: %p -   +1: %v\n", list.Tail().Next(), list.Tail().Next())
		log.Printf("Len: %d\n", list.Len())
	}

	assert.EqualValues(t, list.Len(), LIMIT, "Wrong length returned!")
	assert.EqualValues(t, list.Tail().Value, LIMIT, "Tail's value does not match LIMIT")
	assert.EqualValues(t, list.Head().Value, START, "Head's value does not match START")
	assert.Equal(t, list.Tail().Next(), list.Head(), "Tail's next does not point to Head")
	assert.Equal(t, list.Head().Prev(), list.Tail(), "Head's prev does not point to Tail")

	e = list.Head()
	o := e.Next()
	for i := uint64(1); i < list.Len(); i++ {
		assert.Equal(t, o.Prev(), e, "Node's prev does not match previous node")
		e, o = e.Next(), o.Next()
	}
}

func TestRemove(t *testing.T) {
	for i := list.Len(); i > 0; i-- {
		e := list.Tail()
		rm := list.Remove()
		prev := e.Prev()
		assert.Equal(t, e, rm, "Remove did not remove the Tail node")
		assert.NotEqual(t, rm, list.Tail(), "Node was not removed from the list")
		assert.Equal(t, prev, list.Tail(), "Old Tail's previous node is not equal to the new Tail")
	}
}
