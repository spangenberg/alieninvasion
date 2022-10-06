package sequential

import (
	"github.com/cip8/autoname"

	"github.com/spangenberg/alieninvasion/internal"
)

// alien is an alien in the world.
type alien struct {
	city *internal.City
	name string
}

// newAlien creates a new alien.
func newAlien(name string, city *internal.City) *alien {
	if name == "" {
		name = autoname.Generate("-")
	}

	myAlien := &alien{
		name: name,
		city: city,
	}
	city.AddAlien(myAlien)

	return myAlien
}

// String returns a string representation of the alien.
func (a *alien) String() string {
	return a.name
}

// move the alien to a random neighbor.
func (a *alien) move() bool {
	nextCity := a.city.RandomNeighbor()
	if nextCity == nil {
		return false
	}

	a.city.RemoveAlien()
	a.city = nextCity
	a.city.AddAlien(a)

	return true
}
