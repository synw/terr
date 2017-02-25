package terr

import (
	"fmt"
	"errors"
	"os"
	"runtime"
	"strconv"
	"github.com/acmacalister/skittles"
)

type Terr struct {
	From string
	Error error
	Level string
}

func (e Terr) Format(args ...string) string {
	prefix := ""
	if len(args) == 1 {
		prefix = args[0]
	}
	var msg string
	sep := " "
	if e.Error != nil {
		msg = prefix+e.Error.Error()
	} else {
		sep = ""
	}
	msg = msg+sep+"("+e.From+")"
	return msg
}

type Trace struct {
	Errors []*Terr
}

func (trace Trace) Format(args ...string) string {
	prefix := ""
	suffix := "\n"
	num_args := len(args)
	if num_args == 1 {
		prefix = args[0]
	} else if num_args == 2 {
		prefix = args[0]
		suffix = args[1]
	}
	var msg string
	errs := reverse(trace.Errors)
	for i, er := range(errs) {
		s := strconv.Itoa(i)
		msg = msg+s+" "+er.Format(prefix)
		if (i+1) < len(errs) {
			msg = msg+suffix
		}
	}
	return msg
}

func (trace Trace) Formatl() string {
	var msg string
	errs := reverse(trace.Errors)
	for i, er := range(errs) {
		label := getLabelWithNum(er, i)
		msg = msg+label+" "+er.Format()
		if (i+1) < len(errs) {
			msg = msg+"\n"
		}
	}
	return msg
}

func (e Trace) Printp(prefix string) {
	fmt.Println(e.Format(prefix, ""))
}

func (e Trace) Prints(suffix string) {
	fmt.Println(e.Format("", suffix))
}

func (e Trace) Printps(suffix string, prefix string) {
	fmt.Println(e.Format(prefix, suffix))
}

func (e Trace) Printl(from string) {
	fmt.Println("-------------- errors ("+from+") --------------")
	fmt.Println(e.Formatl())
}

func (e Trace) Printf(from string) {
	fmt.Println("-------------- errors ("+from+") --------------")
	fmt.Println(e.Format())
	//fmt.Println("---------------------------------------------")
}

func (e Trace) Print(from string) {
	fmt.Println("-------------- errors ("+from+") --------------")
	fmt.Println(e.Format())
}

func (trace Trace) Errs() []error {
	var errs []error
	if len(trace.Errors) > 0 {
		for _, er := range(trace.Errors) {
			if er != nil {
				errs = append(errs, er.Error)
			}
		}
	}
	return errs
}

func (trace Trace) Err() error {
	var err_str string
	if len(trace.Errors) > 0 {
		for _, er := range(trace.Errors) {
			if er != nil {
				err_str = err_str+er.Error.Error()
			}
		}
	}
	e := errors.New(err_str)
	return e
}

func New(from string, err error) *Trace {
	er := &Terr{from, err, ""}
	var prev *Trace
	t := newFromErr(er, from, err, prev)
	return t
}

func Add(from string, err error, previous_traces ...*Trace) *Trace {
	er := &Terr{from, err, ""}
	t := newFromErr(er, from, err, previous_traces...)
	return t
}

func Pass(from string, previous_traces ...*Trace) *Trace {
	var err error
	er := &Terr{from, err, ""}
	t := newFromErr(er, from, err, previous_traces...)
	return t
}

func Push(from string, err error, previous_traces ...*Trace) *Trace {
	er := &Terr{from, err, ""}
	t := newFromErr(er, from, err, previous_traces...)
	fmt.Println(er.Format())
	return t
}

func Stack(from string, err error, previous_traces ...*Trace) *Trace {
	er := &Terr{from, err, ""}
	t := newFromErr(er, from, err, previous_traces...)
	fmt.Println(getLabel(er), er.Format())
	var stack [4096]byte
	runtime.Stack(stack[:], false)
	fmt.Println(er.Format())
	fmt.Printf("%s\n", stack[:])
	return t
}

func Fatal(from string, trace *Trace) {
	msg := "From "+skittles.BoldWhite(from)
	fmt.Println(msg)
	fmt.Println(trace.Format())
	os.Exit(1)
}

func Critical(from string, err error, previous_traces ...*Trace) *Trace {
	er := &Terr{from, err, "critical"}
	t := newFromErr(er, from, err, previous_traces...)
	return t
}

func Important(from string, err error, previous_traces ...*Trace) *Trace {
	er := &Terr{from, err, "important"}
	t := newFromErr(er, from, err, previous_traces...)
	return t
}

func Minor(from string, err error, previous_traces ...*Trace) *Trace {
	er := &Terr{from, err, "minor"}
	t := newFromErr(er, from, err, previous_traces...)
	return t
}

func Debug(from string, err error, previous_traces ...*Trace) *Trace {
	er := &Terr{from, err, "debug"}
	t := newFromErr(er, from, err, previous_traces...)
	return t
}

func newFromErr(er *Terr, from string, err error, previous_traces ...*Trace) *Trace {
	var new_errors []*Terr
	new_errors = append(new_errors, er)
	if len(previous_traces) > 0 {
		for _, trace := range(previous_traces) {
			if trace != nil {
				if len(trace.Errors) > 0 {
					for _, err := range(trace.Errors) {
						new_errors = append(new_errors, err)
					}
				}
			}
		}
	}
	new_trace := &Trace{new_errors}
	return new_trace
}

func reverse(array []*Terr) []*Terr {
	var new []*Terr
	for i := len(array) - 1; i >= 0; i-- {
		new = append(new, array[i])
	}
	return new
}


func getLabel(er *Terr) string {
	label := "["+skittles.Red("error")+"]"
	if er.Level == "critical" {
		label = "["+skittles.BoldRed("critical")+"]"
	} else if er.Level == "minor" {
		label = "[minor error]"
	} else if er.Level == "debug" {
		label = "["+skittles.Yellow("debug")+"]"
	} else if er.Level == "important" {
		label = "["+skittles.BoldGreen("important")+"]"
	}
	return label
}

func getLabelWithNum(er *Terr, i int) string {
	s := strconv.Itoa(i)
	label := "["+skittles.Red("error")+" "+s+"]"
	if er.Level == "critical" {
		label = "["+skittles.BoldRed("critical")+" "+s+"]"
	} else if er.Level == "minor" {
		label = "[minor error]"
	} else if er.Level == "debug" {
		label = "["+skittles.Yellow("debug")+" "+s+"]"
	} else if er.Level == "important" {
		label = "["+skittles.BoldGreen("important")+" "+s+"]"
	}
	return label
}
