package config

type Redis struct {
	DB           int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr         string    `mapstructure:"addr" json:"addr" yaml:"addr"`
	PassWord     string `mapstructure:"password" json:"password" yaml:"password"`
}
