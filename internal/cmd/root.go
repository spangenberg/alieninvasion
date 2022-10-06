package cmd

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/spangenberg/alieninvasion/internal"
	"github.com/spangenberg/alieninvasion/internal/version"
)

// Execute executes the root command.
func Execute() {
	rand.Seed(time.Now().Unix())

	if err := NewCmdRoot().Execute(); err != nil {
		os.Exit(1)
	}
}

// NewCmdRoot creates a new root command.
func NewCmdRoot() *cobra.Command {
	cfg := new(internal.Config)
	cmd := &cobra.Command{
		Use:   "alieninvasion",
		Short: "alieninvasion is a CLI tool to simulate an alien invasion",
		Long: `alieninvasion is a CLI tool to simulate an alien invasion.
Find more information at: https://spangenberg.github.io/alieninvasion/`,
		Version: version.String(),
	}
	cmd.AddCommand(newCmdGenerate())
	cmd.AddCommand(newCmdSimulate(cfg))

	return cmd
}
