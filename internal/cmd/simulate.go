package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/spangenberg/alieninvasion/internal"
	"github.com/spangenberg/alieninvasion/internal/concurrent"
	"github.com/spangenberg/alieninvasion/internal/sequential"
)

const minAliens = 2

func newCmdSimulate(cfg *internal.Config) *cobra.Command {
	var concurrently bool

	cmd := &cobra.Command{
		Use:   "simulate [number of aliens]",
		Short: "Simulates an alien invasion",
		Long:  `Simulates an alien invasion.`,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if cfg.Aliens, err = strconv.Atoi(args[0]); err != nil {
				return fmt.Errorf("invalid number of aliens: %w", err)
			}
			if cfg.Aliens < minAliens {
				return fmt.Errorf("invalid number of aliens: %d\nplease specify at least two alians", cfg.Aliens)
			}

			return nil
		},
		Run: handleError(func(cmd *cobra.Command, args []string) error {
			earth, err := internal.NewWorld(cfg)
			if err != nil {
				return fmt.Errorf("failed to create world: %w", err)
			}

			cmd.Print("Starting simulation...\n\n")
			out := func(s string) {
				cmd.Println(s)
			}
			if concurrently {
				if err = concurrent.SimulateInvasion(cfg, earth, out); err != nil {
					return fmt.Errorf("failed to simulate invasion: %w", err)
				}
			} else {
				if err = sequential.SimulateInvasion(cfg, earth, out); err != nil {
					return fmt.Errorf("failed to simulate invasion: %w", err)
				}
			}

			cmd.Print("\nWorld after simulation:\n\n")
			earth.PrintWorld(out)

			return nil
		}),
	}
	cmd.Flags().BoolVar(&concurrently, "concurrent", false, "run simulation concurrently")
	cmd.Flags().BoolVar(&cfg.GenerateAlienNames, "generate-alien-names", false, "Generate random alien names")
	cmd.Flags().StringVar(&cfg.MapPath, "map-path", "", "Location of the file with the world map")

	return cmd
}
