package internal

import (
	"fmt"
	"math/rand"
	"strings"
)

// City is a city in the world.
type City struct {
	// Aliens is the list of aliens in the city.
	Aliens []Alien
	// North is the city to the north.
	North *City
	// East is the city to the east.
	East *City
	// South is the city to the south.
	South *City
	// West is the city to the west.
	West *City

	directions *Directions
	name       string
}

// AddAlien adds an alien to the city.
func (c *City) AddAlien(alien Alien) {
	c.Aliens = append(c.Aliens, alien)
}

// Destroy destroys the city.
func (c *City) Destroy() {
	c.Aliens = nil

	if c.North != nil {
		c.North.South = nil
		c.North = nil
	}

	if c.East != nil {
		c.East.West = nil
		c.East = nil
	}

	if c.South != nil {
		c.South.North = nil
		c.South = nil
	}

	if c.West != nil {
		c.West.East = nil
		c.West = nil
	}
}

// LenAliens returns the number of aliens in the city.
func (c *City) LenAliens() int {
	return len(c.Aliens)
}

// MapFormat returns a string representation of the city in map format.
func (c *City) MapFormat() string {
	chunks := []string{c.String()}
	if c.North != nil {
		chunks = append(chunks, fmt.Sprintf("north=%s", c.North))
	}

	if c.East != nil {
		chunks = append(chunks, fmt.Sprintf("east=%s", c.East))
	}

	if c.South != nil {
		chunks = append(chunks, fmt.Sprintf("south=%s", c.South))
	}

	if c.West != nil {
		chunks = append(chunks, fmt.Sprintf("west=%s", c.West))
	}

	return strings.Join(chunks, " ")
}

// String returns a string representation of the city.
func (c *City) String() string {
	return c.name
}

// RandomNeighbor returns a random neighbor.
func (c *City) RandomNeighbor() *City {
	var neighbors []*City

	if c.North != nil {
		neighbors = append(neighbors, c.North)
	}

	if c.East != nil {
		neighbors = append(neighbors, c.East)
	}

	if c.South != nil {
		neighbors = append(neighbors, c.South)
	}

	if c.West != nil {
		neighbors = append(neighbors, c.West)
	}

	if len(neighbors) == 0 {
		return nil
	}

	return neighbors[rand.Intn(len(neighbors))]
}

// RemoveAlien removes all aliens from the city.
func (c *City) RemoveAlien() {
	c.Aliens = nil
}
