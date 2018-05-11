TAG 		= "Makefile"

NPM		= npm

.PHONY: install
install:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing dependencies"
	@ rm -rf node_modules
	@ $(NPM) install

.PHONY: run
run:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Running frontend application"
	@ $(NPM) run serve
