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
GOGET		= go get -u

MAIN		= main.go
BIN_DIR		= bin
CMD_DIR		= cmd
VENDOR_DIR	= vendor


GIT_COMMIT		= $(shell git rev-parse HEAD)
GIT_DIRTY		= $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
VERSION 		= $(shell cat ./VERSION)

APP				= server


SWAGGER_SPEC 	= "../schema-swagger/gateway-rest.yaml"

.PHONY: prepare
prepare:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing prerequisites"

	@ echo "\t golang:dep \t\t(dependency)"
	@ $(GOGET) github.com/golang/dep/cmd/dep

	@ echo "\t golang:go-swagger \t(build)"
	@ $(GOGET) github.com/go-swagger/go-swagger/cmd/swagger

	@ echo "\t golang:govvv \t\t(build)"
	@ $(GOGET) github.com/ahmetb/govvv

	@ echo "\t golang:ginkgo \t\t(testing)"
	@ $(GOGET) github.com/ahmetb/govvv

	@ echo "\t golang:gomega \t\t(testing)"
	@ $(GOGET) github.com/ahmetb/govvv

	@ echo ""

.PHONY: install
install:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing dependencies"
	@ $(GODEP) ensure
	@ echo ""

.PHONY: swagger-version
swagger-version:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Updating swagger spec version: $(VERSION)"
	@VERSION=$(VERSION) SWAGGER_FILE=$(SWAGGER_SPEC) ../script/update-swagger-version.sh
	@ echo ""

.PHONY: swagger-server
swagger-server: swagger-version
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Cleaning swagger files"
	@ rm -rf pkg/generated/swagger/* || true
	@ mkdir -p pkg/generated/swagger || true
	@ echo ""

	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Validating swagger spec"
	@ swagger validate $(SWAGGER_SPEC)
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Generating swagger code"
	@ swagger generate server --spec=$(SWAGGER_SPEC) --exclude-main \
		 --target=pkg/generated/swagger --model-package=swagmodel \
		 --server-package=swagserver --api-package=swagapi
	@ echo ""


.PHONY: swagger
swagger: swagger-version swagger-server

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


.PHONY: release
release: prepare install clean swagger build
