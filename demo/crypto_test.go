package main

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

func BenchmarkSHA256(b *testing.B) {
	hasher := sha256.New()
	bs := make([]byte, 1024)
	_, err := rand.Read(bs)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasher.Write(bs)
		hasher.Sum(nil)
	}
}
