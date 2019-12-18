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
	Validation struct {
		MaxPicsNumber       int `yaml:"max_pics_num"`
		MaxDesciptionLength int `yaml:"max_description_len"`
		MaxNameLength       int `yaml:"max_name_len"`
	} `yaml:"validation"`
	AdvertList struct {
		MaxAdvertsOnPage int `yaml:"max_adverts_on_page"`
	} `yaml:"advert_list_page"`
}
