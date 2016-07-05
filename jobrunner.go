package jobrunner

import (
	"log"
	"reflect"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"gopkg.in/robfig/cron.v2"
)

// innerJob is jobrunner's job implementation
type innerJob struct {
	Name    string
	inner   cron.Job
	status  uint32
	Status  string
	Latency string
	running sync.Mutex
}

const unnamed = "(unnamed)"

func newJob(job cron.Job) *innerJob {

	name := reflect.TypeOf(job).Name()

	if name == "Func" {
		name = unnamed
	}

	return &innerJob{
		Name:  name,
		inner: job,
	}
}

func (j *innerJob) statusUpdate() string {

	if atomic.LoadUint32(&j.status) > 0 {
		j.Status = "RUNNING"
		return j.Status
	}

	j.Status = "IDLE"

	return j.Status
}

// Run invokes the current job, should only be run internally
func (j *innerJob) Run() {

	start := time.Now()

	// If the job panics, just print a stack trace.
	// Don't let the whole process die.
	defer func() {
		if err := recover(); err != nil {
			log.Panic(err, string(debug.Stack()))
		}
	}()

	// so only one instance can run at a time.
	j.running.Lock()
	defer j.running.Unlock()

	atomic.StoreUint32(&j.status, 1)
	j.statusUpdate()

	defer atomic.StoreUint32(&j.status, 0)
	defer j.statusUpdate()

	j.inner.Run()

	end := time.Now()
	j.Latency = end.Sub(start).String()

}
