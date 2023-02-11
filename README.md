# Geospatial Tracking Or Device Service


# mqtt payload and topic

Last Change: 2023-11-02

Kontrak dan topic name untuk broker mqtt


## MQTT Topic

|Topic                  |Description                                            |
|-----------------------|-------------------------------------------------------|
|/device/req/register   |Topic untuk request atau action register device        |
|/device/resp/register  |Topic untuk response register device                   |
|/device/req/tracking   |Topic for record gps tracking from device              |
|/device/geofence/detect|Topic for publish device detected in geofenece area    |
|/device/logs           |Topic for publish device log like heartbeat etc..      |
|/device/resp/tracking  |Topic for response gps tracking and will               |
|/device/req/{device_id}/sensor/{sensor_id}|Topic for publish sensor data from iot device or device |
|/device/resp/sensor/{device_id}|Topic for response record sensor               |

## MQTT Message payload per topic

Payload or contract message pertopic mqtt. for this version message used with native json or struct in Golang, in next version will migrate to protbuf.

### Payload Topic : /device/req/register
topic for request register device 

```
{
    "device_id":string,
    "mac_address":string,
    "device_type":int,
    "chip_id":string,
    "i2c_address":string,
    "timestamp":int64,
    "network_mode":string,
    "phone_no":string,
    "imei":string,
    "imsi":"imsi",
    "sim_operator":int64,
    "apn":string,
    "embedded_sensor":[]int
}
```

in this version supported network mode:

```
const (
	NETWORK_MODE_WLAN        = "WLAN"
	NETWORK_MODE_MOBILE_DATA = "MOBILE_DATA"
)
```

### Payload Topic : /device/resp/register
topic for response register device 

```
{
    "device_id":string,
    "status":string,
    "message":string,
    "valid_sensor_id":[]int,
    "count_sensor_not_valid":int,
}
```

in this version supported status:

```

var (
	StatusValidDeviceId            Status = "VALID_DEVICE_ID"
	StatusSuccessRegister          Status = "SUCCESS_REGISTER_DEVICE"
	StatusFailedRegister           Status = "FAILED_REGISETER_DEVICE"
	StatusGPSTrackingStart         Status = "START_RECORD_TRACKING"
	StatusGPSTrackingStop          Status = "STOP"
	StatusGPSTracingRecordTracking Status = "TRACKING_RECORDED"
	StatusLowSignal                Status = "LOW_SIGNAL"
	StatusHeartBeat                Status = "HEARTBEAT"
)
```

### Payload Topic : /device/req/tracking

payload for record gps tracking

```
{
    "device_id":string ,
    "speed":float64,
    "status":string ,
    "temp":float64,
    "lat":float64,
    "long":float64,
    "signal":float64,
    "angle":float64,
    "altitude":float64,
    "timestamp":int64,
}
```

### Payload Topic : /device/geofence/detect
paylod for detect device inside or in gefence area
```
{
    "device_id":int64,
    "detect":string ,
    "lat":float64,
    "long":float64,
    "detect_time":int64  
}
```

### Payload Topic : /device/logs
payload for publish hearbeat or logs from device   
```
{
    "device_id":string,
    "status":string,
    "reason":string,
    "signal_strength":float64,
    "recorded_at":int64
}
```

### Payload Topic : /device/resp/tracking
topic for response gps tracking with detail
```
{
    "device_id":string ,
    "device_type":int,
    "id":int64,
    "status":string ,
    "message":Message,
    "gps_data":GPSData,
    "sensor_data":Sensor,
}
```

Child payload or struct

### Payload: Message
```
{
   "value":string,
   "reason":string
}
```
### Payload: GPSData
struct contain data from gps
```
{
    "lat":float64,
    "long":float64,
    "altitude":float64,
    "speed":float64,
    "angle":float64
}
```

### Payload: Sensor
struct contain data sensor
```
{
    "temp":float64,
	"signal":float64
}
```

### Payload Topic : /device/req/{device_id}/sensor/{sensor_id}
payload for publish data sensor non gps tracking

```
{
    "device_id":string,
    "sensor_id":int,
    "value":float64,
    "timestamp":int,
    "message":string,
    "code":int,
}
```

### Payload Topic : /device/resp/sensor/{device_id}
payload for response record sensor
```
{
    "device_id":string,
    "message":string,
    "code":int,
}
```

response code in response record sensor
```
// enum event code...
type EventCode int
const (
	// sending sensor data from device to backend
	SendSensorData                   EventCode = 1
	BadRequest                                 = 2
	DeviceIdNotFound                           = 3
	UserRFIDAndDeviveNotPairedToUser           = 4
	DeviceNotPairedToUser                      = 5
)
```