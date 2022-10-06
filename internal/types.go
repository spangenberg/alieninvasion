package internal

// Alien is a type that represents an alien.
type Alien interface {
	// String returns the name of the alien.
	String() string
}

// Config is the configuration for the world.
type Config struct {
	// Aliens is the number of aliens in the world.
	Aliens int
	// GenerateAlienNames determines if alien names should be generated.
	GenerateAlienNames bool
	// MapPath is the path to the map file.
	MapPath string
}

// Directions is a set of directions.
type Directions struct {
	// North is the name of the city in the north.
	North string
	// East is the name of the city in the east.
	East string
	// South is the name of the city in the south.
	South string
	// West is the name of the city in the west.
	West string
}
