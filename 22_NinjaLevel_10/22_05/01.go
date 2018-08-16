package main

import (
	"fmt"
)

/*
https://play.golang.org/p/YHOMV9NYKK
 */
func main() {
	c := make(chan int, 1)

	c <- 10

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
