package api

import (
	"fmt"

	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	runtimeScheme = runtime.NewScheme()
	codecs        = serializer.NewCodecFactory(runtimeScheme)
	deserializer  = codecs.UniversalDecoder()
)

func init() {
	err := v1.AddToScheme(runtimeScheme)
	if err != nil {
		fmt.Println("add to scheme error.,", err)
	}
}

func admitPod(ar v1beta1.AdmissionRequest) *v1beta1.AdmissionResponse {
	admissionResp := &v1beta1.AdmissionResponse{
		Allowed: true,
		UID:     ar.UID,
	}
	podResource := v1.Resource("pods")
	if ar.RequestResource.Resource != podResource.Resource {
		fmt.Println("pod 不相等")
		return admissionResp
	}
	// 接下来是Pod资源类型，对其先反序列化拿到Pod相关资源数据
	raw := ar.Object.Raw
	pod := v1.Pod{}
	_, _, err := deserializer.Decode(raw, nil, &pod)
	if err != nil {
		fmt.Println("反序列化成Go中数据出错", err)
		admissionResp.Allowed = false
		admissionResp.Result = &metav1.Status{
			Message: err.Error(),
		}
		return admissionResp
	}

	// 自定义Hook操作验证
	for _, c := range pod.Spec.Containers {
		// 判断pod request是否为0
		if c.Resources.Requests.Cpu().IsZero() {

			admissionResp.Allowed = false
			admissionResp.Result = &metav1.Status{
				Message: "Must have CPU requests.",
			}
			return admissionResp
		}
	}
	return admissionResp

}
