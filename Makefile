# https://stackoverflow.com/questions/18136918/how-to-get-current-relative-directory-of-your-makefile
DIR := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

.PHONY: all
all: gazelle
	@bazel build //...
	@make update-binary

# build/run
.PHONY: run
run: all
	@bazel run //



.PHONY: gazelle
gazelle:
	@bazel run //:gazelle

.PHONY: clean
clean:
	@bazel clean


# binary
.PHONY: update-binary
update-binary:
	mkdir -p bin
	ln -fs $(DIR)/bazel-bin/darwin_amd64_stripped/oidc-helper $(DIR)/bin/oidc-helper
