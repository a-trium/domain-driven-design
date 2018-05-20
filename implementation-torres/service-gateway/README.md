# Service - Gateway

## library
- [Errors](https://github.com/go-errors/errors) | ERROR
- [Gin](https://github.com/gin-gonic/gin) | REST API
- [GORM](http://doc.gorm.io/) | ORM
- [Zap](https://github.com/uber-go/zap) | logging
- [envconfig](https://github.com/kelseyhightower/envconfig) | Env configuration

```bash
$ dep ensure -add github.com/gin-gonic/gin  # Gin
$ dep ensure -add github.com/uber-go/zap    # Zap
$ dep ensure --add github.com/jinzhu/gorm   # GORM
$ dep ensure --add github.com/rubenv/sql-migrate/...  # sql-migrate
$ go get -v github.com/rubenv/sql-migrate/...         # sql-migrate for CLI

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