package infra

import (
	config "github.com/aditya37/get-env"
)

// get config in param infra...
// ex: infra mysql client..

// declare or set param for connect to mqtt....
func GetMQTTConfig() MQTTConf {
	return MQTTConf{
		Host:     config.GetString("MQTT_HOST", "127.0.0.1"),
		Port:     int64(config.GetInt("MQTT_PORT", 1883)),
		Username: config.GetString("MQTT_USERNAME", ""),
		Password: config.GetString("MQTT_PASSWORD", ""),
		ClientId: "tracking1",
	}
}

// declare mysql config instance...
func GetMysqlConfig() MysqlConfigParam {
	return MysqlConfigParam{
		Host:     config.GetString("DB_HOST", "127.0.0.1"),
		Port:     config.GetInt("DB_PORT", 3306),
		Name:     config.GetString("DB_NAME", "db_geofencing"),
		User:     config.GetString("DB_USER", "root"),
		Password: config.GetString("DB_PASSWORD", "admin"),
	}
}

// declare redis config....
func GetRedisConfig() RedisConfigParam {
	return RedisConfigParam{
		Port:     config.GetInt("REDIS_PORT", 6379),
		Host:     config.GetString("REDIS_HOST", "127.0.0.1"),
		Password: config.GetString("REDIS_PASSWORD", ""),
	}
}

func GetFirebaseConfig() FirebaseConfig {
	return FirebaseConfig{
		StorageBucketName: config.GetString("FIREBAE_BUCKET_PATH", "gs://device-service-1029d.appspot.com/"),
		ProjectId:         config.GetString("FIREBASE_PROJECT_ID", "device-service-1029d"),
		PathCredFile:      config.GetString("FIREBASE_CRED_FILE", "sa.fbs.device.service.json"),
	}
}
