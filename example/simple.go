package main

import (
	"fmt"
	"errors"
	"github.com/synw/terr"
)


func f4() *terr.Trace {
	err := errors.New("Error from f4")
	terr := terr.Debug("f3", err)
	return terr
}

func f3() *terr.Trace {
	err := errors.New("Error from f3")
	perr := f4()
	terr := terr.Minor("f3", err, perr)
	return terr
}

func f2() *terr.Trace {
	err := errors.New("error from f2")
	perr := f3()
	newerr := terr.Critical("f2", err, perr)
	return newerr
}

func f1() *terr.Trace {
	err := errors.New("error from f1")
	perr := f2()
	newerr := terr.New("f1", err, perr)
	return newerr
}

func main() {
	err := f1()
	if err != nil {
		err.Print()
		fmt.Println("--------------------------")
		fmt.Println("With custom separators:")
		err.Print("-> ", "\n")
		fmt.Println("--------------------------")
		fmt.Println("With others custom separators:")
		err.Print("", "> ")
	}
}
