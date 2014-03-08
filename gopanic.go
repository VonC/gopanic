package main

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)
import "fmt"

// http://stackoverflow.com/questions/6359318/how-do-i-send-a-message-to-stderr-from-cmd
// a_command 2>&1 | gopanic
func main() {
	// http://stackoverflow.com/questions/12363030/read-from-initial-stdin-in-go
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(-1)
	}

	lines := strings.Split(string(b), "\n")
	lexer := &lexer{lines: lines}
	for state := lookForReason; state != nil; {
		state = state(lexer)
	}
	fmt.Println("done")
}

type stateFn func(*lexer) stateFn
type lexer struct {
	lines []string
	pos   int
}

var fileLineRx, _ = regexp.Compile(`\s*?\*?\s*?(/?[^\*\s/\\]+(?:[/\\][^/\\:]+)+):?(\d+)?`)
var causeRx, _ = regexp.Compile(`Line (\d+):[^:]+:\s+(.*?)$`)

func lookForReason(l *lexer) stateFn {
	line := l.lines[l.pos]
	//fmt.Printf("Look at line '%v': '%v'\n", l.pos, line)
	if strings.Contains(line, " *") {
		var fl *fileLine
		var err error
		if fl, err = newFileLine(line); err != nil {
			return l.errorf("Unable to read file for reason in line '%v'\n Cause: '%v'", l.pos, err)
		}
		l.pos = l.pos + 1
		line := l.lines[l.pos]
		res := causeRx.FindStringSubmatch(line)
		if res == nil {
			return l.errorf("Unable to read cause in line '%v': '%v'", l.pos, line)
		}
		var ln int
		if ln, err = strconv.Atoi(res[1]); err != nil {
			return l.errorf(fmt.Sprintf("Couldn't extract cause line number for from line '%v': '%v'", l.pos, line))
		}
		fl.line = ln
		r := &reason{cause: res[2], file: fl}
		fmt.Println("PANIC:\n" + r.String())
		l.pos = l.pos + 1
		return lookForStack
	}
	l.pos = l.pos + 1
	return lookForReason
}

type stack struct {
	function string
	fileLine *fileLine
}

var functionRx, _ = regexp.Compile(`\s*?([^ ]+/[^\.]+)\.([^\)]+\))`)

func (s *stack) String() string {
	msg := ""
	if s.fileLine != nil {
		msg = msg + s.fileLine.String() + " "
	}
	msg = msg + s.function
	return msg
}

func lookForStack(l *lexer) stateFn {
	line := l.lines[l.pos]
	if strings.Contains(line, "[running]:") ||
		strings.Contains(line, "runtime.panic") ||
		strings.Contains(line, "runtime/panic") {
		l.pos = l.pos + 1
		return lookForStack
	}
	if strings.Contains(line, "testing.tRunner(") ||
		strings.Contains(line, "created by testing.RunTests") {
		l.pos = l.pos + 2
		return lookForStack
	}
	if strings.TrimSpace(line) == "" {
		return nil
	}
	res := functionRx.FindStringSubmatch(line)
	if res == nil {
		return l.errorf("Unable to read function in stack line '%v': '%v'\n", l.pos, line)
	}
	function := res[1] + "." + res[2]

	l.pos = l.pos + 1
	line = l.lines[l.pos]

	var fl *fileLine
	var err error
	if fl, err = newFileLine(line); err != nil {
		return l.errorf("Unable to read file for reason in line '%v'\n Cause: '%v'", l.pos, err)
	}

	s := &stack{fileLine: fl, function: function}
	fmt.Println(s.String())

	l.pos = l.pos + 1
	return lookForStack
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	fmt.Printf(format, args...)
	return nil
}

type fileLine struct {
	file string
	line int
}

func newFileLine(line string) (*fileLine, error) {
	res := fileLineRx.FindStringSubmatch(line)
	if res == nil {
		return nil, errors.New(fmt.Sprintf("No file-line found in line '%v'", line))
	}
	var ln int
	var err error
	if res[2] != "" {
		if ln, err = strconv.Atoi(res[2]); err != nil {
			return nil, errors.New(fmt.Sprintf("Couldn't extract line number for from line '%v' '%v'", res[2], res))
		}
	}
	fl := &fileLine{file: res[1], line: ln}
	return fl, nil
}

func (fl *fileLine) String() string {
	res := fl.file
	if fl.line > 0 {
		res = res + ":" + strconv.Itoa(fl.line)
	}
	return res
}

type reason struct {
	file  *fileLine
	cause string
}

func (r *reason) String() string {
	return r.file.String() + " " + r.cause
}
