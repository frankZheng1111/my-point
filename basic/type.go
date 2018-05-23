package main

// When you create a type declaration by defining a new type from an existing (non-interface) type, 
// you don't inherit the methods defined for that existing type.
//

import "sync"
//
// type myMutex sync.Mutex
//
// func main() {  
//   var mtx myMutex
//   mtx.Lock() //error
//   mtx.Unlock() //error  
// }


// If you do need the methods from the original type you can define a new struct type embedding the original type as an anonymous field.
// Interface type declarations also retain their method sets

type myLocker struct {  
  sync.Mutex
}

// type myLocker sync.Locker

func main() {  
  var lock myLocker
  lock.Lock() //ok
  lock.Unlock() //ok
}
