package main

// When you create a type declaration by defining a new type from an existing (non-interface) type,
// you don't inherit the methods defined for that existing type.
//

import "sync"

type myMutex sync.Mutex

func main() {
	var mtx myMutex
	mtx.Lock()   //error
	mtx.Unlock() //error
}
