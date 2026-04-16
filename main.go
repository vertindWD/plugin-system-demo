package main

import (
	"context"
	"fmt"
	"time"

	"plugin-system/manager"
)

func main() {
	pm := manager.NewPluginManager()

	if err := pm.LoadPlugin("./plugins/filter/filter.so"); err != nil {
		fmt.Println("load err:", err)
		return
	}

	// 1. 测试正常执行
	out1, err1 := pm.Execute(context.Background(), "TextFilter", map[string]interface{}{"text": "test"})
	fmt.Printf("1. Normal  -> out: %v, err: %v\n", out1, err1)

	// 2. 测试崩溃隔离
	out2, err2 := pm.Execute(context.Background(), "TextFilter", map[string]interface{}{"panic": true})
	fmt.Printf("2. Panic   -> out: %v, err: %v\n", out2, err2)

	// 3. 测试超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	out3, err3 := pm.Execute(ctx, "TextFilter", map[string]interface{}{"delay": true})
	fmt.Printf("3. Timeout -> out: %v, err: %v\n", out3, err3)
}
