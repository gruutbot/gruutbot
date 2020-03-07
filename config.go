package gruutbot

import "github.com/spf13/viper"

var gviper *viper.Viper

type Config struct {
	Logger   *Logger
	LogLevel *LogLevel
}

func init() {
	gviper = viper.New()
	setupEnv()
}

func setupEnv() {
	gviper.AutomaticEnv()
	gviper.SetEnvPrefix("gruutbot")
	_ = gviper.BindEnv(logLevelKey)
	gviper.AllowEmptyEnv(true)
}
