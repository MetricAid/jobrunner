package jobrunner

import (
	"sync"

	"gopkg.in/robfig/cron.v2"
)

var (
	// Singleton instance of the underlying job scheduler.
	mainCron *cron.Cron
	once     sync.Once
)

func init() {
	once.Do(func() {
		mainCron = cron.New()
	})
}
