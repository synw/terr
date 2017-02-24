# Traced errors

Error tracing library

## Usage

   ```go
package main

import (
	"errors"
	"github.com/synw/terr"
)

func f2() *terr.Trace {
	err := errors.New("error from f2")
	newerr := terr.Debug("f2", err)
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
	}
}
```

Start tracing errors:

   ```go
// terr(from string, terr error)
trace := terr.New("function_path", err)
return trace
//return a *Trace instead of an error
   ```

Continue tracing:

   ```go
// trace is the previous returned *Trace
terr.New("function_path", err, trace)
   ```

## Options

   ```go
terr.Critical("function_path", err)
terr.Minor("function_path", err)
terr.Debug("function_path", err)
   ```

Custom formating:
   ```go
// trace is a *Trace
// trace.Print(prefix, suffix) or trace.Print(prefix)
trace.Print("->", "\n")
// get the trace output without printing
formated_trace := trace.Format()
   ```
