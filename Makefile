.PHONY:  generate dependencies upload cleaner cleaner_cronjob migrate_up sqlc lint fmt test-upload test-cleaner test-cleaner-cronjob .remove_empty_dirs .pre-check-go

SRCS = $(patsubst ./%,%,$(shell find . -name "*.go" -not -path "*vendor*" -not -path "*.pb.go"))
MOCK_PACKAGES = \
	internal/app/upload/server/handlers \
	internal/app/upload/core \
	internal/app/imagecleaner/core \
	internal/app/consumer \
	internal/app/synchronizer \
    internal/pkg/database \
    internal/pkg/imghandler

MOCKED_FILES = $(shell find . -name DOES_NOT_EXIST_FILE $(patsubst %,-or -path "./%/mocks/*.go",$(MOCK_PACKAGES)))
MOCKED_FOLDERS = $(patsubst %,%/mocks,$(MOCK_PACKAGES))

.remove_empty_dirs:
	-@find . -type d -print | xargs rmdir 2>/dev/null | true

generate: $(MOCKED_FILES) $(MOCKED_FOLDERS) | .remove_empty_dirs ## to generate all auto-generated files

dependencies:
	$(GO) mod download

migrate_up:
	migrate -path ./db/migration -database "${DB__SERVER}://${DB__USERNAME}:${DB__PASSWORD}@${DB__HOST}:${DB__PORT}/${DB__NAME}?sslmode=disable" -verbose up ${num}

sqlc:
	cd ./db && sqlc generate

lint: .bin/golangci-lint | generate ## to lint the files
	.bin/golangci-lint run --config=.golangci-lint.yml ./...

.bin/golangci-lint:
	@if ! ( [ -f ".bin/golangci-lint" ] && ( [ "v$(shell .bin/golangci-lint version --format short)" = "$(LINTER_VERSION)" ] || [ "$(shell .bin/golangci-lint version --format short)" = "$(LINTER_VERSION)" ] ) ) ; then \
		if ( [ "$(shell which golangci-lint)" != "" ] && ( [ "v$(shell golangci-lint version --format short)" = "$(LINTER_VERSION)" ] || [ "$(shell golangci-lint version --format short)" = "$(LINTER_VERSION)" ] ) ); then \
			mkdir -p .bin; ln -s "$(shell which golangci-lint)" .bin ; \
		else \
			echo "--> Downloading from Github"; \
			curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b .bin $(LINTER_VERSION); \
		fi \
	fi

.bin/gofumpt:
	@if [ -z "$$(which gofumpt)" ]; then \
		go install mvdan.cc/gofumpt@latest; \
	else \
		mkdir -p .bin; ln -s "$$(which gofumpt)" $@; \
	fi

fmt: .bin/gofumpt ## to run `go fmt` on all source code
	gofumpt -l -w $(SRCS)

.SECONDEXPANSION:
$(MOCKED_FOLDERS): | .pre-check-go
	cd $(patsubst %/mocks,%,$@) && mockery --quiet --all --keeptree --outpkg mocks --output mocks

.SECONDEXPANSION:
$(MOCKED_FILES): $$(shell find $$(patsubst %/mocks,%,$$(patsubst %/mocks/,%,$$(dir $$@))) -maxdepth 1 -name "*.go") | $(MOCKED_FOLDERS)
	rm -rf $(dir $@)
	cd $(patsubst %/mocks,%,$(patsubst %/mocks/,%,$(dir $@))) && mockery --quiet --all --keeptree --outpkg mocks --output mocks

.pre-check-go:
	@if [ -z "$$(which mockery)" ]; then \
		go install github.com/vektra/mockery/v2@v2.14.0; \
	fi

# Variables
ROOT := git.cafebazaar.ir/divar/image/image
GO ?= go
GO_BIN_PATH ?= $(Go env GOPATH)/bin
GIT ?= git
COMMIT := $(shell $(GIT) rev-parse --short HEAD)
VERSION ?= $(strip $(if $(CI_COMMIT_TAG),$(CI_COMMIT_TAG),$(shell $(GIT) describe --tag 2> /dev/null || echo "$(COMMIT)")))
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
LINTER_VERSION ?= v1.51.0
