package configuration

type Datasource struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
