package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spangenberg/alieninvasion/internal/genmap"
)

const (
	defaultHeight = 10
	defaultWidth  = 10
	defaultCities = 10
)

func newCmdGenerate() *cobra.Command {
	var cities, height, width int

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a new world",
		Long:  `Generate a new world.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if height < 1 {
				return fmt.Errorf("invalid height: %d", height)
			}
			if width < 1 {
				return fmt.Errorf("invalid width: %d", width)
			}
			if cities < 1 {
				return fmt.Errorf("invalid number of cities: %d", cities)
			}
			if cities > height*width {
				return fmt.Errorf("cities must be less than or equal to height * width (%d * %d = %d)", height, width, height*width)
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			earth := genmap.NewWorld(height, width)
			earth.PlaceCities(cities)
			cmd.SetOut(os.Stdout)
			earth.Output(func(s string) {
				cmd.Println(s)
			})
		},
	}
	cmd.Flags().IntVar(&height, "height", defaultHeight, "Height of the world")
	cmd.Flags().IntVar(&width, "width", defaultWidth, "Width of the world")
	cmd.Flags().IntVar(&cities, "cities", defaultCities, "Number of cities in the world")

	return cmd
}
