package unstructure

import (
	"fmt"
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"log"
)

var _ familiar.Familiar = &ReplicaSet{}

type ReplicaSet struct {
	uns *unstructured.Unstructured
}

func (r *ReplicaSet) NeedMore() []*familiar.Id {
	selector, exists, err := unstructured.NestedStringMap(r.uns.Object, "spec", "selector", "matchLabels")
	if !exists || err != nil {
		log.Fatalf("%+v %+v", exists, err)
	}
	return []*familiar.Id{
		{
			GroupVersionKind: familiar.GroupVersionKind{
				Group:   "",
				Version: "v1",
				Kind:    "Pod",
			},
			Namespace: r.uns.GetNamespace(),
			Labels:    selector,
		},
	}
}

func (r *ReplicaSet) String() string {
	return fmt.Sprintf("%v %v", r.uns.GetNamespace(), r.uns.GetName())
}
