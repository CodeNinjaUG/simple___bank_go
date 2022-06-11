package util

import "github.com/spf13/viper"

//stores all environment variables of the application
//the values are read by viper from config variables
type Config struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

//reads configuration from the file or environment variables
func LoadConfig(path string)(config Config,err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil{
       return 
	}
	err = viper.Unmarshal(&config)
	return
}
