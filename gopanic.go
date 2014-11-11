package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)
import "fmt"

var gopath = os.Getenv("gopath") + "/src"
var pwd, _ = os.Getwd()

// http://stackoverflow.com/questions/6359318/how-do-i-send-a-message-to-stderr-from-cmd
// a_command 2>&1 | gopanic
func main() {
	gopath = strings.Replace(gopath, "\\", "/", -1)
	pwd = strings.Replace(pwd, "\\", "/", -1)
	// http://stackoverflow.com/questions/12363030/read-from-initial-stdin-in-go
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(-1)
	}

	lines := strings.Split(string(b), "\n")
	lexer := &lexer{lines: lines, stacks: []*stack{}}
	for state := lookForReason; state != nil; {
		state = state(lexer)
	}
	for _, stack := range lexer.stacks {
		stack.max = lexer.max + 2
		fmt.Println(stack)
	}
	fmt.Println("done")
}

type stateFn func(*lexer) stateFn
type lexer struct {
	lines  []string
	pos    int
	stacks []*stack
	max    int
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
	max      int
}

var functionRx, _ = regexp.Compile(`\s*?([^ ]+/[^\.]+)\.((?:(?:[^\)]+\))\.?)+)`)

func (s *stack) String() string {
	msg := ""
	f := s.function
	if s.fileLine != nil {
		fl := s.fileLine.String()
		l := s.max - len(fl)
		msg = msg + fl + strings.Repeat(" ", l)
		if strings.HasPrefix(f, s.fileLine.prefix) {
			f = f[len(s.fileLine.prefix)+1:]
		}
	}
	msg = msg + f
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
	//fmt.Println(res)
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
	l.stacks = append(l.stacks, s)
	if l.max < fl.lenf {
		l.max = fl.lenf
	}

	l.pos = l.pos + 1
	return lookForStack
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	fmt.Printf(format, args...)
	return nil
}

type fileLine struct {
	file   string
	prefix string
	line   int
	lenf   int
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
	file := strings.TrimSpace(res[1])
	filedir := filepath.Dir(file)
	f := filedir
	rel, _ := filepath.Rel(pwd, filedir)
	//fmt.Println("aaa: " + rel)
	if strings.HasPrefix(file, gopath) {
		file = file[len(gopath)+1:]
	}
	if strings.HasPrefix(pwd, gopath) {
		rel = strings.Replace(rel, "\\", "/", -1)
		rels := strings.Split(rel, "/")
		m := ""
		b := false
		for _, arel := range rels {
			if arel == ".." {
				filedir = filepath.Dir(filedir)
				m = m + "../"
			} else if arel != "" {
				b = true
			}
		}
		if !b && m != "" {
			filedir = f
		}
		if !strings.Contains(rel, "..") && rel != "." {
			filedir = filedir[:len(filedir)-len(rel)-1]
		}
		filedir = strings.Replace(filedir, "\\", "/", -1)
		if strings.HasPrefix(filedir, gopath) {
			filedir = filedir[len(gopath)+1:]
		}
		//fmt.Printf("filedir='%v' => '%v'\n", f, filedir)
		if strings.HasPrefix(file, filedir) {
			file = file[len(filedir)+1:]
		}
		file = m + file
	}
	fl := &fileLine{file: file, line: ln, prefix: filedir, lenf: len(file) + len(res[2])}
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
