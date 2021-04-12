package main

import (
	"fmt"
	"time"
)

func main() {
	// time.RFC3339: "2006-01-02T15:04:05Z07:00"
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// 0   1   2  3  4  5   7          6
	t, err := time.Parse(time.RFC3339, "2019-05-02T10:04:00Z")
	fmt.Println(t)

	t, err = time.Parse("2006-01-DUPA-02", "2019-05-DUPA-10")
	fmt.Println(t)

	t, err = time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println(err)

	t, err = time.Parse("2019-05-10T10:04:00", "2019-05-10T10:04:00")
	fmt.Println(err)

	// time.Kitchen: "3:04PM"
	t, err = time.Parse(time.Kitchen, "10:04AM")
	fmt.Println(t.Format(time.Kitchen))
}
