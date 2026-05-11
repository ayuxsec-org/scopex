package config

func New() Config {
	return Config{
		Hackerone: HackerOneCfg{
			RateLimitPerMin: 600, // https://api.hackerone.com/getting-started/#rate-limits
		},
	}
}

type Config struct {
	Hackerone HackerOneCfg `yaml:"hackerone"`
}

type HackerOneCfg struct {
	Creds           HackerOneApiCreds `yaml:"creds"`
	RateLimitPerMin int               `yaml:"rpm"`
}

type HackerOneApiCreds struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}
