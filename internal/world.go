package internal

import (
	"fmt"
	"math/rand"
	"sync"
)

// World represents the world in which the aliens move around and fight.
type World struct {
	// Cities is a map of all cities in the world.
	Cities map[int]*City

	cityNameLookup map[string]int
	m              sync.RWMutex
}

// NewWorld creates a new world.
func NewWorld(cfg *Config) (*World, error) {
	world := &World{
		Cities:         make(map[int]*City),
		cityNameLookup: make(map[string]int),
	}
	if err := ParseMapFile(cfg.MapPath, world.addCity); err != nil {
		return nil, fmt.Errorf("failed to parse map file: %w", err)
	}

	for _, city := range world.Cities {
		if city.directions == nil {
			continue
		}

		err := world.processCity(city)
		if err != nil {
			return nil, err
		}

		city.directions = nil
	}

	return world, nil
}

// Len returns the number of cities in the world.
func (w *World) Len() int {
	w.m.RLock()
	defer w.m.RUnlock()

	return len(w.Cities)
}

// PrintWorld prints the world.
func (w *World) PrintWorld(out func(s string)) {
	w.m.RLock()
	defer w.m.RUnlock()

	for _, city := range w.Cities {
		if city.Destroyed {
			continue
		}

		out(city.MapFormat())
	}
}

// RandomCity returns a random city from the world.
func (w *World) RandomCity() *City {
	w.m.RLock()
	defer w.m.RUnlock()

	if len(w.Cities) == 0 {
		return nil
	}

	return w.Cities[rand.Intn(len(w.Cities))]
}

// RemoveCity removes the city from the world and itself from all neighboring cities.
func (w *World) RemoveCity(city *City) {
	w.m.Lock()
	defer w.m.Unlock()

	delete(w.Cities, w.cityNameLookup[city.name])
	delete(w.cityNameLookup, city.name)
}

// addCity adds a city to the world.
func (w *World) addCity(name string, directions *Directions) error {
	if name == "" {
		return fmt.Errorf("invalid city name: %s", name)
	}

	if _, ok := w.cityNameLookup[name]; ok {
		return fmt.Errorf("city already exists: %s", name)
	}

	city := &City{
		name:       name,
		directions: directions,
	}
	w.cityNameLookup[name] = len(w.Cities)
	w.Cities[len(w.Cities)] = city

	return nil
}

// getCity returns the city with the given name.
func (w *World) getCity(name string) *City {
	if idx, ok := w.cityNameLookup[name]; ok {
		return w.Cities[idx]
	}

	return nil
}

// processCity processes a city's directions.
// It looks up the city's directions and sets the city's neighbours.
// It also adds the city to the neighbours' neighbours and cleans their directions.
func (w *World) processCity(city *City) error {
	if city.directions.North != "" {
		if city.North = w.getCity(city.directions.North); city.North == nil {
			return fmt.Errorf("city not found: %s", city.directions.North)
		}

		city.North.South = city
		if city.North.directions != nil {
			city.North.directions.South = ""
		}
	}

	if city.directions.East != "" {
		if city.East = w.getCity(city.directions.East); city.East == nil {
			return fmt.Errorf("city not found: %s", city.directions.East)
		}

		city.East.West = city
		if city.East.directions != nil {
			city.East.directions.West = ""
		}
	}

	if city.directions.South != "" {
		if city.South = w.getCity(city.directions.South); city.South == nil {
			return fmt.Errorf("city not found: %s", city.directions.South)
		}

		city.South.North = city
		if city.South.directions != nil {
			city.South.directions.North = ""
		}
	}

	if city.directions.West != "" {
		if city.West = w.getCity(city.directions.West); city.West == nil {
			return fmt.Errorf("city not found: %s", city.directions.West)
		}

		city.West.East = city
		if city.West.directions != nil {
			city.West.directions.East = ""
		}
	}

	return nil
}
