package system

import (
	"context"
	"sync"
	"time"

	"github.com/coreos/go-systemd/v22/daemon"
)

type ControlLoop struct {
	service      Service
	ctx          context.Context
	triggerMutex sync.Mutex
	runMutex     sync.Mutex
	interval     time.Duration
	handler      func() error
	running      bool
	counter      int
}

func NewControlLoop(
	service Service,
	ctx context.Context,
	interval time.Duration,
	handler func() error,
) *ControlLoop {
	return &ControlLoop{
		service:  service,
		ctx:      ctx,
		interval: interval,
		handler:  handler,
		running:  false,
		counter:  0,
	}
}

func (loop *ControlLoop) incrementCounter() {
	loop.triggerMutex.Lock()
	defer loop.triggerMutex.Unlock()
	loop.counter++
}

func (loop *ControlLoop) resetCounter() {
	loop.triggerMutex.Lock()
	defer loop.triggerMutex.Unlock()
	loop.counter = 0
}

func (loop *ControlLoop) Trigger() {
	if loop.running {
		loop.incrementCounter()
	} else {
		loop.run()
		if loop.counter > 0 {
			loop.resetCounter()
			loop.Trigger()
		}
	}
}

func (loop *ControlLoop) run() {
	// this means that only 1 version this of function can be running at a time
	loop.runMutex.Lock()
	defer loop.runMutex.Unlock()
	loop.running = true
	err := loop.handler()
	loop.running = false
	if err != nil {
		Error(loop.service, "error running control loop handler", err)
	}
}

func (loop *ControlLoop) Start(runInitially bool) error {
	ticker := time.NewTicker(loop.interval)

	if err := notifyWatchdogReady(); err != nil {
		Error(loop.service, "unable to notify watchdog: ready", err)
	}

	if runInitially {
		err := loop.handler()
		if err != nil {
			return err
		}
	}

	go func() {
		for {
			select {
			case <-loop.ctx.Done():
				return
			case <-ticker.C:
				// Notify systemd watchdog that we are alive
				if err := notifyWatchdogHeartbeat(); err != nil {
					Error(loop.service, "unable to notify watchdog: heartbeat", err)
				}

				loop.Trigger()
			}
		}
	}()

	return nil
}

func notifyWatchdogReady() error {
	// Notify systemd that the service is ready
	_, err := daemon.SdNotify(false, daemon.SdNotifyReady)
	return err
}

func notifyWatchdogHeartbeat() error {
	_, err := daemon.SdNotify(false, daemon.SdNotifyWatchdog)
	return err
}
