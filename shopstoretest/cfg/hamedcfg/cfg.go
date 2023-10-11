package hamedcfg

import "shopstoretest/cfg"

const (
	ServerHost       = "localhost:"
	ServerPort       = "5555"
	DatabaseHost     = "localhost"
	DatabasePort     = "3308"
	DatabaseName     = "shopstore_db"
	DatabaseUser     = "shopstore"
	DatabasePassword = "shopstoret0lk2o20"
	DatabaseProtocol = "tcp"
)

var Database = cfg.DataBaseCfg{
	UserName: DatabaseUser,
	Password: DatabasePassword,
	Host:     DatabaseHost,
	Port:     DatabasePort,
	Protocol: DatabaseProtocol,
	Database: DatabaseName,
}
