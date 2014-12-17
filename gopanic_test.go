package main

import (
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGoPanic(t *testing.T) {
	Convey("Test main", t, func() {
		fmt.Println("test main")
	})
}
