package core

import "context"

type Plugin interface {
	Name() string
	Version() string
	Run(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error)
}
