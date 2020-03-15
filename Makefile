OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
MAKE_COLOR=\033[33;01m%-20s\033[0m

## all              : build, fmt, test
all: deps fmt test

## deps             : sync go mod deps
deps:
	@echo "$(OK_COLOR)--> Download go.mod dependencies$(NO_COLOR)"
	go mod download
	go mod vendor

## test             : run all tests
test:
	@echo "$(OK_COLOR)--> Running unit tests$(NO_COLOR)"
	go test --race --count=1 ./...

## fmt              : format go files
fmt:
	@echo "$(OK_COLOR)--> Formatting go files$(NO_COLOR)"
	go fmt ./...

## help             : show this help screen
help : Makefile
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*? "}; {printf "$(MAKE_COLOR) : %s\n", $$1, $$2}'

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: all deps test fmt clean help