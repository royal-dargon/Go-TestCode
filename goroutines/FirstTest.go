/*
package main

type value struct {
	mu sync.Mutex
	value int
}

var wg sync.WaitGroup
printSum := func(v1, v2 *value) {
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()
	time.Sleep(2*time.Second)
	v2.mu.Lock()
	defer v2.mu.Unlock()

}
*/