package gruutbot

type Plugin interface {
	Register(manager *PluginManager)
	GetName() string
	GetDescription() string
	GetVersion() string
}
