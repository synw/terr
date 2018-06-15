package main

import (
	"fmt"
	"github.com/synw/terr"
)

func f1() *terr.Trace {
	tr := terr.New("f1", "First error")
	return tr
}

func f2() *terr.Trace {
	tr := f1()
	tr = tr.Add("f2", "Second error")
	return tr
}

func main() {
	tr := f2()
	err := tr.Err()
	fmt.Println(err)
}
