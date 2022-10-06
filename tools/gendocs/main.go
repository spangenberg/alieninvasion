package main

import (
	"github.com/spangenberg/snakecharmer"

	"github.com/spangenberg/alieninvasion/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot())
}
