package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex

func dupa() {
	fmt.Println("start dupa")

	lock.RLock()
	defer lock.RUnlock()
	fmt.Println("Rlocked")

	time.Sleep(1 * time.Second)

	lock.Lock()
	fmt.Println("RW locked")
	lock.Unlock()
	fmt.Println("RW unlocked")

	fmt.Println("stop dupa")
}

func main() {
	go dupa()
	go dupa()

	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
