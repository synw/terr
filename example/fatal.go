package main

import (
	"fmt"
	"github.com/synw/terr"
)

func f1() *terr.Trace {
	tr := terr.New("First error")
	return tr
}

func f2() {
	tr := f1()
	tr.Fatal("Second error")
}

func main() {
	f2()
	fmt.Println("I am not supposed to be printed: the program has stopped")
}
