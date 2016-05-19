To reproduce:

```
go get github.com/henrikhodne/go-cover-test-panic
cd $GOPATH/src/github.com/henrikhodne/go-cover-test-panic
go test ./...
# panic is reported to be in foo.go:7
go test -covermode=count ./...
# panic is reported to be in foo.go:9
```

Output from my machine:

```
henrik@skyhawk:~/gopath/src/github.com/henrikhodne/go-cover-test-panic$ go test ./...
--- FAIL: TestFoo (0.00s)
panic: oops [recovered]
	panic: oops

goroutine 5 [running]:
panic(0x51b2a0, 0xc82000a5d0)
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/runtime/panic.go:464 +0x3e6
testing.tRunner.func1(0xc820098000)
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/testing/testing.go:467 +0x192
panic(0x51b2a0, 0xc82000a5d0)
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/runtime/panic.go:426 +0x4e9
github.com/henrikhodne/go-cover-test-panic.Foo(0xc80cdb49a9)
	/home/henrik/gopath/src/github.com/henrikhodne/go-cover-test-panic/foo.go:7 +0x72
github.com/henrikhodne/go-cover-test-panic.TestFoo(0xc820098000)
	/home/henrik/gopath/src/github.com/henrikhodne/go-cover-test-panic/foo_test.go:6 +0x18
testing.tRunner(0xc820098000, 0x646df0)
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/testing/testing.go:473 +0x98
created by testing.RunTests
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/testing/testing.go:582 +0x892
FAIL	github.com/henrikhodne/go-cover-test-panic	0.004s
henrik@skyhawk:~/gopath/src/github.com/henrikhodne/go-cover-test-panic$ go test -covermode=count ./...
--- FAIL: TestFoo (0.00s)
panic: oops [recovered]
	panic: oops

goroutine 5 [running]:
panic(0x51b380, 0xc82000a5d0)
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/runtime/panic.go:464 +0x3e6
testing.tRunner.func1(0xc820086000)
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/testing/testing.go:467 +0x192
panic(0x51b380, 0xc82000a5d0)
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/runtime/panic.go:426 +0x4e9
github.com/henrikhodne/go-cover-test-panic.Foo(0xc821269bb2)
	github.com/henrikhodne/go-cover-test-panic/_test/_obj_test/foo.go:9 +0x92
github.com/henrikhodne/go-cover-test-panic.TestFoo(0xc820086000)
	/home/henrik/gopath/src/github.com/henrikhodne/go-cover-test-panic/foo_test.go:6 +0x18
testing.tRunner(0xc820086000, 0x647e30)
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/testing/testing.go:473 +0x98
created by testing.RunTests
	/home/henrik/.gimme/versions/go1.6.linux.amd64/src/testing/testing.go:582 +0x892
FAIL	github.com/henrikhodne/go-cover-test-panic	0.004s
henrik@skyhawk:~/gopath/src/github.com/henrikhodne/go-cover-test-panic$ go env
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/henrik/gopath"
GORACE=""
GOROOT="/home/henrik/.gimme/versions/go1.6.linux.amd64"
GOTOOLDIR="/home/henrik/.gimme/versions/go1.6.linux.amd64/pkg/tool/linux_amd64"
GO15VENDOREXPERIMENT="1"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0"
CXX="g++"
CGO_ENABLED="1"
henrik@skyhawk:~/gopath/src/github.com/henrikhodne/go-cover-test-panic$ go version
go version go1.6 linux/amd64
```
