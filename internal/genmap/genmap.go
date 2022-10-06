package genmap

import (
	"fmt"
	"math/rand"
	"strings"
)

// World represents a world.
type World struct {
	matrix      [][]*city
	coordinates []coordinate
	cities      []*city
}

// NewWorld creates a new world with the given height and width.
func NewWorld(height, width int) *World {
	matrix := make([][]*city, width)

	coordinates := make([]coordinate, 0, width*height)

	for lat := 0; lat < width; lat++ {
		matrix[lat] = make([]*city, height)

		for lon := 0; lon < height; lon++ {
			coordinates = append(coordinates, coordinate{
				lat: lat,
				lon: lon,
			})
		}
	}

	return &World{
		matrix:      matrix,
		coordinates: coordinates,
	}
}

// Output prints the world to stdout.
func (w *World) Output(out func(s string)) {
	for _, cty := range w.cities {
		cty.north = w.lookup(cty.lat, cty.lon-1)
		cty.east = w.lookup(cty.lat+1, cty.lon)
		cty.south = w.lookup(cty.lat, cty.lon+1)
		cty.west = w.lookup(cty.lat-1, cty.lon)
		out(cty.mapFormat())
	}
}

// PlaceCities places the given number of cities in the world.
func (w *World) PlaceCities(cities int) {
	w.randomizeCoordinates(cities)

	var idx int
	for i := len(w.coordinates) - 1; i >= len(w.coordinates)-cities; i-- {
		idx++

		c := w.coordinates[i]
		cty := &city{
			name:       fmt.Sprintf("city%d", idx),
			coordinate: c,
		}
		w.matrix[c.lat][c.lon] = cty
		w.cities = append(w.cities, cty)
	}
}

// lookup returns the city at the given coordinate.
func (w *World) lookup(lat, lon int) *city {
	if lat < 0 {
		lat = len(w.matrix) - 1
	}

	if lat >= len(w.matrix) {
		lat = 0
	}

	if lon < 0 {
		lon = len(w.matrix[0]) - 1
	}

	if lon >= len(w.matrix[0]) {
		lon = 0
	}

	return w.matrix[lat][lon]
}

// randomizeCoordinates randomizes the coordinates.
// It iterates over the slice and swap each value with a random value but only as many times as the number of nCities.
func (w *World) randomizeCoordinates(cities int) {
	for i := len(w.coordinates); i > len(w.coordinates)-cities; i-- {
		randIndex := rand.Intn(i)
		w.coordinates[i-1], w.coordinates[randIndex] = w.coordinates[randIndex], w.coordinates[i-1]
	}
}

// city represents a city.
type city struct {
	name                     string
	north, east, south, west *city
	coordinate
}

// String returns a string representation of the city.
func (c *city) String() string {
	return c.name
}

// mapFormat returns a string representation of the city and its neighbors.
func (c *city) mapFormat() string {
	chunks := []string{c.name}
	if c.north != nil {
		chunks = append(chunks, fmt.Sprintf("north=%s", c.north))
	}

	if c.east != nil {
		chunks = append(chunks, fmt.Sprintf("east=%s", c.east))
	}

	if c.south != nil {
		chunks = append(chunks, fmt.Sprintf("south=%s", c.south))
	}

	if c.west != nil {
		chunks = append(chunks, fmt.Sprintf("west=%s", c.west))
	}

	return strings.Join(chunks, " ")
}

// coordinate represents a coordinate in the world.
type coordinate struct {
	lat, lon int
}
