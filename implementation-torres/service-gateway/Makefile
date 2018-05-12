# application info
APP			= server
BIN_DIR		= bin
CMD_DIR		= cmd
VENDOR_DIR	= vendor
BIN			= ${BIN_DIR}/${APP}
MAIN		= ${CMD_DIR}/${APP}/main.go
TAG			= make

# build options
VERSION		= 0.0.1
BUILDTIME	= $(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')
IMPORTPATH  = github.com/a-trium/domain-driven-design/implementation-torres/service-gateway
GITSHA		= $(shell git rev-parse HEAD)
LDFLAGS 	= $(shell govvv -flags -pkg $(shell go list ./internal/pkg/config))

GOARCH  = amd64
GOOS    = $(OS)
ifeq ($(GOOS),)
  ifeq ($(shell  uname -s), Darwin)
    GOOS	= darwin
  else
    GOOS	= linux
  endif
endif

# go command
GOCMD		= go
GOGET		= $(GOCMD) get -u -v
GOBUILD		= GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOCMD) build
GOINSTALL	= $(GOCMD) install
GODEP		= dep
GOX			= gox
GOFMT 		= gofmt
GOLINT 		= gometalinter
GOTEST		= ginkgo
GOTEST_OPT	= -r -p -v -cover
GO_FILES	= $(shell $(GOCMD) list ./... | grep -v /vendor/)
RM			= rm -rf

# make commands
all: install lint build-all test

$(APP):
	@$(eval TARGET := $@)
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) - building app: $(TARGET)"

prepare:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) - installing prerequisites"
	# [install] go tools
	$(GOGET) github.com/onsi/ginkgo/ginkgo			# for test
	$(GOGET) github.com/mitchellh/gox	 			# for cross compile
	$(GOGET) github.com/githubnemo/CompileDaemon 	# for reload
	$(GOGET) github.com/alecthomas/gometalinter 	# for lint
	$(GOGET) github.com/ahmetb/govvv				# for build wrapper

install:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) - installing / updating dependencies"
	# [install] go lint install & update
	$(GOLINT) --install --update --force

	# [install] dep, Gopkg.toml
	@$(GODEP) ensure -update
	$(GODEP) ensure

build: clean $(APP)
	$(GOBUILD) \
		-ldflags="$(LDFLAGS)" \
		-i \
		-o ${BIN_DIR}/${APP} \
		-v ${MAIN}

build-all:
	$(GOX) -verbose \
		-ldflags=$(LDFLAGS) \
		-os="linux darwin windows freebsd openbsd netbsd" \
		-arch="amd64 386 armv5 armv6 armv7 arm64" \
		-osarch="!darwin/arm64" \
		-output="${BIN_DIR}/${APP}" \
		${MAIN}

run:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) - running app: $(APP)"
	${BIN_DIR}/${APP}

run-cont: build
	CompileDaemon \
		-exclude-dir="${VENDOR_DIR}" \
		-build="make build" \
		-command="make run"\
		-graceful-kill=true

clean:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) - cleaning"
	find . -type f -name '*.coverprofile' -exec rm {} +
	$(RM) ${BIN_DIR}/
	$(RM) vendor/$(IMPORTPATH)

test:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) - testing"
	$(GOTEST) $(GOTEST_OPT)

test-cont:
	@echo "\n[MAKEFILE] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) test-cont"
	$(GOTEST) watch $(GOTEST_OPT)

lint:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) - linting"
	@$(GOLINT) --vendor --errors ./... \
		--skip=internal/mock --skip=pkg \
		--enable=unparam --enable=nakedret --enable=staticcheck
style:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) - checking code style"
	@! $(GOFMT) -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

format:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')) - formatting code"
	@$(GOCMD) fmt $(GO_FILES)

check: format style lint
