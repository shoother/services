package g

type RedisConfig struct {
	Addr     string `json:"addr"`
	Idle     int    `json:"idle"`
	Password string `json:"password"`
}

type MysqlConfig struct {
	Addr string `json:"addr"`
	Idle int    `json:"idle"`
	Max  int    `json:"max"`
}

type GlobalConfig struct {
	Database      *MysqlConfig `json:"database"`
}

var (
	RedisKVConfig RedisConfig
	RedisDBConfig RedisConfig
	RedisQueueConfig RedisConfig
	Root string
)