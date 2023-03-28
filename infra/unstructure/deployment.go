package unstructure

import (
	"fmt"
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"log"
)

var _ familiar.Familiar = &Deployment{}

type Deployment struct {
	uns *unstructured.Unstructured
}

func (d *Deployment) NeedMore() []*familiar.Id {
	selector, exists, err := unstructured.NestedStringMap(d.uns.Object, "spec", "selector", "matchLabels")
	if !exists || err != nil {
		log.Fatalf("%+v %+v", exists, err)
	}
	return []*familiar.Id{
		{
			GroupVersionKind: familiar.GroupVersionKind{
				Group:   "apps",
				Version: "v1",
				Kind:    "ReplicaSet",
			},
			Namespace: d.uns.GetNamespace(),
			Labels:    selector,
		},
	}
}

func (d *Deployment) String() string {
	return fmt.Sprintf("%v %v", d.uns.GetNamespace(), d.uns.GetName())
}
