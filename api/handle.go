package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReqData struct {
	ID string `json:"id"`
}

type OceanResp map[string]any

// 封装一个只能返回AdmissionResponse响应对象的函数
func writeWebhookResponse(w http.ResponseWriter, resp *v1beta1.AdmissionResponse) {
	// 将响应对象序列化为json（字节切片）
	responseBytes, err := json.Marshal(v1beta1.AdmissionReview{
		Response: resp,
		// 由于序列化AdmissionReview对象，那么要带上TypeMeta的数据字段，否则准入控制器校验response对象会报错。
		TypeMeta: v1.TypeMeta{
			Kind:       "AdmissionReview",
			APIVersion: "admission.k8s.io/v1",
		},
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(responseBytes)
}

// 普通的response
func writeResponse(w http.ResponseWriter, respBody OceanResp) {
	// 对返回的resp进行序列化
	respBytes, err := json.Marshal(respBody)
	if err != nil {
		log.Println("反序列化resp body出错", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(respBytes)
	if err != nil {
		log.Println("返回响应式数据出错", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// 实际handleFunc函数
func OceanHook(w http.ResponseWriter, r *http.Request) {
	var ar v1beta1.AdmissionReview
	// 对AdmissionRequest请求对象进行反序列化（转换并存储在go结构体中）
	err := json.NewDecoder(r.Body).Decode(&ar)
	if err != nil {
		log.Println("解析apiserver的hook请求失败", err)
		resp := v1beta1.AdmissionResponse{
			Allowed: false,
			Result: &v1.Status{
				Message: err.Error(),
			},
		}
		// 将失败的AdmissionResponse对象进行响应
		writeWebhookResponse(w, &resp)
		return
	}
	// 通过前置判断，进入我们Hook的逻辑中进行校验
	resp := admitPod(*ar.Request)
	writeWebhookResponse(w, resp)
}

// 健康检查
func OceanHealthCheck(w http.ResponseWriter, r *http.Request) {
	// 健康检查
	fmt.Println(r.Body)
	w.Write([]byte("200 OK"))
}

// Post请求根据id获取task（仅prometheus暴露指标使用）
func OceanQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST method is allowed"))
		return
	}
	var reqData ReqData
	respBody := OceanResp{
		"code": 200,
		"data": map[string]string{},
		"msg":  "",
	}
	// 获取请求体并反序列化
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		log.Println("读取请求体中的数据失败", err)
		respBody["data"] = nil
		respBody["code"] = 200
		respBody["msg"] = "读取请求体中的数据失败" + err.Error()
		writeResponse(w, respBody)
		return
	}
	fmt.Println(reqData)
	switch reqData.ID {
	case "1":
		respBody["msg"] = "执行1号任务"
	case "2":
		respBody["msg"] = "执行2号任务"
	default:
		respBody["msg"] = "没有匹配的任务可做..."
	}
	respBody["data"] = map[string]string{
		"name": "oceanwang",
	}
	// 计数器+1 （并发安全的方式）
	myMetrics.QueryCounter.Inc()
	writeResponse(w, respBody)
}

// 暴露给prometheus的metrics接口
func OceanGetMetric() http.Handler {
	return promhttp.Handler()
}

var myMetrics *OceanMetrics

type OceanMetrics struct {
	QueryCounter prometheus.Counter
}

func NewOceanMetrics() *OceanMetrics {
	return &OceanMetrics{
		QueryCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "ocean_query_count",
			Help: "用于收集OceanQuery的请求次数",
		}),
	}
}

func RegisterMetrics() {
	myMetrics = NewOceanMetrics()
	prometheus.MustRegister(myMetrics.QueryCounter)
}
