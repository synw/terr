package terr

import (
	"errors"
	"fmt"
	"github.com/acmacalister/skittles"
	"os"
	"runtime"
	"strconv"
)

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

// constructor
func New(errObj interface{}, level ...string) *Trace {
	var err error
	err, found := errObj.(error)
	if found == false {
		errMsg, found := errObj.(string)
		if found == true {
			err = errors.New(errMsg)
		} else {
			panic("The second parameter must be a string or an error")
		}
	}
	lvl := "error"
	if len(level) > 0 {
		lvl = level[0]
	}

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	from := f.Name()

	//_, file, line, _ := runtime.Caller(1)
	ter := &Terr{from, lvl, err, file, line}
	ters := []*Terr{ter}
	tr := Trace{ters}
	return &tr
}

// ------------------- Methods -------------------

// add a new error to the trace
func (trace Trace) Add(errObj interface{}, level ...string) *Trace {
	var err error
	err, found := errObj.(error)
	if found == false {
		errMsg, found := errObj.(string)
		if found == true {
			err = errors.New(errMsg)
		} else {
			panic("The second parameter must be a string or an error")
		}
	}

	lvl := "error"
	if len(level) > 0 {
		lvl = level[0]
	}

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	from := f.Name()

	//_, file, line, _ := runtime.Caller(1)
	ter := &Terr{from, lvl, err, file, line}
	trace.Errors = append(trace.Errors, ter)
	return &trace
}

// add a new error to the trace with no message
func (trace Trace) Pass(level ...string) *Trace {
	lvl := "error"
	if len(level) > 0 {
		lvl = level[0]
	}
	err := errors.New("")

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	from = f.Name()

	//_, file, line, _ := runtime.Caller(1)
	ter := &Terr{from, lvl, err, file, line}
	trace.Errors = append(trace.Errors, ter)
	return &trace
}

// prints the trace if some errors are found
func (trace Trace) Check() {
	if len(trace.Errors) > 0 {
		trace.Print()
	}
}

// print the trace
func (trace Trace) Print() {
	msg := trace.Msg()
	fmt.Println(msg)
}

// get the coloured error message from the trace
func (trace Trace) Msg() string {
	var msg string
	trs := reverse(trace.Errors)
	for i, tr := range trs {
		label := getLabelWithNum(tr, i)
		msg = msg + label + " " + tr.Error.Error() + " from " + skittles.BoldWhite(tr.From)
		msg = msg + " line " + skittles.White(strconv.Itoa(tr.Line)) + " in " + tr.File
		if (i + 1) < len(trace.Errors) {
			msg = msg + "\n"
		}
	}
	return msg
}

// get the error message from the trace
func (trace Trace) Error() string {
	return trace.Err().Error()
}

// returns the trace as a standard error
func (trace Trace) Err() error {
	var msg string
	trs := reverse(trace.Errors)
	if len(trs) > 0 {
		for i, tr := range trs {
			if tr != nil {
				if tr.Error != nil {
					lvl := "Error"
					if tr.Level != "" {
						lvl = tr.Level
					}
					label := "[" + lvl + " " + strconv.Itoa(i) + "]"
					msg = msg + label + " " + tr.Error.Error() + " from " + tr.From
					msg = msg + " line " + strconv.Itoa(tr.Line) + " in " + tr.File
					if i+1 < len(trace.Errors) {
						msg = msg + ("\n")
					}
				}
			}
		}
	}
	err := errors.New(msg)
	return err
}

// same as Add but the error message will contain the stack trace
func (trace Trace) Stack(errObj interface{}, level ...string) *Trace {
	var err error
	err, found := errObj.(error)
	if found == false {
		errMsg, found := errObj.(string)
		if found == true {
			err = errors.New(errMsg)
		} else {
			panic("The second parameter must be a string or an error")
		}
	}

	lvl := "error"
	if len(level) > 0 {
		lvl = level[0]
	}

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	from := f.Name()

	//_, file, line, _ := runtime.Caller(1)
	var stack [4096]byte
	runtime.Stack(stack[:], false)
	st := fmt.Sprintf("%s\n", stack[:])
	err = errors.New(err.Error() + "\n" + st)
	ter := &Terr{from, lvl, err, file, line}
	trace.Errors = append(trace.Errors, ter)
	return &trace
}

// same as Stack but stops the program
func (trace Trace) Fatal(errObj interface{}) {
	var err error
	err, found := errObj.(error)
	if found == false {
		errMsg, found := errObj.(string)
		if found == true {
			err = errors.New(errMsg)
		} else {
			panic("The second parameter must be a string or an error")
		}
	}

	lvl := "fatal"
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	from := f.Name()

	//_, file, line, _ := runtime.Caller(1)
	var stack [4096]byte
	runtime.Stack(stack[:], false)
	st := fmt.Sprintf("%s\n", stack[:])
	err = errors.New(err.Error() + "\n" + st)
	ter := &Terr{from, lvl, err, file, line}
	trace.Errors = append(trace.Errors, ter)
	trace.Print()
	os.Exit(1)
}

// ------------------- Functions -------------------

// print debug info about something
func Debug(args ...interface{}) {
	num_args := len(args)
	if num_args < 1 {
		return
	}
	for _, o := range args {
		msg := "[" + skittles.Yellow("debug") + "] "
		fmt.Println(msg + fmt.Sprintf("Type: %T Value: %#v", o, o))
	}
}

// get the colored label and error number
func getLabelWithNum(tr *Terr, i int) string {
	s := strconv.Itoa(i)
	label := s + " [" + skittles.Red(tr.Level) + "]"
	if tr.Level == "debug" {
		label = s + " [" + skittles.Yellow(tr.Level) + "]"
	} else if tr.Level == "info" {
		label = s + " [" + skittles.Green(tr.Level) + "]"
	} else if tr.Level == "warning" {
		label = s + " [" + skittles.Magenta(tr.Level) + "]"
	} else if tr.Level == "error" {
		label = s + " [" + skittles.Red(tr.Level) + "]"
	} else if tr.Level == "fatal" {
		label = s + " [" + skittles.BoldRed(tr.Level) + "]"
	} else if tr.Level == "panic" {
		label = s + " [" + skittles.BoldRed(tr.Level) + "]"
	}
	return label
}

// reverses and array
func reverse(array []*Terr) []*Terr {
	var new []*Terr
	for i := len(array) - 1; i >= 0; i-- {
		new = append(new, array[i])
	}
	return new
}
