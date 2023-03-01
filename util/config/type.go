package config

type MongoConfig struct {
	Host       string `yaml:"Host"`
	Port       string `yaml:"Port"`
	DB         string `yaml:"DB"`
	User       string `yaml:"User"`
	Password   string `yaml:"Password"`
	Collection string `yaml:"Collection"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"Brokers"`
	Topic   string   `yaml:"Topic"`
	GroupID string   `yaml:"GroupID"`
}

type HttpConfig struct {
	Port string `yaml:"Port"`
}

type ServerConfig struct {
	Http HttpConfig `yaml:"Http"`
}

type Config struct {
	Version string       `yaml:"Version"`
	Server  ServerConfig `yaml:"Server"`
	Mongo   MongoConfig  `yaml:"Mongo"`
	Kafka   KafkaConfig  `yaml:"Kafka"`
}
