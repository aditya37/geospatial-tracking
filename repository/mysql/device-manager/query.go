package devicemanager

const (
	mysqlQueryInsertDevice = `INSERT INTO mst_device(
			 device_id,
			 mac_address,
			 device_type,
			 chip_id,
			 i2c_address,
			 created_at
	   ) VALUES(?,?,?,?,?,NOW())`
	mysqlQueryInsertTracking = `INSERT INTO trx_gps_tracking(
		device_id,
		signal_strength,
		speed,
		status,
		temp,
		waypoints,
		modified_at
	)VALUES(?,?,?,?,?,?,NOW())`
	mysqlQueryUpdateTracking = `UPDATE trx_gps_tracking SET 
		modified_at=NOW(),
		waypoints=JSON_ARRAY_INSERT(waypoints,'$.geometry.coordinates[0]',JSON_ARRAY(?,?)),
		status=?,
		temp=?,
		speed=?,
		signal_strength=?
	WHERE id =? AND device_id = ?`
	mysqlQueryInsertDeviceLog = `INSERT INTO trx_device_log(
		   device_id,
		   status,
		   reason,
		   signal_strength,
		   recorded_at
	)VALUES(?,?,?,?,?)`
	mysqlQueryInsertDeviceDetect   = "INSERT INTO trx_detect_device(device_id,detect,lat,`long`,detected_at)VALUES(?,?,?,?,?)"
	mysqlQueryGetCountDeviceDetect = `SELECT
		count(td.detect) AS count,
		md.device_id,
		mt.type,(
			SELECT detected_at FROM trx_detect_device WHERE device_id = md.id LIMIT 1
		) AS last_detect
   	FROM trx_detect_device td 
   	INNER JOIN mst_device md ON md.id = td.device_id
   	INNER JOIN mst_device_type mt ON md.device_type = mt.id GROUP BY md.device_id`
	// Read....
	mysqlQueryGetDeviceByDeviceId       = `SELECT device_id,mac_address,device_type,chip_id,created_at,id,device_type FROM mst_device WHERE device_id = ?`
	mysqlQueryGetLastTrackingByInterval = `SELECT status,id FROM trx_gps_tracking WHERE device_id = ? AND modified_at >= DATE_SUB(NOW(), INTERVAL ? SECOND)`
	mysqlQueryGetCounter                = `SELECT COUNT(id) AS recorded_tracking FROM trx_gps_tracking WHERE status="STOP"`
	mysqlQueryGetDeviceLogs             = `SELECT device_id,status,reason,recorded_at FROM trx_device_log WHERE DATE(recorded_at) = DATE(?) ORDER BY recorded_at DESC LIMIT ?,?`
	mysqlQueryGetDeviceMonitoringById   = `SELECT 
		mt.id,
		mt.device_id,
		IFNULL(
			(SELECT tdl.status FROM trx_device_log tdl WHERE tdl.device_id = mt.device_id ORDER BY tdl.recorded_at DESC LIMIT 1 ),"NO DATA"
		) AS device_log_status,
		IFNULL(
			(SELECT tdl.reason FROM trx_device_log tdl WHERE tdl.device_id = mt.device_id ORDER BY tdl.recorded_at DESC LIMIT 1 ),"NO DATA"
		) AS log_reason,
		IFNULL(
			(SELECT tdl.signal_strength FROM trx_device_log tdl WHERE tdl.device_id = mt.device_id ORDER BY tdl.recorded_at DESC LIMIT 1 ),0
		) AS log_signal_strength,
		IFNULL(
			(SELECT tdl.recorded_at FROM trx_device_log tdl WHERE tdl.device_id = mt.device_id ORDER BY tdl.recorded_at DESC LIMIT 1 ),NOW()
		) AS log_record_at,
		IFNULL(
			(SELECT tgt.speed FROM trx_gps_tracking tgt WHERE tgt.device_id = mt.device_id ORDER BY tgt.modified_at DESC LIMIT 1),0
		) AS gps_speed
	FROM mst_device mt WHERE mt.device_id = ?`
)
