# Ensure that the build tools are compiled. Go's caching makes this quick.
$(shell cd build/manifest && go build -o ../bin/manifest)

BUILD_TOOLS_DIR ?= build/bin
BUNDLE_DIR ?= build/dist
PLUGIN_ENTRYPOINT ?= $(shell build/bin/manifest entrypoint)

## Checks the code style, tests, builds and bundles the plugin.
all: prepare build bundle clean

## Runs any lints and unit tests defined for the server.
.PHONY: prepare
prepare:
	golangci-lint run
	go test -v -race ./server/...

## Propagates plugin manifest information into the server folder and builds the server.
.PHONY: build
build:
	mkdir $(BUNDLE_DIR)
	$(BUILD_TOOLS_DIR)/manifest apply
	cd server && env GOOS=linux GOARCH=amd64 go build -o ../$(BUNDLE_DIR)/$(PLUGIN_ENTRYPOINT)

## Generates a tar bundle of the plugin for install.
.PHONY: bundle
bundle:
	cp plugin.json $(BUNDLE_DIR)/
	cp -r assets $(BUNDLE_DIR)/
	tar -cvzf plugin.tar.gz -C $(BUNDLE_DIR) .

## Clean removes all build artifacts.
.PHONY: clean
clean:
	rm -rf $(BUILD_TOOLS_DIR)
	rm -rf $(BUNDLE_DIR)
