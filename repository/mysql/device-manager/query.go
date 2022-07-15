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
	// Read....
	mysqlQueryGetDeviceByDeviceId       = `SELECT device_id,mac_address,device_type,chip_id,created_at FROM mst_device WHERE device_id = ?`
	mysqlQueryGetLastTrackingByInterval = `SELECT status,id FROM trx_gps_tracking WHERE device_id = ? AND modified_at >= DATE_SUB(NOW(), INTERVAL ? SECOND)`
)
