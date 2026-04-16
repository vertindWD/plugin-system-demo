package manager

import (
	"context"
	"errors"
	"fmt"
	"plugin"
	"sync"

	"plugin-system/core"
)

type PluginManager struct {
	plugins map[string]core.Plugin
	mu      sync.RWMutex
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make(map[string]core.Plugin),
	}
}

func (pm *PluginManager) LoadPlugin(path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return err
	}

	sym, err := p.Lookup("Plugin")
	if err != nil {
		return err
	}

	pl, ok := sym.(core.Plugin)
	if !ok {
		return errors.New("invalid plugin type")
	}

	pm.mu.Lock()
	pm.plugins[pl.Name()] = pl
	pm.mu.Unlock()

	return nil
}

func (pm *PluginManager) Execute(ctx context.Context, name string, input map[string]interface{}) (res map[string]interface{}, err error) {
	pm.mu.RLock()
	pl, exists := pm.plugins[name]
	pm.mu.RUnlock()

	if !exists {
		return nil, errors.New("plugin not found")
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("plugin panic recovered: %v", r)
		}
	}()

	return pl.Run(ctx, input)
}
