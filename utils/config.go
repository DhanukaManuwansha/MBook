package utils

import "github.com/spf13/viper"

//config stores all configurations of the application
//The values are read by viper from a config file or environment varibales

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE_NAME"`
}

type PatientConfig struct {
	PatientDBDriver string `mapstructure:"DB_DRIVER"`
	PatientDBSource string `mapstructure:"PatientDB_SOURCE_NAME"`
}

//loadconfig reads configuration from file ore environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func LoadPatientConfig(path string) (config PatientConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
