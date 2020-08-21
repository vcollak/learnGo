package config

//configuration in simple variables
var (
	Path     = "myPath"
	Version  = "1.1.1"
	Hostname = "localhost"
	Port     = 3000
)

//Config degines configuration of some kind
type Config struct {
	Path     string
	Version  string
	Hostname string
	Port     int
}

//New creates a new config using defined defaults
func New() *Config {

	return &Config{
		Path:     Path,
		Version:  Version,
		Hostname: Hostname,
		Port:     Port,
	}

}
