package lib

import "fmt"

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
	root Node
}

func (c *CircularList) Len() uint64 { return c.len }

func (c *CircularList) Init() *CircularList {
	c.root.next = &c.root
	c.root.prev = &c.root
	c.len = 0
	return c
}

func (c *CircularList) Tail() *Node {
	if c.len == 0 {
		return nil
	}
	return &c.root
}

func (c *CircularList) Head() *Node {
	if c.len == 0 {
		return nil
	}
	return c.root.Next()
}

func (c *CircularList) Insert(n *Node) {
	if n == nil {
		return
	}

	if c.len == 0 {
		n.next = n
		n.prev = n
		c.root = *n
	} else {
		//Set n.prev to point to current root (last element)
		n.prev = &c.root

		//Set n.next to point to Head
		n.next = c.root.next

		//Set current root (last element) next's element to be n
		c.root.next = n

		fmt.Printf("mem: %v - cRoot: %v\n", &c.root, c.root)
		fmt.Printf("mem: %v - n: %v\n", &n, n)

		//Set root (last element) to n
		c.root = *n

		//Set Head.prev to point to root (n)
		c.root.next.prev = &c.root
	}
	c.len++
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
	c.root = *t
	c.len--

	fmt.Printf("T: %v\n", t)
	fmt.Printf("E: %v\n", e)

	return &e
}
