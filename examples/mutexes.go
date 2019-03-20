package main

import "sync"

var mu sync.Mutex

func main() {
	mu.Lock()
	defer mu.Unlock()
	// the lock is held until we return from the function
	// The defer statement is particularly useful when there are multiple return statements. Without defer, we’d need a call to Unlock just before every return statement, and it would be very easy to forget one of those.
}

// Visited tracks whether web pages have been visited.
// Its methods may be used concurrently from multiple goroutines.
type Visited struct {
	// mu guards the visited map.
	mu      sync.Mutex
	visited map[string]int // Declare a map from URL (string) keys to integer values
}

// VisitLink tracks that the page with the given URL has
// been visited, and returns the updated link count.
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}
