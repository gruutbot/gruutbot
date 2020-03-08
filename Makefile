NAME := gruutbot
MAIN_FILE := ./cli/cmd.go
BUILD_DIR := build
LINT_FLAGS := run --fast --enable=golint --enable=goconst --enable=gocyclo --enable=gocognit --enable=goimports --enable=maligned --enable=wsl --fix --color=always --print-issued-lines=false
PLUGINS_DIR := plugins

.PHONY: clean lint plugins

run: $(MAIN_FILE)
	@ go run $(MAIN_FILE)

clean:
	find $(BUILD_DIR)/ -type f -not -name '.gitkeep' -print0 | xargs -0 rm -f --
	find $(BUILD_DIR)/ -type d -not -name 'build' -print0 | xargs -0 rm -rf --

lint:
	golangci-lint $(LINT_FLAGS)

plugins:
	@ $(CURDIR)/$(PLUGINS_DIR)/build.sh $(PLUGINS_DIR) $(BUILD_DIR)/$(PLUGINS_DIR)

PLATFORMS := linux darwin
os = $(word 1, $@)

$(PLATFORMS):
	GOOS=$(os) GOARCH=amd64 go build -ldflags "$(VFLAG)" -o $(BUILD_VERSIONED_DIR)-$(os)/$(NAME)-$(os)-amd64 ./cli