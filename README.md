# Alien Invasion

Alien Invasion is a CLI tool to simulate an alien invasion.

Getting Started & Documentation
-------------------------------

All documentation is available on the [Alien Invasion documentation](https://github.com/spangenberg/alieninvasion/tree/main/docs).

### Run locally

```sh
$ go run .
...
```

### Build

```sh
$ make build
...
```

### Test

```sh
$ make test
...
```

### Generating a map

This command will generate a decently sized map.

```sh
$ bin/alieninvasion generate --height 300 --width 300 --cities 50000 | grep = > 50k.map
```

### Running the simulation

```sh
$ bin/alieninvasion simulate --map-path 50k.map 10000
```

### Notes

- City names in the map file must not contain spaces, if they do the parsing will fail.
- Alien names can be randomly generated, but they are not guaranteed to be unique and are also hard to read in the output.
- The simulation will run until all aliens are dead or the maximum number of turns is reached.
- The random number generator is seeded with the current time, so the simulation will be different each time, for this use case it's good enough.
- Basic user validation is performed on the CLI, but it is not exhaustive.
- The test coverage is not exhaustive and only cover some core logic.
- The simulation runs sequentially, an experimental concurrency version is available in the [feat/concurrency](https://github.com/spangenberg/alieninvasion/tree/feat/concurrency) branch.
