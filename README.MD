# LRU Cache Implementation in Go

This repository contains a simple implementation of an LRU (Least Recently Used) cache in Go.

## Concept Overview

LRU cache is a data structure that maintains a fixed-size cache of items, evicting the least recently used item when the cache reaches its capacity. It's useful for scenarios where you want to cache frequently accessed items to improve performance.

## Code Structure

The provided Go code consists of the following components:

- **Node**: Represents a node in the doubly linked list.
- **Queue**: Represents the doubly linked list used in the cache.
- **Cache**: Represents the LRU cache, consisting of the doubly linked list and a hash map.
- **Hash**: Represents a hash map mapping string keys to Node pointers.

The key operations in the LRU cache are:

- **check()**: Checks if an item exists in the cache, removes it if it does, adds it otherwise, and updates the hash map accordingly.
- **Remove()**: Removes a node from the queue and updates the length counter and hash map.
- **Add()**: Adds a node to the queue, removing the least recently used node if the cache size exceeds the predefined size.
- **Display()**: Displays the contents of the cache and the queue.

## Usage

To use the LRU cache implementation:

1. Clone this repository.
2. Navigate to the directory containing the code.
3. Run the Go program using `go run main.go`.

You can modify the `main()` function in `main.go` to test the LRU cache with different input scenarios.

## Example

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start Cache")
	cache := NewCache()
	for _, word := range []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 1", "Item 5", "Item 2"} {
		cache.check(word)
		cache.Display()
	}
}
