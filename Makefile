export DOCS_PATH=docs
export OUTPUT=bin/alieninvasion
export PACKAGE=github.com/spangenberg/alieninvasion

.PHONY: build
# Build compiles this package
build:
	@scripts/build.sh

.PHONY: docs
# Generate docs
docs:
	@scripts/docs.sh

.PHONY: test
# Run tests
test:
	@scripts/test.sh
