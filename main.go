package main

import (
	"containers/api"
	"encoding/json"
	"fmt"
	"log"
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
	http.HandleFunc("/call", api.LoggerForTestAPI(
		func(w http.ResponseWriter, r *http.Request) {
			// 这是call接口的原始代码逻辑
			if r.Method != "GET" {
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte("Only GET method is allowed"))
				return
			}
			var score int
			v := map[string]any{
				"code": 0,
				"msg":  fmt.Sprintf("你的成绩是:%d", score),
			}
			respData, err := json.Marshal(v)
			if err != nil {
				log.Println("序列化数据发生错误,", err)
				return
			}
			w.Write(respData)
		}))
	// 创建HTTP服务器（启用TLS传输）
	api.StartServer()
}
