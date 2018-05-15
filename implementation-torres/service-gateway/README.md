# Service - Gateway

## library
- [Errors](https://github.com/go-errors/errors)
- [Gin](https://github.com/gin-gonic/gin) 
- [GORM](http://doc.gorm.io/)
- [Zap](https://github.com/uber-go/zap)

```bash
$ dep ensure -add github.com/gin-gonic/gin

```

## directory
- /bin
- /cmd/server/main.go
- /internal
- Makefile

## build & run

- test command
```{bash}
$ GOOS=darwin GOARCH=amd64 go build -o bin/hello -v cmd/hello/main.go
$ ./bin/hello
``` 
- install
```bash
$ make install
```
- check : format, style, lint
```bash
$ make check
```
- build
```bash
$ make build
```
- run
```bash
$ make run
```
- start
```bash
$ make clean install build run
```

## Makefile
- usage go command : go build, go install, dep, gometalinter, ginkgo, rm, gofmt
- make command : all, lint, install, build, build-all, run, run-cont, clean, test, test-cont
- go build options : -ldflags, -i, -o, -v
    - -i : flag installs the packages that are dependencies of the target.
    - -v : print the names of packages as they are compiled.
    - -ldflags : arguments to pass on each go tool link invocation.