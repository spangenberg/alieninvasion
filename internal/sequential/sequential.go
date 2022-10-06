package sequential

import (
	"fmt"
	"strings"

	"github.com/spangenberg/alieninvasion/internal"
)

const (
	// maxRounds is the maximum number of rounds to run the simulation for.
	maxRounds = 10000
	// minAliens is the minimum number of aliens to run the simulation for.
	minAliens = 2
)

// SimulateInvasion simulates the invasion.
func SimulateInvasion(cfg *internal.Config, world *internal.World, out func(s string)) error {
	myInvaders, err := newInvaders(cfg, world)
	if err != nil {
		return fmt.Errorf("failed to create invaders: %w", err)
	}

	// Run the simulation for a maximum of 10,000 rounds
	for i := 0; i < maxRounds; i++ {
		// At least two aliens must be in the world to pick a fight with each other
		if myInvaders.len() < minAliens {
			return nil
		}

		if myInvaders.len() > world.Len() {
			processCities(world, myInvaders, out)
		} else {
			processAliens(world, myInvaders, out)
		}
	}

	return nil
}

// checkCity checks if there are multiple aliens in the city and if so, picks a fight.
func checkCity(world *internal.World, invaders *invaders, city *internal.City) string {
	if city.LenAliens() > 1 {
		var aliens []string

		for _, myAlien := range city.Aliens {
			aliens = append(aliens, myAlien.String())
		}

		for _, myAlien := range city.Aliens {
			myAlienT, ok := myAlien.(*alien)
			if !ok {
				panic("alien is not of type *alien")
			}

			invaders.removeAlien(myAlienT)
		}

		city.Destroy()
		world.RemoveCity(city)

		return fmt.Sprintf("%s has been destroyed by %s and %s!", city,
			strings.Join(aliens[:len(aliens)-1], ", "), aliens[len(aliens)-1])
	}

	return ""
}

// processAliens processes all aliens in the world.
func processAliens(world *internal.World, invaders *invaders, out func(s string)) {
	for myAlien := range invaders.aliens {
		if str := checkCity(world, invaders, myAlien.city); str != "" {
			out(str)

			continue
		}

		if !myAlien.move() {
			invaders.removeAlien(myAlien)
		}
	}
}

// processCities processes all aliens in all cities.
func processCities(world *internal.World, invaders *invaders, out func(s string)) {
	for _, cty := range world.Cities {
		if str := checkCity(world, invaders, cty); str != "" {
			out(str)

			continue
		}

		if cty.LenAliens() == 1 {
			myAlien, ok := cty.Aliens[0].(*alien)
			if !ok {
				panic("alien is not of type *alien")
			}

			if !myAlien.move() {
				invaders.removeAlien(myAlien)
			}
		}
	}
}
