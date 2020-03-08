package gruutbot

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"sync"

	"github.com/sirupsen/logrus"
)

var pluginManager *PluginManager
var pmMux sync.Mutex

type PluginManager struct {
	pluginsPath string
	plugins     map[string]Plugin
	Log         Logger
	commands    map[string]func(CommandMessage, Logger) error
}

func GetPluginManager(pluginsPath string, log Logger) *PluginManager {
	pmMux.Lock()
	defer pmMux.Unlock()

	if pluginManager == nil {
		pluginManager = &PluginManager{
			pluginsPath: pluginsPath,
			plugins:     make(map[string]Plugin),
			Log:         log,
			commands:    make(map[string]func(CommandMessage, Logger) error),
		}
	}

	return pluginManager
}

func (pm *PluginManager) LoadPlugins() (err error) {
	var files []string
	files, err = findPlugins(pm.pluginsPath, "*.so")

	if err != nil {
		return
	}

	for _, f := range files {
		p, err := plugin.Open(f)
		if err != nil {
			pm.Log.Errorf("error loading plugin %s: %s", f, err)
			continue
		}

		symPlugin, err := p.Lookup("Plugin")
		if err != nil {
			pm.Log.Errorf("error registering plugin %s: %s", f, err)
			continue
		}

		var plug Plugin

		plug, ok := symPlugin.(Plugin)
		if !ok {
			logrus.Errorf("unexpected type from plugin %s", f)
			continue
		}

		logrus.Infof("Registering plugin %s version %s", plug.GetName(), plug.GetVersion())
		plug.Register(pm)
	}

	pm.Log.Info("Finished loading plugins")
	pm.Log.Debug("Loaded plugins", pm.plugins)
	pm.Log.Debug("Registered commands", pm.commands)

	return
}

func (pm *PluginManager) RegisterPlugin(plugin Plugin) {
	pm.plugins[plugin.GetName()] = plugin
}

func findPlugins(root, pattern string) (matches []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) (er error) {
		if err != nil {
			er = err
			return
		}
		if info.IsDir() {
			return
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			er = err
		} else if matched {
			matches = append(matches, path)
		}
		return
	})
	if err != nil {
		return nil, err
	}

	return
}

func (pm *PluginManager) RegisterCommand(command string, f func(CommandMessage, Logger) error) error {
	if pm.commands[command] != nil {
		return fmt.Errorf("command already registered")
	}

	pm.commands[command] = f

	return nil
}
