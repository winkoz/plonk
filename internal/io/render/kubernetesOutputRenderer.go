package render

import (
	"fmt"
	"log"
	"os"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
)

// PlainOutputRenderer renders bytes output as simple strings to console
type KubernetesOutputRenderer struct {
	log *log.Logger
}

// NewPlainOutputRenderer creates a plain output renderer that renders buffers into the console via stdout
func NewKubernetesOutputRenderer() KubernetesOutputRenderer {
	return KubernetesOutputRenderer{
		log: log.New(os.Stdout, "", 0),
	}
}

// RenderComponents renders the passed in output to the console via simple stdout call
func (pr KubernetesOutputRenderer) RenderComponents(output []byte) {
	decoder := serializer.NewCodecFactory(scheme.Scheme).UniversalDecoder()
	object := &v1.List{}
	err := runtime.DecodeInto(decoder, output, object)
	if err != nil {
		pr.log.Panicf("ERROR ESTO ES UN ERROR %v+", err)
	}

	pods := make([]*v1.Pod, len(object.Items))

	for _, v := range object.Items {
		pod := &v1.Pod{}
		runtime.DecodeInto(decoder, v.Raw, pod)
		pods = append(pods, pod.DeepCopy())
	}

	for _, p := range pods {
		pr.log.Print(fmt.Sprintf("%s - %s", (*p).Namespace, (*p).Name))
	}

}

// RenderLogs renders the passed in output to the console via simple stdout call
func (pr KubernetesOutputRenderer) RenderLogs(output []byte) {
	pr.log.Print(string(output))
}
