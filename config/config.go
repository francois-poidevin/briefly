package config

// Configuration contains conectivity settings
type Configuration struct {
	Log struct {
		Level         string `toml:"level" default:"warning" comment:"Log level: trace, debug, info, warning, error, panic, and fatal"`
		JSONFormatter bool   `toml:"jsonformatter" default:"false" comment:"Allow to display logs in Json format if true"`
	} `toml:"Log" comment:"###############################\n Logs Settings \n##############################"`

	Briefly struct {
		GRPC struct {
			ListenPort int `toml:"listenPort" default:"5556" comment:"On which port REST HTTP service will listen"`
		} `toml:"GRPC" comment:"###############################\n GRPC API settings \n##############################"`
	} `toml:"Briefly" comment:"###############################\n Briefly Settings \n##############################"`
}
