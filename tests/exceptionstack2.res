PANIC:
../regexps/regexps_test.go:679 invalid memory address or nil pointer dereference
../../../../../../../Temp/2/makerelease250988475/go/src/pkg/regexp/regexp.go:610  .  regexp.(*Regexp).allMatches(0x0, 0x607530, 0xd0, 0x0, 0x0, ...)
../../../../../../../Temp/2/makerelease250988475/go/src/pkg/regexp/regexp.go:1052 .  regexp.(*Regexp).FindAllStringSubmatchIndex(0x0, 0x607530, 0xd0, 0xd1, 0xc0840009c0, ...)
/regexps.go:48                                                                    .  github.com/VonC/asciidocgo/consts/regexps.NewReres(0x607530, 0xd0, 0x0, 0xc08405f780)
/regexps.go:246                                                                   .  github.com/VonC/asciidocgo/consts/regexps.NewInlineAnchorRxres(0x607530, 0xd0, 0x3)
../regexps/regexps_test.go:679                                                    .  github.com/VonC/asciidocgo/consts/regexps.funcÂ·050()
../regexps/regexps_test.go:686                                                    .  github.com/VonC/asciidocgo/consts/regexps.TestRegexps(0xc08403a1b0)
