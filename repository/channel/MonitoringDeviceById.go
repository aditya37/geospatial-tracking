package channel

import (
	"log"
	"sync"

	"github.com/aditya37/geospatial-tracking/proto"
)

type (
	MonitoringDeviceByIdClient struct {
		StreamerId string
		Streamer   proto.Geotracking_MonitoringDeviceByIdServer
		Data       *proto.ResponseMonitoringDeviceById
	}
	MonitoringDeviceByIdPool struct {
		mutex    sync.RWMutex
		register chan *MonitoringDeviceByIdClient
		clients  map[string]*MonitoringDeviceByIdClient
		send     chan *MonitoringDeviceByIdClient
		close    chan *MonitoringDeviceByIdClient
	}
)

func NewMonitoringDeviceById() *MonitoringDeviceByIdPool {
	return &MonitoringDeviceByIdPool{
		mutex:    sync.RWMutex{},
		register: make(chan *MonitoringDeviceByIdClient),
		clients:  map[string]*MonitoringDeviceByIdClient{},
		send:     make(chan *MonitoringDeviceByIdClient),
		close:    make(chan *MonitoringDeviceByIdClient),
	}
}

// Register Channel...
func (mp *MonitoringDeviceByIdPool) Register(client *MonitoringDeviceByIdClient) {
	mp.mutex.Lock()
	mp.register <- client
	defer mp.mutex.Unlock()
}

// Set
func (mp *MonitoringDeviceByIdPool) Send(data MonitoringDeviceByIdClient) {
	mp.mutex.Lock()
	mp.send <- &data
	defer mp.mutex.Unlock()
}

// Close...
func (mp *MonitoringDeviceByIdPool) Close(client *MonitoringDeviceByIdClient) {
	mp.mutex.Lock()
	mp.close <- client
	defer mp.mutex.Unlock()
}

// Run...
func (mp *MonitoringDeviceByIdPool) Run() {
	for {
		select {
		case cl := <-mp.register:

			mp.clients[cl.StreamerId] = cl
			log.Printf(
				"Streamer ID: %s Joined Current Client: %d",
				cl.StreamerId,
				len(mp.clients),
			)
		case c := <-mp.close:
			if l, ok := mp.clients[c.StreamerId]; ok {
				log.Printf("Streamer ID : %s Leave", l.StreamerId)
				delete(mp.clients, l.StreamerId)
			}
		case d := <-mp.send:
			if cl, ok := mp.clients[d.StreamerId]; ok {
				cl.Streamer.Send(d.Data)
			}
		}
	}
}
