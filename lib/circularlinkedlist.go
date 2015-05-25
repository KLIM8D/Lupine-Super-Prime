package lib

import (
	"errors"
)

type Node struct {
	prev  *Node
	next  *Node
	Value interface{}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type CircularList struct {
	len  uint64
	root *Node
}

func (c *CircularList) Len() uint64 { return c.len }

func (c *CircularList) Init() *CircularList {
	c.root = nil
	c.len = 0
	return c
}

func (c *CircularList) Tail() *Node {
	if c.len == 0 {
		return nil
	}
	return c.root
}

func (c *CircularList) Head() *Node {
	if c.len == 0 {
		return nil
	}
	return c.root.Next()
}

func (c *CircularList) Insert(n *Node) error {
	if n == nil {
		return errors.New("Can't insert <nil>")
	}

	if c.len == 0 {
		n.next = n
		n.prev = n
		c.root = n
	} else {
		e := c.root

		//Set newnode to point to current root (last element)
		n.prev = e
		//Set newnode.next to point to Head
		n.next = e.next
		//Set current root (last element) next's element to be n
		e.next = n
		//Set root (last element) to n
		c.root = n
		//Set Head.prev to point to root (n)
		c.root.next.prev = c.root
	}
	c.len++

	return nil
}

func (c *CircularList) Remove() *Node {
	e := c.root
	t := c.root.prev

	//Set e-1 to point to head
	//Set head to point to e-1
	c.root.prev.next = c.root.next
	c.root.next.prev = c.root.prev

	//Avoid memory leaks
	e.next = nil
	e.prev = nil

	//Set root element to e-1
	c.root = t
	c.len--

	return e
}
