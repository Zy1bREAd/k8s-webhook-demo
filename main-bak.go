package main

import (
	"encoding/json"
	"log"
	"net/http"

	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	runtimeScheme = runtime.NewScheme()
	codecs        = serializer.NewCodecFactory(runtimeScheme)
	deserializer  = codecs.UniversalDecoder()
)

func admitPod(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	podResource := v1.Resource("pods")
	if ar.Request.Resource != podResource {
		return &v1beta1.AdmissionResponse{
			Allowed: true,
		}
	}
	raw := ar.Request.Object.Raw
	pod := v1.Pod{}
	_, _, err := deserializer.Decode(raw, nil, &pod)
	if err != nil {
		return &v1beta1.AdmissionResponse{
			Allowed: false,
			Result:  &v1beta1.Status{Message: err.Error()},
		}
	}
	// 简单验证：检查容器是否有CPU请求
	for _, container := range pod.Spec.Containers {
		if container.Resources.Requests.Cpu().IsZero() {
			return &v1beta1.AdmissionResponse{
				Allowed: false,
				Result:  &v1beta1.Status{Message: "Container must have CPU request"},
			}
		}
	}
	return &v1beta1.AdmissionResponse{
		Allowed: true,
	}
}

func serve(w http.ResponseWriter, r *http.Request) {
	var ar v1beta1.AdmissionReview
	if err := json.NewDecoder(r.Body).Decode(&ar); err != nil {
		log.Println(err)
		resp := v1beta1.AdmissionResponse{
			Allowed: false,
			Result:  &v1beta1.Status{Message: err.Error()},
		}
		writeResponse(w, resp)
		return
	}
	resp := admitPod(ar)
	writeResponse(w, resp)
}

func writeResponse(w http.ResponseWriter, resp *v1beta1.AdmissionResponse) {
	responseBytes, err := json.Marshal(v1beta1.AdmissionReview{Response: resp})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(responseBytes)
}

func init() {
	_ = v1.AddToScheme(runtimeScheme)
}

// func main() {
// 	fmt.Println("Micro Service - Containers")
// 	// gRPC客户端
// 	// grpc.NewClient("localhost:7777")
// 	http.HandleFunc("/validate-ocean", serve)
// 	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))

// }
