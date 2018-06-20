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

**New** (errObj *interface{}*, level *...string*) *Trace : create a trace 
from an error or a string

   ```go
   func foo() *terr.Trace {
      tr := terr.New("Error one")
      return tr
   }
   
   func bar() *terr.Trace {
      err := errors.New("Error one")
      tr := terr.New(err)
      return tr
   }
   ```
   
**Trace.Add** (errObj *interface{}*, level *...string*) *Trace : adds to a trace 
from an error message string

   ```go
   func myfunc2() *terr.Trace {
      tr := myfunc()
      tr := tr.Add("Error two", "warning")
      return tr
   }
   ```
   
**Trace.Pass** (level *...string*) *Trace : adds a trace with no new error message: it just records
the function call

   ```go
   tr := myfunc2()
   tr = tr.Pass()
   ```
   
**Trace.Check**: prints the trace if some errors are present

   ```go
   tr := myfunc2()
   tr.Check()
   ```
   
**Trace.Err** () error: returns the full trace as a standard error

   ```go
   tr := myfunc2()
   err := tr.Err()
   ```
   
**Trace.Print**: print the trace's coloured error message

   ```go
   tr := myfunc2()
   tr.Print()
   ```
   
**Trace.Msg** () string: returns the trace's coloured error message

   ```go
   tr := myfunc2()
   fmt.Println(tr.Msg())
   ```
   
**Trace.Log** () string: returns the last error message of the trace 

   ```go
   tr := myfunc2()
   tr.Log()
   ```
   
**Trace.Stack** (errObj *interface{}*, level *...string*): 
same as Trace.Add but adds the stack trace message of the error in the message 

   ```go
   func myfunc3() *terr.Trace {
      tr := myfunc2()
      tr := tr.Stack("Error two", "debug")
      return tr
   }
   ```
   
**Trace.Fatal** (errObj *interface{}*, level *...string*): 
same as Trace.Stack and exits the program 

   ```go
   func myfunc4() *terr.Trace {
      tr := myfunc3()
      tr := tr.Fatal("Error two")
      return tr
   }
   ```
   
### Error levels

debug, info, warning, error (default), fatal, panic

### Examples 
   
Check the [examples](https://github.com/synw/terr/tree/master/example)

   ```go
   package main

   import (
      "errors"
	  "github.com/synw/terr"
   )

   func f1() *terr.Trace {
	  tr := terr.New("First error")
	  return tr
   }

   func f2() *terr.Trace {
	  tr := f1()
	  err := errors.New("Second error")
	  tr = tr.Add(err)
	  return tr
   }

   func main() {
	  tr := f2()
	  tr.Check()
   }
   ```
   
Output:
   
   ```
   0 [error ] Second error from main.f2 line 18 in /home/me/terr/example/simple.go
   1 [error ] First error from main.f1 line 11 in /home/me/terr/example/simple.go
   ```
