default: help

PROJECTNAME=$(shell basename "$(PWD)")

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## compile: compiles project in current system
compile: clean get generate fmt vet test

## watch: format, test and build project at go files modification
watch:
	@echo "  >  Watching go files..."
	@if type "ag" > /dev/null 2>&1; then if type "entr" > /dev/null 2>&1; then ag -l | entr make clean generate fmt vet test-colorized; else echo "Please install entr: http://eradman.com/entrproject/"; fi else echo "Please install silver searcher: https://github.com/ggreer/the_silver_searcher"; fi

# ---------------------------------------------------------------------------

clean:
	@echo "  >  Cleaning build cache"
	@-go clean

fmt:
	@echo "  >  Formatting code"
	@go fmt ./...

generate:
	@echo "  >  Go generate"
	@if type "stringer" > /dev/null 2>&1; then go generate; else GO111MODULE=off go get golang.org/x/tools/cmd/stringer && go generate; fi

get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get

test:
	@echo "  >  Executing unit tests"
	@go test -v -timeout 60s -race ./...

test-colorized:
	@echo "  >  Executing unit tests"
	@if type "richgo" > /dev/null 2>&1; then richgo test -v -timeout 60s -race ./...; else GO111MODULE=off go get github.com/kyoh86/richgo && richgo test -v -timeout 60s -race ./...; fi

vet:
	@echo "  >  Checking code with vet"
	@go vet ./...

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
