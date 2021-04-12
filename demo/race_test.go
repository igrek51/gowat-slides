package main

import (
	"fmt"
	"testing"
)

func TestRace(t *testing.T) {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
	if m["1"] != "a" {
		t.Errorf("expected a, got %s", m["1"])
	}
}
