package internal

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

// City is a city in the world.
type City struct {
	// Aliens is the list of aliens in the city.
	Aliens []Alien
	// Destroyed is true if the city is destroyed.
	Destroyed bool
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

	sync.RWMutex
}

// AddAlien adds an alien to the city.
func (c *City) AddAlien(alien Alien) {
	c.Lock()
	defer c.Unlock()

	c.Aliens = append(c.Aliens, alien)
}

// Destroy destroys the city.
func (c *City) Destroy() {
	c.Lock()
	defer c.Unlock()

	c.Aliens = nil
	c.Destroyed = true

	if c.North != nil {
		c.North.Lock()
		defer c.North.Unlock()
		c.North.South = nil
		c.North = nil
	}

	if c.East != nil {
		c.East.Lock()
		defer c.East.Unlock()
		c.East.West = nil
		c.East = nil
	}

	if c.South != nil {
		c.South.Lock()
		defer c.South.Unlock()
		c.South.North = nil
		c.South = nil
	}

	if c.West != nil {
		c.West.Lock()
		defer c.West.Unlock()
		c.West.East = nil
		c.West = nil
	}
}

// LenAliens returns the number of aliens in the city.
func (c *City) LenAliens() int {
	c.RLock()
	defer c.RUnlock()

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
		c.North.RLock()
		defer c.North.RUnlock()
		neighbors = append(neighbors, c.North)
	}

	if c.East != nil {
		c.East.RLock()
		defer c.East.RUnlock()
		neighbors = append(neighbors, c.East)
	}

	if c.South != nil {
		c.South.RLock()
		defer c.South.RUnlock()
		neighbors = append(neighbors, c.South)
	}

	if c.West != nil {
		c.West.RLock()
		defer c.West.RUnlock()
		neighbors = append(neighbors, c.West)
	}

	if len(neighbors) == 0 {
		return nil
	}

	return neighbors[rand.Intn(len(neighbors))]
}

// RemoveAlien removes all aliens from the city.
func (c *City) RemoveAlien() {
	c.Lock()
	defer c.Unlock()

	c.Aliens = nil
}
