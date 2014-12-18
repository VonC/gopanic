package main

import (
	"fmt"
	"io/ioutil"
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
			fmt.Println(file.Name())
		}
	})
}
