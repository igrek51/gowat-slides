package main

import (
	"fmt"
	"time"
)

type Dupa struct {
	Name string
}

func whatis(v interface{}) string {
	return fmt.Sprintf("value:%v; pointer: %p", v, v)
}

func main() {
	dupas := []Dupa{
		Dupa{Name: "dupa1"},
		Dupa{Name: "dupa2"},
		Dupa{Name: "dupa3"},
	}

	fmt.Println("dupas =", whatis(dupas))

	range_memory_reuse(dupas)

	time.Sleep(2 * time.Second)
}

func range_memory_reuse(dupas []Dupa) {
	for _, dupa := range dupas {
		fmt.Println("dupa =", whatis(dupa))
		print_deferred(&dupa)
	}
}

func range_memory_reuse_fix(dupas []Dupa) {
	for i := range dupas {
		dupa := dupas[i]
		fmt.Println("dupa =", whatis(dupa))
		print_deferred(&dupa)
	}
}

func range_memory_reuse_fix2(dupas []Dupa) {
	for _, d := range dupas {
		dupa := d
		fmt.Println("dupa =", whatis(dupa))
		print_deferred(&dupa)
	}
}

func print_deferred(dupa *Dupa) {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("go dupa =", whatis(dupa))
	}()
}
