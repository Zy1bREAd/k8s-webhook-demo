package main

import (
	"containers/api"
	"net/http"
)

func main() {
	// 实现自定义Webhook,并注册到路由中
	http.HandleFunc("/ocean-validate", api.OceanHook)
	// 创建HTTP服务器（启用TLS传输）
	api.StartServer()
}
