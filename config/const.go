package config

type Const struct {
	Capacity                      int `mapstructure:"capacity" json:"capacity" yaml:"capacity"`
	ProdTokenIntervalMs           int `mapstructure:"prodTokenIntervalMs" json:"prodTokenIntervalMs" yaml:"prodTokenIntervalMs"`
	ProdTokenNumEveryInterval     int `mapstructure:"prodTokenNumEveryInterval" json:"prodTokenNumEveryInterval" yaml:"prodTokenNumEveryInterval"`
}