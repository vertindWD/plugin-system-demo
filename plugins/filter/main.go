package main

import (
	"context"
	"errors"
	"time"
)

type FilterPlugin struct{}

func (f *FilterPlugin) Name() string    { return "TextFilter" }
func (f *FilterPlugin) Version() string { return "1.1.0" }

func (f *FilterPlugin) Run(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	// 模拟崩溃
	if data["panic"] == true {
		panic("simulated panic")
	}

	// 模拟耗时操作，响应 Context 超时信号
	if data["delay"] == true {
		select {
		case <-time.After(3 * time.Second):
		case <-ctx.Done():
			return nil, errors.New("plugin timeout")
		}
	}

	data["status"] = "ok"
	return data, nil
}

var Plugin FilterPlugin
