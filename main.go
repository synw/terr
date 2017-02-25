package terr

import (
	"fmt"
	"errors"
	"os"
	"runtime"
	"github.com/acmacalister/skittles"
)

type Terr struct {
	From string
	Error error
	Level string
}

func (e Terr) Formatl(args ...string) string {
	prefix := ""
	if len(args) == 1 {
		prefix = args[0]
	}
	level := "["+skittles.Red("error")+"]"
	if e.Level == "critical" {
		level = "["+skittles.BoldRed("critical")+"]"
	} else if e.Level == "minor" {
		level = "[minor error]"
	} else if e.Level == "debug" {
		level = "["+skittles.Yellow("debug")+"]"
	} else if e.Level == "important" {
		level = "["+skittles.BoldGreen("important")+"]"
	}
	var msg string
	if e.Error != nil {
		msg = prefix+e.Error.Error()
	}
	msg = msg+" ("+e.From+")"
	msg = level+" "+msg+" "
	return msg
}

func (e Terr) Format(args ...string) string {
	prefix := ""
	if len(args) == 1 {
		prefix = args[0]
	}
	var msg string
	if e.Error != nil {
		msg = prefix+e.Error.Error()
	}
	msg = msg+" ("+e.From+")"
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
	for i, terr := range(errs) {
		msg = msg+terr.Format(prefix)
		if (i+1) < len(errs) {
			msg = msg+suffix
		}
	}
	return msg
}

func (trace Trace) Formatl(args ...string) string {
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
	for i, terr := range(errs) {
		msg = msg+terr.Formatl(prefix)
		if (i+1) < len(errs) {
			msg = msg+suffix
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

func (e Trace) Printl(sep ...string) {
	fmt.Println(e.Formatl(sep...))
}

func (e Trace) Print(sep ...string) {
	fmt.Println(e.Format(sep...))
}

func (e Trace) Printf(from string) {
	fmt.Println("-------------- errors ("+from+") --------------")
	fmt.Println(e.Format())
	//fmt.Println("---------------------------------------------")
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
	terr := &Terr{from, err, ""}
	var prev *Trace
	t := newFromErr(terr, from, err, prev)
	return t
}

func Add(from string, err error, previous_traces ...*Trace) *Trace {
	terr := &Terr{from, err, ""}
	t := newFromErr(terr, from, err, previous_traces...)
	return t
}

func Pass(from string, previous_traces ...*Trace) *Trace {
	var err error
	terr := &Terr{from, err, ""}
	t := newFromErr(terr, from, err, previous_traces...)
	return t
}

func Push(from string, err error, previous_traces ...*Trace) *Trace {
	terr := &Terr{from, err, ""}
	fmt.Println("PUSH")
	t := newFromErr(terr, from, err, previous_traces...)
	fmt.Println(terr.Formatl())
	return t
}

func Stack(from string, err error, previous_traces ...*Trace) *Trace {
	terr := &Terr{from, err, ""}
	t := newFromErr(terr, from, err, previous_traces...)
	fmt.Println(terr.Format())
	var stack [4096]byte
	runtime.Stack(stack[:], false)
	fmt.Println(terr.Format())
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
	terr := &Terr{from, err, "critical"}
	t := newFromErr(terr, from, err, previous_traces...)
	return t
}

func Important(from string, err error, previous_traces ...*Trace) *Trace {
	terr := &Terr{from, err, "important"}
	t := newFromErr(terr, from, err, previous_traces...)
	return t
}

func Minor(from string, err error, previous_traces ...*Trace) *Trace {
	terr := &Terr{from, err, "minor"}
	t := newFromErr(terr, from, err, previous_traces...)
	return t
}

func Debug(from string, err error, previous_traces ...*Trace) *Trace {
	terr := &Terr{from, err, "debug"}
	t := newFromErr(terr, from, err, previous_traces...)
	return t
}

func newFromErr(terr *Terr, from string, err error, previous_traces ...*Trace) *Trace {
	var new_errors []*Terr
	new_errors = append(new_errors, terr)
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
