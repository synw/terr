package terr

import (
	"errors"
	"fmt"
	"github.com/acmacalister/skittles"
	//"os"
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

func New(from string, errMsg string, level ...string) *Trace {
	lvl := ""
	if len(level) > 0 {
		lvl = level[0]
	}
	err := errors.New(errMsg)
	_, file, line, _ := runtime.Caller(1)
	ter := &Terr{from, lvl, err, file, line}
	ters := []*Terr{ter}
	tr := Trace{ters}
	return &tr
}

// ------------------- Methods -------------------

func (trace Trace) Add(from string, errMsg string, level ...string) *Trace {
	lvl := ""
	if len(level) > 0 {
		lvl = level[0]
	}
	err := errors.New(errMsg)
	_, file, line, _ := runtime.Caller(1)
	ter := &Terr{from, lvl, err, file, line}
	trace.Errors = append(trace.Errors, ter)
	return &trace
}

func (trace Trace) Check() {
	if len(trace.Errors) > 0 {
		trace.Print()
	}
}

func (trace Trace) Print() {
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
	fmt.Println(msg)
}

func (trace Trace) ToErr() error {
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

/*
func (trace Trace) Debug(msg string) {
	_, file, line, _ := runtime.Caller(1)
	//fmt.Printf("[cgl] debug %s:%d %v\n", file, line, info)
	fmt.Println(line)
	fmt.Println(file)
	fmt.Println(msg)
}

func getInfo() (string, int) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(line)
	fmt.Println(file)
	return file, line
}*/

// ------------------- Functions -------------------

// errors stack traces management

/*
func Pass(from string, previous_traces ...*Trace) *Trace {
	var err error
	er := &Terr{from, err, ""}
	t := newFromErr(er, from, err, "", previous_traces...)
	return t
}

func Stack(from string, errMsg string, previous_traces ...*Trace) *Trace {
	err := errors.New(errMsg)
	ter := &Terr{from, err, ""}
	tr := &Trace{[]*Terr{ter}}
	t := newTr(ter, from, errMsg, "", previous_traces...)
	fmt.Println(getLabel(ter), tr.Msg())
	var stack [4096]byte
	runtime.Stack(stack[:], false)
	tr.Print()
	fmt.Printf("%s\n", stack[:])
	return t
}*/

/*
func (trace Trace) Error() string {
	trace.Print()
}

func (trace Trace) Fatal(from ...string) {
	fr := ""
	if len(from) > 0 {
		fr = from[0]
	}
	msg := skittles.BoldRed("Fatal error") + " from " + skittles.BoldWhite(fr)
	fmt.Println(msg)
	trace.Print()
	os.Exit(1)
}*/

/*
func Push(from string, err error, previous_traces ...*Trace) *Trace {
	tr := &Terr{from, err, ""}
	ntr := newFromErr(tr, from, err, "", previous_traces...)
	fmt.Println(tr.Print())
	return ntr
}

func Fatal(from string, tr *Trace) {
	msg := skittles.BoldRed("Fatal error") + " from " + skittles.BoldWhite(from)
	fmt.Println(msg)
	tr.Print()
	os.Exit(1)
}

func Ok(msg string) {
	msg = "[" + skittles.Green("ok") + "] " + msg
	fmt.Println(msg)
}

func Debug(args ...interface{}) {
	num_args := len(args)
	if num_args < 1 {
		return
	}
	t := fmt.Sprintf("%T", args[0])
	objs := args
	if t == "string" {
		msg := "[" + skittles.Yellow("debug") + "] " + args[0].(string)
		fmt.Println(msg)
	} else {
		objs = args[1:]
	}
	for _, o := range objs {
		fmt.Println(fmt.Sprintf("%T %#v", o, o))
	}
}

func Err(msg string) error {
	msg = "[" + skittles.Red("error") + "] " + msg
	err := errors.New(msg)
	return err
}

// internal methods
/*
func newTr(ter *Terr, from string, errMsg string, level string, ptr ...*Trace) *Trace {
	var ters []*Terr
	ters = append(ptr, ter)
	if len(ptr) > 0 {
		for _, tr := range ptr {
			if tr != nil {
				if len(tr.Errors) > 0 {
					for _, otr := range tr.Errors {
						trs = append(trs, otr)
					}
				}
			}
		}
	}
	new_trace := &Trace{trs}
	return new_trace
}*/

func getLabelWithNum(er *Terr, i int) string {
	s := strconv.Itoa(i)
	label := "[" + skittles.Red("error") + " " + s + "]"
	if er.Level == "critical" {
		label = "[" + skittles.BoldRed("critical") + " " + s + "]"
	} else if er.Level == "minor" {
		label = "[minor error]"
	} else if er.Level == "debug" {
		label = "[" + skittles.Yellow("debug") + " " + s + "]"
	} else if er.Level == "important" {
		label = "[" + skittles.BoldGreen("important") + " " + s + "]"
	}
	return label
}

func reverse(array []*Terr) []*Terr {
	var new []*Terr
	for i := len(array) - 1; i >= 0; i-- {
		new = append(new, array[i])
	}
	return new
}
