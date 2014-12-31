package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
			if file.Name() != "exceptionstack1" {
				continue
			}
			Pdbgf(file.Name())
			if in, err = os.Open("tests/" + file.Name()); err != nil {
				Pdbgf("Unable to access open file '%v'\n'%v'\n", file.Name(), err)
				t.Fail()
			}
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
