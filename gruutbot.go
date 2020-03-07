package gruutbot

import (
	"github.com/andersfylling/disgord"
	"github.com/spf13/viper"
)

type GruutBot struct {
	viper  *viper.Viper
	client *disgord.Client
	log    Logger
}

func New(configs ...Config) *GruutBot {
	var c Config
	if len(configs) < 1 {
		c = Config{}
	} else {
		c = configs[0]
	}

	v := gviper

	return &GruutBot{viper: v, log: getLogger(c)}
}

func (g *GruutBot) Start() {
	g.log.Info("Starting bot")
}

func getLogger(c Config) (l Logger) {
	if c.Logger == nil {
		logLevel := InfoLevel
		if c.LogLevel != nil {
			logLevel = *c.LogLevel
		}

		l = logrusLogger(logLevel)

		return
	}

	l = *c.Logger

	return
}
