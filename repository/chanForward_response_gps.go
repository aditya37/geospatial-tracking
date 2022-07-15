package repository

import (
	"sync"

	"github.com/aditya37/geospatial-tracking/usecase"
)

type ChanTrackingForward struct {
	message chan usecase.ForwardTrackingPayload
}

type TrackingForward struct {
	RWMute          sync.RWMutex
	forwardTracking *ChanTrackingForward
	result          chan *usecase.ForwardTrackingPayload
}

func NewTrackingForward() *TrackingForward {
	tf := &TrackingForward{
		RWMute: sync.RWMutex{},
		forwardTracking: &ChanTrackingForward{
			message: make(chan usecase.ForwardTrackingPayload),
		},
		result: make(chan *usecase.ForwardTrackingPayload),
	}
	// run tracking forward
	go tf.run()
	return tf
}

// Publish....
func (tf *TrackingForward) Publish(message usecase.ForwardTrackingPayload) {
	tf.RWMute.Lock()
	// assign data to channel
	tf.forwardTracking.message <- message
	defer tf.RWMute.Unlock()
}

// Subscribe...
// Subsribe data from channel TrackingForward
func (tf *TrackingForward) Subscribe(callback func(m *usecase.ForwardTrackingPayload)) {
	go func() {
		for {
			msg := <-tf.result
			callback(msg)
		}
	}()
}

//
func (tf *TrackingForward) run() {
	for {
		select {
		case data := <-tf.forwardTracking.message:
			tf.result <- &data
		default:
			continue
		}
	}
}
