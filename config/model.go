package config

type ConfStruct struct {
	DebugMode bool `yaml:"debug_mode"`

	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		DBserver struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"dbserver"`
		Credentials struct {
			User   string `yaml:"user"`
			Pass   string `yaml:"pass"`
			DbName string `yaml:"dbname"`
			Schema string `yaml:"schema"`
		} `yaml:"credentials"`
		Maintainance struct {
			AnalyzeInterval int `yaml:"analyze_period"`
			VacuumInterval  int `yaml:"vacuum_period"`
		} `yaml:"maintainance"`
	} `yaml:"database"`
	Cache struct {
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		Expiration int    `yaml:"expiration"`
	}
	Validation struct {
		MaxPicsNumber       int `yaml:"max_pics_num"`
		MaxDesciptionLength int `yaml:"max_description_len"`
		MaxNameLength       int `yaml:"max_name_len"`
	} `yaml:"validation"`

	AdvertList struct {
		MaxAdvertsOnPage int `yaml:"max_adverts_on_page"`
	} `yaml:"advert_list_page"`
}
