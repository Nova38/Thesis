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
	viper.BindEnv("CCID", "CHAINCODE_ID")
	viper.BindEnv("Address", "CHAINCODE_SERVER_ADDRESS")
	viper.BindEnv("ChaincodeName", "CHAINCODE_NAME")
	viper.BindEnv("LogLevel", "LOG_LEVEL")
}

func init() {
	LoadConfig()
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
