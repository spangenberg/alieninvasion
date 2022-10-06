#!/bin/bash

set -euo pipefail

if [ "${STATIC_ENABLED:=1}" == "0" ]; then
	unset STATIC_ENABLED
fi

ldflags=$(go run github.com/ahmetb/govvv@v0.3.0 build -flags -pkg github.com/spangenberg/alieninvasion/internal/version)

set -x

env CGO_ENABLED="${CGO_ENABLED:-0}" go build ${FORCE_REBUILD:+-a} -tags netgo -ldflags "$ldflags${DEBUG_ENABLED:+ -w}${STATIC_ENABLED:+ -extldflags \"-static\"}" -o "$OUTPUT" "$PACKAGE"
