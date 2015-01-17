package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"regexp"
	"strings"
	"testing"

	. "github.com/VonC/godbg"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGoPanic(t *testing.T) {
	Convey("Test main", t, func() {
		fmt.Println("test main")
		files, err := ioutil.ReadDir("./tests")
		if err != nil {
			Pdbgf("Unable to access tests folder\n'%v'\n", err)
			t.Fail()
		}
		for _, file := range files {
			if file.Name() != "exceptionstack5" {
				continue
			}
			Pdbgf(file.Name())
			user, err := user.Current()
			if err != nil {
				Pdbgf("Unable to access current user\n'%v'\n", err)
				t.Fail()
			}
			username := user.Username
			// If domain\username => keep only username
			re := regexp.MustCompile(`^.*\\`)
			username = re.ReplaceAllLiteralString(username, "")
			var inb []byte
			if inb, err = ioutil.ReadFile("tests/" + file.Name()); err != nil {
				Pdbgf("Unable to access open file '%v'\n'%v'\n", file.Name(), err)
				t.Fail()
			}
			ins := string(inb)
			ins = strings.Replace(ins, "C:/Users/vonc/", "C:/Users/"+username+"/", -1)
			in = strings.NewReader(ins)
			// Pdbgf("In '%s'", ins)
			var resb []byte
			if resb, err = ioutil.ReadFile("tests/" + file.Name() + ".res"); err != nil {
				Pdbgf("Unable to access open RES file '%v'\n'%v'\n", file.Name()+".res", err)
				t.Fail()
			}
			res := string(resb)
			Pdbgf("ok open")
			writers = NewPdbg(SetBuffers)
			main()
			So(writers.OutString(), ShouldEqual, res)
		}
	})
}
