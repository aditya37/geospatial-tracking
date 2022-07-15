package repository

import (
	"sync"

	"github.com/aditya37/geospatial-tracking/usecase"
)

type Channel struct {
	// payload or struct for sent callback to device
	streamGPSData chan usecase.MqttGpsTrackingPayload
	err           chan error
}
type GPSChannelStream struct {
	GPSTrackPayload *Channel
	mute            sync.RWMutex
	StreamGPSTrack  chan usecase.MqttGpsTrackingPayload
	ChanStreamError chan error
	isDone          chan bool
}

func NewChannelStreamGPS() *GPSChannelStream {
	return &GPSChannelStream{
		mute: sync.RWMutex{},
		GPSTrackPayload: &Channel{
			streamGPSData: make(chan usecase.MqttGpsTrackingPayload),
			err:           make(chan error),
		},
		StreamGPSTrack:  make(chan usecase.MqttGpsTrackingPayload),
		ChanStreamError: make(chan error),
		isDone:          make(chan bool),
	}
}

// StoreTrackingToStreamChan....
func (gs *GPSChannelStream) StoreTrackingToStreamChan(streamData usecase.MqttGpsTrackingPayload, err error) {
	gs.mute.Lock()
	gs.isDone <- false
	gs.GPSTrackPayload.streamGPSData <- streamData
	gs.GPSTrackPayload.err <- err
	defer gs.mute.Unlock()
}

// Done...
func (gs *GPSChannelStream) Done() {
	gs.isDone <- true
}

//
func (gs *GPSChannelStream) Run() {
	for {
		select {
		case streamer := <-gs.GPSTrackPayload.streamGPSData:
			gs.StreamGPSTrack <- streamer
		case chanStreamErr := <-gs.GPSTrackPayload.err:
			if chanStreamErr != nil {
				gs.ChanStreamError <- chanStreamErr
			}
		case done := <-gs.isDone:
			if !done {
				continue
			}
			gs.StreamGPSTrack <- usecase.MqttGpsTrackingPayload{}
		}
	}

}
