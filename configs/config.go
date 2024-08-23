package configs

import "github.com/spf13/viper"

type Config struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	TokenExpiresIn int    `mapstructure:"TOKEN_EXPIRES_IN"`
}

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
