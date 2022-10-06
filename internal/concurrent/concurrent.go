package concurrent

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/spangenberg/alieninvasion/internal"
)

// SimulateInvasion simulates the invasion.
func SimulateInvasion(cfg *internal.Config, world *internal.World, out func(s string)) error {
	channel := make(chan *alien, cfg.Aliens)
	done := make(chan struct{}, cfg.Aliens)

	for i := 0; i < cfg.Aliens; i++ {
		cty := world.RandomCity()
		if cty == nil {
			return fmt.Errorf("no cities")
		}

		var name string
		if !cfg.GenerateAlienNames {
			name = fmt.Sprintf("alien %d", i+1)
		}

		channel <- newAlien(name, cty)
	}

	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		wg.Add(1)
		go worker(ctx, &wg, channel, done, out)
	}

	go func() {
		var count int
		for {
			select {
			case <-done:
				count++
				if count >= cfg.Aliens {
					cancel()
					return
				}
			}
		}
	}()

	wg.Wait()

	return nil
}

func worker(ctx context.Context, wg *sync.WaitGroup, aliens chan *alien, done chan struct{}, out func(s string)) {
	defer wg.Done()

	for {
		select {
		case myAlien, ok := <-aliens:
			if !ok {
				return
			}

			if myAlien.move(out) {
				aliens <- myAlien
			} else {
				done <- struct{}{}
			}
		case <-ctx.Done():
			return
		}
	}
}
