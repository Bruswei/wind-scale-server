package config

type AppConfig struct {
	CSVFilePath string `env:"CSV_FILE_PATH"`
	ListenPort  string `env:"LISTEN_PORT"`
}

func GetConfig() AppConfig {
	return AppConfig{
		CSVFilePath: "internal/csvdata/wind-speed.csv",
		ListenPort:  "8080",
	}
}
