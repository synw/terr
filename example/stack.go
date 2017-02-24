package main

import (
	"fmt"
	"time"
	"errors"
	"strconv"
	"github.com/synw/terr"
)


func doWork() *terr.Trace {
	trace := &terr.Trace{}
	for i := 1;  i<=10; i++ {
		tick := strconv.Itoa(i)
		fmt.Println("Work being done: tick", tick)
		if i == 2 || i == 5 {
			err := errors.New("Error")
			from := "doWork: tick "+tick
			trace = terr.Stack(from, err, trace)
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
