
TAG = "Makefile"

.PHONY: prepare
prepare:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing prerequisites"

	@ echo "\t golang:dep"
	@ go get -u github.com/golang/dep/cmd/dep

	@ echo ""