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
			if file.Name() != "exceptionstack2" {
				continue
			}
			Pdbgf(file.Name())
			if in, err = os.Open("tests/" + file.Name()); err == nil {
				Pdbgf("ok open")
				main()
			} else {
				Pdbgf("Unable to access open file '%v'\n'%v'\n", file.Name(), err)
				t.Fail()
			}
		}
	})
}
