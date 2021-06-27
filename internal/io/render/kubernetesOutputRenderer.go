package render

import (
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
		pr.log.Panicf("Error while trying to decode kubectl show pods %v+", err)
	}

	pods := make([]*v1.Pod, 0, len(object.Items))
	for _, v := range object.Items {
		pod := &v1.Pod{}
		runtime.DecodeInto(decoder, v.Raw, pod)
		pods = append(pods, pod)
	}

	// for _, p := range pods {
	// 	pr.log.Printf("%s - %s", p.Namespace, p.Name)
	// }

	// if err := ui.Init(); err != nil {
	// 	log.Fatalf("failed to initialize termui: %v", err)
	// }
	// defer ui.Close()
	// table1 := widgets.NewTable()
	// table1.Rows = [][]string{
	// 	[]string{"header1", "header2", "header3"},
	// 	[]string{"你好吗", "Go-lang is so cool", "Im working on Ruby"},
	// 	[]string{"2016", "10", "11"},
	// }
	// table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	// table1.SetRect(0, 0, 60, 10)

	// ui.Render(table1)
	// uiEvents := ui.PollEvents()
	// for {
	// 	e := <-uiEvents
	// 	switch e.ID {
	// 	case "q", "<C-c>":
	// 		return
	// 	}
	// }
}

// RenderLogs renders the passed in output to the console via simple stdout call
func (pr KubernetesOutputRenderer) RenderLogs(output []byte) {
	pr.log.Print(string(output))
}
