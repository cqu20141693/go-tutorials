package resource

type MongoConfig struct {
	Address        string `json:"address"`
	Database       string `json:"database"`
	User           string `json:"user"`
	Password       string `json:"password"`
	MaxConnections int    `json:"maxConnections" yaml:"maxConnections"`
	Timeout        int    `json:"timeout"`
	Mechanism      string `json:"mechanism"`
	AuthSource     string `json:"authSource" yaml:"authSource"`
	Debug          bool   `json:"debug"`
}

type Config struct {
	Port  uint16      `json:"port"`
	Mongo MongoConfig `json:"mongo"`
}
