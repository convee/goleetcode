package main

import "fmt"

func chanRange(chanName chan int) {
	for e := range chanName {
		fmt.Printf("Get element from chan: %d\n", e)
	}
}
