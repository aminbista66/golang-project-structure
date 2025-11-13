package tasks

import (
	"sync"
	"myapp/internal/app"
)

var wg sync.WaitGroup

// Go runs a background function safely and tracks it in the central WaitGroup.
// The name parameter is optional but helps with debugging/logging.
func Background(name string, fn func()) {
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				app.L.PrintError("panic recovered in background task ["+name+"]", map[string]string{})
			}
			wg.Done()
			app.L.PrintInfo("task ["+name+"] finished", map[string]string{})
		}()
		app.L.PrintInfo("starting background task ["+name+"]", map[string]string{})
		fn()
	}()
}

