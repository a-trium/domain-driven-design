TAG = "Makefile"
VERSION = $(shell cat ./VERSION)
SWAGGER_SPEC = "../schema-swagger/gateway.yaml"

.PHONY: prepare
prepare:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing prerequisites"

	@ echo "\t golang:dep"
	@ go get -u github.com/golang/dep/cmd/dep

	@ echo "\t golang:go-swagger"
	@ go get -u github.com/go-swagger/go-swagger/cmd/swagger

	@ echo ""

.PHONY: swagger-version
swagger-version:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Updating swagger spec version: $(VERSION)"
	@VERSION=$(VERSION) SWAGGER_FILE=$(SWAGGER_SPEC) ../script/update-swagger-version.sh


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

