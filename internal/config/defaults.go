package config

var supportedAlgorithms = [3]string{
	"round-robin",
	"least-connection",
	"weighted-response-time",
}

var algorithmsAliases = [3]string{
	"rr",
	"lc",
	"wrt",
}

func (cfg *Config) ApplyDefaults() *Config {
	if cfg.Version != 1 {
		cfg.Version = 1
	}

	if cfg.Port == 0 {
		cfg.Port = 4000
	}

	for i, algorithm := range supportedAlgorithms {
		if cfg.Algorithm == algorithm {
			break
		}

		if cfg.Algorithm == algorithmsAliases[i] {
			cfg.Algorithm = algorithm
			break
		}

		if i == len(supportedAlgorithms)-1 {
			cfg.Algorithm = "round-robin"
		}
	}

	return cfg
}
