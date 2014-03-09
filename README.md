gopanic
=======

Go panic reader, for quickly see where the error comes from

Quick, tell me what are the classes of *your* program involved in the following error:

````
  * C:/Users/vonc/prog/go/src/github.com/VonC/asciidocgo/substitutors_test.go
  Line 638: - runtime error: slice bounds out of range
  goroutine 10 [running]:
  runtime.panic(0x5dc2c0, 0x8b4b6a)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/runtime/panic.c:248 +0x11b
  github.com/VonC/asciidocgo/consts/regexps.newReresLA(0x66c090, 0x5d, 0xc08400caa0, 0x0, 0x1)
        C:/Users/vonc/prog/go/src/github.com/VonC/asciidocgo/consts/regexps/regexps.go:97 +0x90c
  github.com/VonC/asciidocgo/consts/regexps.NewReresLAGroup(0x66c090, 0x5d, 0xc08400caa0, 0x5)
        C:/Users/vonc/prog/go/src/github.com/VonC/asciidocgo/consts/regexps/regexps.go:56 +0x47
  github.com/VonC/asciidocgo/consts/regexps.NewIndextermInlineMacroRxres(0x66c090, 0x5d, 0x617920)
        C:/Users/vonc/prog/go/src/github.com/VonC/asciidocgo/consts/regexps/regexps.go:294 +0x41
  github.com/VonC/asciidocgo.(*substitutors).SubMacros(0xc0840ce780, 0x66c090, 0x5d, 0x3, 0xc084067780)
        github.com/VonC/asciidocgo/_test/substitutors.go:1298 +0x4016
  github.com/VonC/asciidocgo.func·290()
        C:/Users/vonc/prog/go/src/github.com/VonC/asciidocgo/substitutors_test.go:638 +0x3d
  github.com/VonC/asciidocgo.TestSubstitutor(0xc084036cf0)
        C:/Users/vonc/prog/go/src/github.com/VonC/asciidocgo/substitutors_test.go:640 +0x130e
  testing.tRunner(0xc084036cf0, 0x8b3208)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/testing/testing.go:391 +0x8e
  created by testing.RunTests
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/testing/testing.go:471 +0x8b5

  goroutine 1 [chan receive]:
  testing.RunTests(0x6851f8, 0x8b3160, 0x8, 0x8, 0x301)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/testing/testing.go:472 +0x8d8
  testing.Main(0x6851f8, 0x8b3160, 0x8, 0x8, 0x8b6fe0, ...)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/testing/testing.go:403 +0x87
  main.main()
        github.com/VonC/asciidocgo/_test/_testmain.go:119 +0x11e
````

... If all you see is a blurring wall of text, you are not the only one.

This `gopanic` utility wants to help filter the output, and display only what matters.

Install `gopanic`:

    go get github.com/VonC/gopanic

Then go to any sub-folder of your project:
````
cd path/to/your/go/project
cd a/subdirectory
go test ../..|gopanic
````

The panic dump stack will now look like:

````
C:\Users\vonc\prog\go\src\github.com\VonC\asciidocgo\consts\compliance>go test ../..|gopanic
PANIC:
../../substitutors_test.go:638 slice bounds out of range
../regexps/regexps.go:97       regexps.newReresLA(0x667550, 0x5d, 0xc08400cbe0, 0x0, 0x1)
../regexps/regexps.go:56       regexps.NewReresLAGroup(0x667550, 0x5d, 0xc08400cbe0, 0x5)
../regexps/regexps.go:294      regexps.NewIndextermInlineMacroRxres(0x667550, 0x5d, 0x613960)
../../substitutors.go:1084     (*substitutors)
../../substitutors_test.go:638 func·290()
../../substitutors_test.go:640 TestSubstitutor(0xc084035cf0)
````

This seems easier on the eyes, and easier to debug.
