# Traced errors

Propagate detailled errors up the call stack

## Data structure

   ```go
   type Terr struct {
      From  string
      Level string
	  Error error
	  File  string
	  Line  int
   }

   type Trace struct {
	  Errors []*Terr
   }
   ```

## Api

**New** (from *string*, errMsg *string*, level *...string*) *Trace : create a trace 
from an error or a string

   ```go
   func myfunc() *terr.Trace {
      tr := terr.New("myfunc", "Error one")
      return tr
   }
   
   func myfunc_() *terr.Trace {
      err := errors.New("Error one")
      tr := terr.New("myfunc", err)
      return tr
   }
   ```
   
**Trace.Add** (from *string*, errMsg *string*, level *...string*) *Trace : adds to a trace 
from an error message string

   ```go
   func myfunc2() *terr.Trace {
      tr := myfunc()
      tr := tr.Add("myfunc2", "Error two", "warning")
      return tr
   }
   ```
   
**Trace.Check**: prints the trace if some errors are present

   ```go
   tr := myfunc2()
   tr.Check()
   ```
   
**Trace.Err**: returns the trace as a standard error

   ```go
   tr := myfunc2()
   err := tr.Err()
   ```
   
**Trace.Msg**: returns the trace's coloured error message

   ```go
   tr := myfunc2()
   fmt.Println(tr.Msg())
   ```
   
**Trace.Print**: print the trace's coloured error message

   ```go
   tr := myfunc2()
   tr.Print()
   ```
   
**Trace.Stack** (from *string*, errMsg *string*, level *...string*): 
same as Trace.Add but adds the stack trace message of the error in the message 

   ```go
   func myfunc3() *terr.Trace {
      tr := myfunc2()
      tr := tr.Stack("myfunc3", "Error two", "debug")
      return tr
   }
   ```
   
**Trace.Fatal** (from *string*, errMsg *string*): 
same as Trace.Stack and exits the program 

   ```go
   func myfunc4() *terr.Trace {
      tr := myfunc3()
      tr := tr.Fatal("myfunc4", "Error two")
      return tr
   }
   ```
   
### Error levels

error (default), minor, warning, debug, fatal

### Examples 
   
Check the [examples](https://github.com/synw/terr/tree/master/example)

   ```go
   package main

   import (
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
	  tr.Check()
   }
   ```
   
Output:
   
   ```
   [error 0] Second error from f2 line 14 in /home/me/terr/example/simple.go
   [error 1] First error from f1 line 8 in /home/me/terr/example/simple.go
   ```
