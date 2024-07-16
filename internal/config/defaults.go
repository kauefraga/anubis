package config

func (cfg *Config) ApplyDefaults() *Config {
	if cfg.Version != 1 {
		cfg.Version = 1
	}

	if cfg.Port == 0 {
		cfg.Port = 4000
	}

	if cfg.Algorithm != "round-robin" && cfg.Algorithm != "least-connection" && cfg.Algorithm != "weighted-response-time" {
		cfg.Algorithm = "round-robin"
	}

	return cfg
}
