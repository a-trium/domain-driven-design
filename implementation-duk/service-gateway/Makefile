GOARCH  = amd64
GOOS    = $(OS)

ifeq ($(GOOS),)
  ifeq ($(shell  uname -s), Darwin)
    GOOS	= darwin
  else
    GOOS	= linux
  endif
endif

TAG 		= "Makefile"

GOCMD		= go
GODEP		= dep
GOVVV		= govvv
GOLINT		= gometalinter
GOBUILD		= GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOCMD) build
GOFMT		= $(GOCMD)fmt

MAIN		= main.go
BIN_DIR		= bin
CMD_DIR		= cmd
VENDOR_DIR	= vendor


GIT_COMMIT		= $(shell git rev-parse HEAD)
GIT_DIRTY		= $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)

APP				= server

.PHONY: install
install:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing dependencies"
	@ $(GODEP) ensure
	@ echo ""


.PHONY: build
build:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Building $(CMD_DIR)/$(APP)/$(MAIN)"
	@ $(eval LDFLAG := $(shell $(GOVVV) -flags -pkg $(shell go list ./internal/config)))
	@ $(GOBUILD) -race -i -ldflags="$(LDFLAG)" -o $(BIN_DIR)/$(APP) -v $(CMD_DIR)/$(APP)/$(MAIN)
	@ echo ""

.PHONY: run
run:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Running app: $(APP)"
	@$(BIN_DIR)/$(APP)

.PHONY: clean
clean:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Clean"
	@rm -rf ./bin
	@ echo ""
