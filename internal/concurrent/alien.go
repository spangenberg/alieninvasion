package concurrent

import (
	"fmt"
	"strings"

	"github.com/cip8/autoname"

	"github.com/spangenberg/alieninvasion/internal"
)

// maxRounds is the maximum number of rounds to run the simulation for.
const maxRounds = 10000

// alien is an alien in the world.
type alien struct {
	city  *internal.City
	name  string
	moves int
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

// move the alien to a random city.
func (a *alien) move(out func(s string)) bool {
	if a.city.Destroyed {
		return false
	}

	nextCity := a.city.RandomNeighbor()
	if nextCity == nil {
		return false
	}

	a.city.RemoveAlien()
	a.city = nextCity
	a.city.AddAlien(a)

	if a.city.LenAliens() > 1 {
		var aliens []string

		for _, myAlien := range a.city.Aliens {
			aliens = append(aliens, myAlien.String())
		}

		out(fmt.Sprintf("%s has been destroyed by %s and %s!", a.city,
			strings.Join(aliens[:len(aliens)-1], ", "), aliens[len(aliens)-1]))

		a.city.Destroy()

		return false
	}

	a.moves++

	if a.moves >= maxRounds {
		return false
	}

	return true
}
