package gruutbot

import (
	"context"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/spf13/viper"
)

type GruutBot struct {
	viper         *viper.Viper
	client        *disgord.Client
	log           Logger
	token         string
	prefix        string
	pluginManager *PluginManager
}

func New(configs ...Config) *GruutBot {
	var c Config
	if len(configs) < 1 {
		c = Config{}
	} else {
		c = configs[0]
	}

	v := gviper
	log := fetchLogger(c)

	pluginManager := GetPluginManager(fetchPluginsPath(c), log)

	err := pluginManager.LoadPlugins()
	if err != nil {
		log.Panic("Error loading plugins:", err)
	}

	return &GruutBot{
		viper:         v,
		log:           log,
		token:         fetchToken(c),
		prefix:        fetchPrefix(c),
		pluginManager: pluginManager,
	}
}

func (g *GruutBot) Start() {
	g.log.Infof("Starting bot. Using %s as prefix", g.prefix)

	g.client = disgord.New(disgord.Config{
		BotToken:     g.token,
		Logger:       g.log,
		DisableCache: true,
	})

	defer func() {
		_ = g.client.StayConnectedUntilInterrupted(context.Background())
	}()

	setupEvents(g)
}

func fetchLogger(c Config) (l Logger) {
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

func fetchToken(c Config) (token string) {
	const tokenKey = "token"

	token = strings.TrimSpace(c.Token)

	if gviper.IsSet(tokenKey) {
		token = gviper.GetString(tokenKey)
	}

	return
}

func fetchPrefix(c Config) (prefix string) {
	const prefixKey = "prefix"

	prefix = strings.TrimSpace(c.Prefix)

	if gviper.IsSet(prefixKey) {
		prefix = gviper.GetString(prefixKey)
	}

	return
}

func fetchPluginsPath(c Config) (path string) {
	const pathKey = "plugins"

	path = strings.TrimSpace(c.Plugins)

	if gviper.IsSet(pathKey) {
		path = gviper.GetString(pathKey)
	}

	if len(path) < 1 {
		path = "plugins"
	}

	return
}
