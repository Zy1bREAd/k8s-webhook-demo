package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func writeResponse(w http.ResponseWriter, resp *v1beta1.AdmissionResponse) {
	responseBytes, err := json.Marshal(v1beta1.AdmissionReview{
		Response: resp,
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

func OceanHook(w http.ResponseWriter, r *http.Request) {
	var ar v1beta1.AdmissionReview
	err := json.NewDecoder(r.Body).Decode(&ar)
	if err != nil {
		log.Println("解析apiserver的hook请求失败", err)
		resp := v1beta1.AdmissionResponse{
			Allowed: false,
			Result: &v1.Status{
				Message: err.Error(),
			},
		}
		writeResponse(w, &resp)
		return
	}
	resp := admitPod(*ar.Request)
	writeResponse(w, resp)

}

func OceanHealthCheck(w http.ResponseWriter, r *http.Request) {
	// 健康检查
	fmt.Println(r.Body)
	w.Write([]byte("200 OK"))
}
