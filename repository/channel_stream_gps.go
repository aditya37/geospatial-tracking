package repository

import (
	"sync"

	"github.com/aditya37/geospatial-tracking/usecase"
)

type Channel struct {
	message chan usecase.MqttGpsTrackingPayload
}
type GPSChannelStream struct {
	GPSTrackPayload *Channel
	mute            sync.RWMutex
	Result          chan usecase.MqttGpsTrackingPayload
}

func NewChannelStreamGPS() *GPSChannelStream {
	return &GPSChannelStream{
		mute: sync.RWMutex{},
		GPSTrackPayload: &Channel{
			message: make(chan usecase.MqttGpsTrackingPayload),
		},
		Result: make(chan usecase.MqttGpsTrackingPayload),
	}
}

func (gs *GPSChannelStream) StoreTrackingToChan(message usecase.MqttGpsTrackingPayload) {
	// set data to Channel
	gs.mute.Lock()
	gs.GPSTrackPayload.message <- message
	defer gs.mute.Unlock()
}

func (gs *GPSChannelStream) Run() {
	for {
		select {
		case d := <-gs.GPSTrackPayload.message:
			// expose data to exit
			gs.Result <- d

		}
	}

}
