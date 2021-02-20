package util

import "github.com/spf13/viper"

// DatabaseConfig is ...
type DatabaseConfig struct {
	DatabaseEngine   string `mapstructure:"DB_ENGINE"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     int    `mapstructure:"DB_PORT"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
	DatabaseName     string `mapstructure:"DB_DATBASE_NAME"`
}

// LoadDatabaseConfig is ...
func LoadDatabaseConfig(path string) (databaseConfig DatabaseConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&databaseConfig)
	return
}
