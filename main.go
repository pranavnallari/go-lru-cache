package main

import (
	"fmt"
)

const SIZE = 5

// Node represents a node in the doubly linked list
type Node struct {
	left  *Node
	right *Node
	val   string
}

// Queue represents the doubly linked list
type Queue struct {
	Head   *Node
	Tail   *Node
	length int
}

// Cache represents the LRU cache
type Cache struct {
	Queue Queue
	Hash  Hash
}

// NewCache initializes a new Cache
func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

// NewQueue initializes a new Queue
func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.right = tail
	tail.left = head
	return Queue{Head: head, Tail: tail, length: 0}
}

// check checks whether an item exists in the cache, and updates the cache accordingly
func (c *Cache) check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node.val = str
	}
	c.Add(node)
	c.Hash[str] = node
}

// Remove removes a node from the queue
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove : %v\n", n.val)
	left := n.left
	right := n.right
	right.left = left
	left.right = right
	c.Queue.length -= 1
	delete(c.Hash, n.val)
	return n
}

// Add adds a node to the queue
func (c *Cache) Add(n *Node) {
	fmt.Printf("Add : %s \n", n.val)
	tmp := c.Queue.Head.right
	c.Queue.Head.right = n
	n.left = c.Queue.Head
	n.right = tmp
	tmp.left = n
	c.Queue.length++
	if c.Queue.length > SIZE {
		c.Remove(c.Queue.Tail.left)
	}
}

// Display displays the contents of the cache
func (c *Cache) Display() {
	c.Queue.Display()
}

// Display displays the contents of the queue
func (q *Queue) Display() {
	node := q.Head.right
	fmt.Printf("%d - [", q.length)
	for i := 0; i < q.length; i++ {
		fmt.Printf("{%s}", node.val)
		if i < q.length-1 {
			fmt.Printf("<-->")
		}
		node = node.right
	}
	fmt.Println("]")
}

// Hash represents a hash map mapping string keys to Node pointers
type Hash map[string]*Node

func main() {
	fmt.Println("Start Cache")
	cache := NewCache()
	for _, word := range []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 1", "Item 5", "Item 2"} {
		cache.check(word)
		cache.Display()
	}
}
