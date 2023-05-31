package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/yaml"
	clientscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"
)

func main() {
	data, err := ioutil.ReadFile("node-templates-non-list.yaml")
	if err != nil {
		panic(err)
	}

	scheme := runtime.NewScheme()
	clientscheme.AddToScheme(scheme)

	decoder := serializer.NewCodecFactory(scheme).UniversalDeserializer()

	multiDocReader := yaml.NewYAMLReader(bufio.NewReader(bytes.NewReader(data)))
	nodeTemplates := []*v1.Node{}

	objs := []runtime.Object{}

	for {
		buf, err := multiDocReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		obj, _, err := decoder.Decode(buf, nil, nil)
		if err != nil {
			panic(err)
		}

		objs = append(objs, obj)
	}

	// This doesn't work
	// nodeList := &v1.NodeList{}
	// _, _, err = decoder.Decode(data, &schema.GroupVersionKind{Group: "", Version: "v1", Kind: "NodeList"}, nodeList)
	// if err != nil {
	// 	panic(err)
	// }

	if len(objs) > 1 {
		for _, obj := range objs {
			if node, ok := obj.(*v1.Node); ok {
				// fmt.Println("node individual", node)
				nodeTemplates = append(nodeTemplates, node)
			}
		}

	} else if nodelist, ok := objs[0].(*v1.List); ok {
		for _, item := range nodelist.Items {

			o, _, err := decoder.Decode(item.Raw, nil, nil)
			if err != nil {
				klog.Fatal(err)
			}

			if node, ok := o.(*v1.Node); ok {
				nodeTemplates = append(nodeTemplates, node)
			}
		}
	}

	for _, nodeTemplate := range nodeTemplates {
		fmt.Println("node", nodeTemplate.Name)
	}

}
