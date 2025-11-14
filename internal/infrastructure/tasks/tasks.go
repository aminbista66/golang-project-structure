package tasks

import (
	jsonlog "myapp/internal/infrastructure/logger/jsonlog"
	"sync"
)

type TaskManager struct {
	wg     *sync.WaitGroup
	Logger *jsonlog.Logger
}

func New(logger *jsonlog.Logger, wg *sync.WaitGroup) *TaskManager {
	return &TaskManager{
		Logger: logger,
		wg:     wg,
	}
}

// Go runs a background function safely and tracks it in the central WaitGroup.
// The name parameter is optional but helps with debugging/logging.
func (tm *TaskManager) Background(name string, fn func()) {
	tm.wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				tm.Logger.PrintError("panic recovered in background task ["+name+"]", map[string]string{})
			}
			tm.wg.Done()
			tm.Logger.PrintInfo("task ["+name+"] finished", map[string]string{})
		}()
		tm.Logger.PrintInfo("starting background task ["+name+"]", map[string]string{})
		fn()
	}()
}
