package cfg

type DataBaseCfg struct {
	UserName string
	Password string
	Host     string
	Port     string
	Protocol string
	Database string
}

func NewDatabaseCfg(userName, password, host, port, protocol, database string) DataBaseCfg {
	newDataBaseCfg := DataBaseCfg{
		UserName: userName,
		Password: password,
		Host:     host,
		Port:     port,
		Protocol: protocol,
		Database: database,
	}

	return newDataBaseCfg
}
