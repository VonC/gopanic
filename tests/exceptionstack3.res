PANIC:
../asciidocgo/substitutors_test.go:800 invalid memory address or nil pointer dereference
../../../../../../../Temp/2/makerelease250988475/go/src/pkg/regexp/regexp.go:610  .  regexp.(*Regexp).allMatches(0x0, 0x63c570, 0x4, 0x0, 0x0, ...)
../../../../../../../Temp/2/makerelease250988475/go/src/pkg/regexp/regexp.go:1052 .  regexp.(*Regexp).FindAllStringSubmatchIndex(0x0, 0x63c570, 0x4, 0x5, 0x4cd038, ...)
../regexps/regexps.go:52                                                          .  github.com/VonC/asciidocgo/consts/regexps.NewReres(0x63c570, 0x4, 0x0, 0x2)
/substitutors.go:621                                                              .  github.com/VonC/asciidocgo.(*substitutors).restorePassthroughs(0xc0840f52a0, 0x63c570, 0x4, 0x3, 0xc08406a750)
../asciidocgo/substitutors_test.go:800                                            .  github.com/VonC/asciidocgo.funcÂ·311()
../asciidocgo/substitutors_test.go:806                                            .  github.com/VonC/asciidocgo.TestSubstitutor(0xc084038c60)
