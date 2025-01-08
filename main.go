package main

import (
	"containers/api"
	"net/http"
)

func main() {
	// 注册prometheus metrics接口
	api.RegisterMetrics()
	// 实现自定义Webhook,并注册到路由中
	http.HandleFunc("/ocean-validate", api.OceanHook)
	http.HandleFunc("/healthz", api.OceanHealthCheck)
	http.HandleFunc("/query", api.OceanQuery)
	http.Handle("/metrics", api.OceanGetMetric())
	// 创建HTTP服务器（启用TLS传输）
	api.StartServer()
}
