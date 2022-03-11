package cron

import (
	"errors"
	"sync"
	"time"
)

const (
	defaultTickerIntervalMS = 100
	defaultErrBufferSize    = 256
	defaultConcurrentMax    = 1
)

type Config struct {
	// TODO: Add Configs for Service

	TickerIntervalMS int
	ConcurrentMax    int
	ErrBufferSize    int
}

type Service struct {
	// TODO: Add Service Parameters

	tickingInterval time.Duration
	running         chan struct{}
	concurrent      int
	concurrentMax   int
	lock            sync.Mutex

	errBufferSize int
	errs          chan error
}

func NewService(config Config) (*Service, error) {
	// Check Config for Invalid Settings
	// TODO: Add

	// Set Defaults
	if config.TickerIntervalMS <= 0 {
		config.TickerIntervalMS = defaultTickerIntervalMS
	}
	if config.ConcurrentMax <= 0 {
		config.ConcurrentMax = defaultConcurrentMax
	}
	if config.ErrBufferSize <= 0 {
		config.ErrBufferSize = defaultErrBufferSize
	}
	// TODO: Add additional defaults, if necessary

	return &Service{
		// TODO: Initialize Service Parameters
		tickingInterval: time.Duration(config.TickerIntervalMS) * time.Millisecond,
		lock:            sync.Mutex{},

		errBufferSize: config.ErrBufferSize,
	}, nil
}

func (svc *Service) Start() error {
	// Initialize the running channel
	err := svc.initStart()
	if err != nil {
		return err
	}

	// Note: This creates a separate ticker for each sync thread.
	//		 To create a shared thread an utilize a 'pool' of routines,
	//		 move the ticker to part of the struct, and initialize it only if nil
	ticker := time.NewTicker(svc.tickingInterval)
	for {
		err := svc.singleOperation()
		if err != nil {
			svc.errs <- err
		}
		select {
		case <-svc.running:
			// this case is selected if `Stop` is called for the serivce
			return nil
		case <-ticker.C:
			// continue to the next iteration
		}
	}

}

func (svc *Service) initStart() error {
	svc.lock.Lock()
	defer svc.lock.Unlock()

	if svc.concurrent >= svc.concurrentMax {
		return errors.New("could not initialize new routine; maximum concurrent threads reached")
	}

	svc.concurrent++

	// Init Running Channel, if not already init
	if svc.running == nil {
		svc.running = make(chan struct{})
	}
	select {
	case <-svc.running:
		// the running channel has been closed, we can start it again
		svc.running = make(chan struct{})
	default:
		// if this case is selected, the running channel already exists
	}

	// Init Err Channel, if not already init
	if svc.errs == nil {
		svc.errs = make(chan error, svc.errBufferSize)
	}
	select {
	case <-svc.errs:
		// the error channel has been closed, we can start it again
		svc.errs = make(chan error, svc.errBufferSize)
	default:
		// if this case is selected, the error channel is already running
	}

	return nil
}

func (svc *Service) singleOperation() error {
	// Regular Operations Here
	// TODO
	return nil
}

func (svc *Service) Stop() {
	svc.lock.Lock()
	defer svc.lock.Unlock()

	// Using a select, we can see if the service has already been stopped
	//	This allows us to stop it from many places at the same time
	select {
	case <-svc.running:
		// Channel is already closed, no other cleanup should be required
		return
	default:
	}

	// Close the `running` channel to signal to all threads that they should stop
	close(svc.running)
	// Close the `errs` channel to cleanup external err logging routines
	close(svc.errs)

	// Add Cleanup Tasks Here, ex:
	// 	Stop a group ticker
	//  Close consumers or DB connections
}

func (svc *Service) Errors() <-chan error {
	return svc.errs
}
