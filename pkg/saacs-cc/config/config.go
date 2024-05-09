package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

type Config struct {
	CCID          string
	Address       string
	ChaincodeName string
	Serialization string
}

func LoadConfig() {
	viper.SetDefault("ChaincodeName", "saacs")
	viper.SetDefault("Serialization", "json")
	viper.MustBindEnv("Serialization", "SERIALIZATION")
	viper.MustBindEnv("CCID", "CHAINCODE_ID")
	viper.MustBindEnv("Address", "CHAINCODE_SERVER_ADDRESS")
	viper.MustBindEnv("ChaincodeName", "CHAINCODE_NAME")
	viper.MustBindEnv("LogLevel", "LOG_LEVEL")
}

func init() {
	LoadConfig()

	config := GetConfig()
	slog.Info("Config loaded",
		"CCID", config.CCID,
		"Address", config.Address,
		"ChaincodeName", config.ChaincodeName,
		"Serialization", config.Serialization,
	)
}

func GetConfig() Config {

	return Config{
		CCID:          viper.GetString("CCID"),
		Address:       viper.GetString("Address"),
		ChaincodeName: viper.GetString("ChaincodeName"),
		Serialization: viper.GetString("Serialization"),
	}
}

func GetCCID() string {
	return viper.GetString("CCID")
}

func GetAddress() string {
	return viper.GetString("Address")
}

func GetChaincodeName() string {
	return viper.GetString("ChaincodeName")
}

func GetSerialization() string {
	return viper.GetString("Serialization")
}

func GetLogLevel() slog.Level {
	switch viper.GetString("LogLevel") {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
