package sequential

import (
	"fmt"

	"github.com/spangenberg/alieninvasion/internal"
)

// invaders is a list of aliens.
type invaders struct {
	aliens map[*alien]struct{}
}

// newInvaders creates a new list of aliens.
func newInvaders(cfg *internal.Config, world *internal.World) (*invaders, error) {
	myInvaders := &invaders{
		aliens: make(map[*alien]struct{}, cfg.Aliens),
	}

	for i := 0; i < cfg.Aliens; i++ {
		if err := myInvaders.addAlien(world, cfg.GenerateAlienNames); err != nil {
			return nil, err
		}
	}

	return myInvaders, nil
}

// addAlien adds an alien to the world.
func (i *invaders) addAlien(world *internal.World, generateAlienNames bool) error {
	cty := world.RandomCity()
	if cty == nil {
		return fmt.Errorf("no cities")
	}

	var name string
	if !generateAlienNames {
		name = fmt.Sprintf("alien %d", len(i.aliens)+1)
	}

	myAlien := newAlien(name, cty)
	i.aliens[myAlien] = struct{}{}

	return nil
}

// len returns the number of aliens.
func (i *invaders) len() int {
	return len(i.aliens)
}

// removeAlien removes an alien from the world.
func (i *invaders) removeAlien(alien *alien) {
	delete(i.aliens, alien)
}
