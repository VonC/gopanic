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
			Pdbgf("ok open")
			writers = NewPdbg(SetBuffers)
			main()
			So(writers.OutString(), ShouldEqual, `PANIC:
gopanic_test.go:25 index out of range
gopanic.go:58      lookForReason(0xc082005080, 0x600208)
gopanic.go:37      main()
gopanic_test.go:25 func·001()
gopanic_test.go:31 TestGoPanic(0xc082044000)
`)

		}
	})
}
