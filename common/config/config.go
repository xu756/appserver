package config

type Config struct {
	Addr        Addr        `yaml:"addr"`
	Rpc         Rpc         `yaml:"rpc"`
	DbConfig    DbConfig    `yaml:"dbConfig"`
	RedisConfig RedisConfig `yaml:"redisConfig"`
	JwtConfig   JwtConfig   `yaml:"jwtConfig"`
	EmailConfig EmailConfig `yaml:"emailConfig"`
	MqUrl       string      `yaml:"mqUrl"`
	MinioConfig MinioConfig `yaml:"minioConfig"`
}

// Addr 服务地址 运行地址
type Addr struct {
	AdminApiAddr string `yaml:"adminApiAddr"`
	WxappApiAddr string `yaml:"wxappApiAddr"`
	UserRpcAddr  string `yaml:"userRpcAddr"`
}

// Rpc 服务发现 调用的地址
type Rpc struct {
	UserRpcAddr string `yaml:"userRpcAddr"`
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

// MinioConfig  配置
type MinioConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	UseSSL          bool   `yaml:"useSSL"`
	BucketName      string `yaml:"bucketName"`
}
