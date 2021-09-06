package config

type Cfg struct {
	Api       *api       `yaml:"api"`
	Redis     *redis     `yaml:"redis"`
	Everdale *everdale `yaml:"everdale"`

}

type redis struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
	IdleTimeout  int    `yaml:"idle_timeout"`
}



type api struct {
	Endpoint string `yaml:"endpoint"`
	RunMode  string `yaml:"run_mode"`
	Env      string `yaml:"env"`
}

type everdale struct {
	Path string `yaml:"path"`
}
