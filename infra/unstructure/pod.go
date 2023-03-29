package unstructure

import (
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"github.com/spyzhov/ajson"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"log"
)

var _ familiar.Familiar = &Pod{}

type Pod struct {
	uns *unstructured.Unstructured
}

func (p *Pod) NeedMore() []*familiar.Id {
	var ret []*familiar.Id
	serviceAccountName, exists, err := unstructured.NestedString(p.uns.Object, "spec", "serviceAccountName")
	if !exists || err != nil {
		log.Fatalf("%+v %+v", exists, err)
	}
	ret = append(ret, &familiar.Id{
		GroupVersionKind: familiar.GroupVersionKind{
			Version: "v1",
			Kind:    "ServiceAccount",
		},
		Namespace: p.uns.GetNamespace(),
		Name:      serviceAccountName,
	})

	bytes, err := p.uns.MarshalJSON()
	if err != nil {
		panic(err)
	}
	root, err := ajson.Unmarshal(bytes)
	if err != nil {
		panic(err)
	}
	configmaps, err := root.JSONPath("$.spec.containers[*].env[*].valueFrom.configMapKeyRef.name")
	if err != nil {
		panic(err)
	}
	for _, configmap := range configmaps {
		value, err := configmap.GetString()
		if err != nil {
			panic(err)
		}
		ret = append(ret, &familiar.Id{
			GroupVersionKind: familiar.GroupVersionKind{
				Version: "v1",
				Kind:    "ConfigMap",
			},
			Namespace: p.uns.GetNamespace(),
			Name:      value,
		})
	}
	secrets, err := root.JSONPath("$.spec.containers[*].env[*].valueFrom.secretKeyRef.name")
	if err != nil {
		panic(err)
	}
	for _, secret := range secrets {
		value, err := secret.GetString()
		if err != nil {
			panic(err)
		}
		ret = append(ret, &familiar.Id{
			GroupVersionKind: familiar.GroupVersionKind{
				Version: "v1",
				Kind:    "Secret",
			},
			Namespace: p.uns.GetNamespace(),
			Name:      value,
		})
	}
	return ret
}
