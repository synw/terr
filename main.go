package terr

import (
	"fmt"
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
	level := "["+skittles.Red("error")+"]"
	if e.Level == "critical" {
		level = "["+skittles.BoldRed("critical")+"]"
	} else if e.Level == "minor" {
		level = "[minor error]"
	} else if e.Level == "debug" {
		level = "["+skittles.Yellow("debug")+"]"
	}
	msg := prefix+e.Error.Error()
	msg = msg+" ("+e.From+")"
	msg = level+" "+msg+" "
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

func (e Trace) Print(sep ...string) {
	fmt.Println(e.Format(sep...))
}

func New(from string, err error, previous_traces ...*Trace) *Trace {
	terr := &Terr{from, err, ""}
	t := newFromErr(terr, from, err, previous_traces...)
	return t
}

func Critical(from string, err error, previous_traces ...*Trace) *Trace {
	terr := &Terr{from, err, "critical"}
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
		for _, e := range(previous_traces[0].Errors) {
			new_errors = append(new_errors, e)
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
