#!/bin/bash

set -euxo pipefail

go run "$PACKAGE/tools/gendocs" --dir "$DOCS_PATH"
