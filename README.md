# Traced errors

Error tracing library

## Api

**New** (from *string*, errMsg *string*, level *...string*) *Trace : create a trace 
from an error message string

   ```go
   func myfunc() *terr.Trace {
      tr := terr.New("myfunc", "Error one")
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
