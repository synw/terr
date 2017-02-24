package main

import (
	"fmt"
	"time"
	"errors"
	"github.com/synw/terr"
)


func doWork() *terr.Trace {
	trace := &terr.Trace{}
	for i := 1;  i<=10; i++ {
		fmt.Println("Work being done")
		if i == 2 || i == 5 {
			err := errors.New("Error")
			trace = terr.Push("doWork", err, trace)
		}
		time.Sleep(1*time.Second)
	}
	return trace
}

func main() {
	trace := doWork()
	fmt.Println("------------- Trace -------------")
	trace.Print()
}
