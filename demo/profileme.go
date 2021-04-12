package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"
)

func main() {
	cpuProfile, _ := os.Create("cpuprof.prof")
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	go doSomething()
	go doSomething()
	go doSomethingElse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func doSomething() {
	fmt.Println("buckle up")
	doSomethingImportant()
}

func doSomethingElse() {
	fmt.Println("buckle up")
	doSomethingImportant()
}

func doSomethingImportant() {
	ch := make(chan string)
	for {
		select {
		case <-ch:
			fmt.Println("received")
		default:
		}
		time.Sleep(1) // BTW: bad practice
	}
}
