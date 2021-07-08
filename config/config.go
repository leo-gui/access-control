package config

type Server struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Const  Const  `mapstructure:"const" json:"const" yaml:"const"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
}

