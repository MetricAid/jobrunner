package jobrunner

import "gopkg.in/robfig/cron.v2"

// Start starts the CRON
func Start() {
	mainCron.Start()
}

// Stop ALL active jobs from running at the next scheduled time
func Stop() {
	go mainCron.Stop()
}

// Remove a specific job from running
// Get EntryID from the list job entries jobrunner.Entries()
// If job is in the middle of running, once the process is finished it will be removed
func Remove(id cron.EntryID) {
	mainCron.Remove(id)
}
