package entity

type QRDevice struct {
	Id          int64
	EventType   int
	DeviceId    string
	Description string
	QrFile      string
	Url         string
}
