package config

type ConfStruct struct {
	DebugMode bool `yaml:"debug_mode"`
	Server    struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		DBserver struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"dbserver"`
		Credentials struct {
			User string `yaml:"user"`
			Pass string `yaml:"pass"`
		} `yaml:"credentials"`
	} `yaml:"database"`
}
