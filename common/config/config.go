package config

type Config struct {
	Addr        Addr        `yaml:"addr"`
	Rpc         Rpc         `yaml:"rpc"`
	DbConfig    DbConfig    `yaml:"dbConfig"`
	RedisConfig RedisConfig `yaml:"redisConfig"`
	JwtConfig   JwtConfig   `yaml:"jwtConfig"`
	EmailConfig EmailConfig `yaml:"emailConfig"`
}

// Addr 服务地址 运行地址
type Addr struct {
	ApiAddr    string `yaml:"apiAddr"`
	PublicAddr string `yaml:"publicAddr"`
}

// Rpc 服务发现 调用的地址
type Rpc struct {
	PublicRpc string `yaml:"publicRpc"`
}

type DbConfig struct {
	DbType   string `yaml:"dbType"`
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type JwtConfig struct {
	SignKey string `yaml:"signKey"`
	Expire  int64  `yaml:"expire"`
}

type EmailConfig struct {
	SMTPAddr string `yaml:"smtpAddr"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
	Host     string `yaml:"host"`
}
