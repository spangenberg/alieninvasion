package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// lineChunks is the number of chunks that a line should be split into.
	lineChunks = 5
	// directionChunks is the number of chunks that a direction should be split into.
	directionChunks = 2
)

// ParseMapFile parses a map file.
func ParseMapFile(name string, apply func(name string, directions *Directions) error) (err error) {
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer func() {
		errFile := file.Close()
		if errFile != nil {
			if err != nil {
				err = fmt.Errorf("%w ; %v", err, errFile)
			} else {
				err = errFile
			}
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if err = processLine(line, apply); err != nil {
			return err
		}
	}

	return nil
}

// processLine processes a line from a map file.
func processLine(line string, apply func(name string, directions *Directions) error) error {
	chunks := strings.SplitN(line, " ", lineChunks)
	if len(chunks) > 0 {
		if chunks[0] == "" {
			return fmt.Errorf("invalid line: %s", line)
		}
	}

	var directions Directions

	for _, direction := range chunks[1:] {
		dir := strings.SplitN(direction, "=", directionChunks)
		if len(dir) > 1 {
			if dir[1] == "" {
				return fmt.Errorf("invalid line: %s", line)
			}
		}

		switch dir[0] {
		case "north":
			directions.North = dir[1]
		case "east":
			directions.East = dir[1]
		case "south":
			directions.South = dir[1]
		case "west":
			directions.West = dir[1]
		default:
			return fmt.Errorf("invalid direction: %#v", direction)
		}
	}

	return apply(chunks[0], &directions)
}
