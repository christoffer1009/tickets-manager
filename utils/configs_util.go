package utils

import "github.com/spf13/viper"

func GetSecretkey() ([]byte, error) {
	viper.SetConfigName("app_config")
	viper.AddConfigPath("./../../configs/")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return []byte(viper.GetString("app.secret")), nil
}
