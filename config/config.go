package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var values *Values

type Values struct {
	// database
	DbHost     string
	DbPort     uint
	DbUser     string
	DbName     string
	DbPassword string

	PrivateKey  string
	Port        int
	GinMode     string
	RPC         string
	Contract    string
	MaxGas      string
	CreateGas   string
	VipMaxGas   string
	VipContract string
}

func InitValues() error {
	viper.SetDefault("port", 8888)
	viper.SetDefault("gin_mode", gin.ReleaseMode)
	viper.SetDefault("CREATE_GAS", "5000000000000000000")
	viper.SetDefault("MAX_GAS", "2000000000000000000")
	viper.SetDefault("VIP_MAX_GAS", "10000000000000000000")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	_ = viper.BindEnv("DB_HOST")
	_ = viper.BindEnv("DB_PORT")
	_ = viper.BindEnv("DB_USER")
	_ = viper.BindEnv("DB_NAME")
	_ = viper.BindEnv("DB_PASSWORD")
	_ = viper.BindEnv("PORT")
	_ = viper.BindEnv("GIN_MODE")
	_ = viper.BindEnv("PRIVATE_KEY")
	_ = viper.BindEnv("RPC")
	_ = viper.BindEnv("CONTRACT")
	_ = viper.BindEnv("CREATE_GAS")
	_ = viper.BindEnv("MAX_GAS")
	_ = viper.BindEnv("VIP_MAX_GAS")
	_ = viper.BindEnv("VIP_CONTRACT")

	values = &Values{
		DbHost:      viper.GetString("DB_HOST"),
		DbPort:      viper.GetUint("DB_PORT"),
		DbUser:      viper.GetString("DB_USER"),
		DbName:      viper.GetString("DB_NAME"),
		DbPassword:  viper.GetString("DB_PASSWORD"),
		PrivateKey:  viper.GetString("PRIVATE_KEY"),
		Port:        viper.GetInt("PORT"),
		GinMode:     viper.GetString("GIN_MODE"),
		RPC:         viper.GetString("RPC"),
		Contract:    viper.GetString("CONTRACT"),
		CreateGas:   viper.GetString("CREATE_GAS"),
		MaxGas:      viper.GetString("MAX_GAS"),
		VipMaxGas:   viper.GetString("VIP_MAX_GAS"),
		VipContract: viper.GetString("VIP_CONTRACT"),
	}
	return nil
}

func Config() *Values {
	if values == nil {
		log.Fatal("config not initial")
	}
	return values
}
