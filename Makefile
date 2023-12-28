# Variable
APP_NAME = msrvct-exec
FULL_APP_NAM = Multiple Software Runtime Version Control Toolbox Executor

VERSION = $(shell git describe --tags --abbrev=0)
GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE = $(shell date +'%F_%T_%Z')
BUILD_TOOL = go_build
BUILD_ARGS = -s -w
LDFLAGS += -X github.com/devops-toolbox/msrvct/cmd._Name_=${APP_NAME}
LDFLAGS += -X github.com/devops-toolbox/msrvct/cmd._Version_=${VERSION}
LDFLAGS += -X github.com/devops-toolbox/msrvct/cmd._GitCommit_=${GIT_COMMIT}
LDFLAGS += -X github.com/devops-toolbox/msrvct/cmd._GitBranch_=${GIT_BRANCH}
LDFLAGS += -X github.com/devops-toolbox/msrvct/cmd._BuildDate_=${BUILD_DATE}
LDFLAGS += -X github.com/devops-toolbox/msrvct/cmd._BuildTool_=${BUILD_TOOL}

# PATH
BUILD_DIR="build"

.PHONY: default help debug build build_all build_darwin build_linux build_windows build_linux_amd64 build_linux_arm64 build_darwin_amd64 build_darwin_arm64 build_windows_amd64 build_windows_arm64


default: help
help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "  help  : Show help"
	@echo "  build : Build the binary of this project for all platform"

debug:
	echo "debug"
tidy:
	@go mod tidy
build: tidy
	@CGO_ENABLED=0 go build -ldflags '$(BUILD_ARGS) $(LDFLAGS)' -o ${BUILD_DIR}/${APP_NAME}
build_all: build_linux build_darwin build_windows
	@echo "build_all completed"
build_linux: clean tidy _build_linux_amd64 _build_linux_arm64
	@echo "build_linux completed"
build_darwin: clean tidy _build_darwin_amd64 _build_darwin_arm64
	@echo "build_darwin completed"
build_windows: clean tidy _build_windows_amd64 _build_windows_arm64
	@echo "build_windows completed"
build_linux_amd64: clean tidy _build_linux_amd64
	@echo "build_linux_amd64 completed"
build_linux_arm64: clean tidy _build_linux_arm64
	@echo "build_linux_arm64 completed"
build_darwin_amd64: clean tidy _build_darwin_amd64
	@echo "build_darwin_amd64 completed"
build_darwin_arm64: clean tidy _build_darwin_arm64
	@echo "build_darwin_arm64 completed"
build_windows_amd64: clean tidy _build_windows_amd64
	@echo "build_windows_amd64 completed"
build_windows_arm64: clean tidy _build_windows_arm64
	@echo "build_windows_arm64 completed"

_build_linux: _build_linux_amd64 _build_linux_arm64
_build_darwin: _build_darwin_amd64 _build_darwin_arm64
_build_windows: _build_windows_amd64 _build_windows_arm64
_build_linux_amd64:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/${APP_NAME}__${VERSION}-linux-amd64 -ldflags '$(BUILD_ARGS) $(LDFLAGS)'
_build_linux_arm64:
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${BUILD_DIR}/${APP_NAME}__${VERSION}-linux-arm64 -ldflags '$(BUILD_ARGS) $(LDFLAGS)'
_build_darwin_amd64:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/${APP_NAME}__${VERSION}-darwin-amd64 -ldflags '$(BUILD_ARGS) $(LDFLAGS)'
_build_darwin_arm64:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ${BUILD_DIR}/${APP_NAME}__${VERSION}-darwin-arm64 -ldflags '$(BUILD_ARGS) $(LDFLAGS)'
_build_windows_amd64:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/${APP_NAME}__${VERSION}-windows-amd64.exe -ldflags '$(BUILD_ARGS) $(LDFLAGS)'
_build_windows_arm64:
	@CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o ${BUILD_DIR}/${APP_NAME}__${VERSION}-windows-arm64.exe -ldflags '$(BUILD_ARGS) $(LDFLAGS)'
clean:
	@rm -rf ${BUILD_DIR}
