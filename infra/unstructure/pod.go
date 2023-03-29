package unstructure

import (
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"log"
)

var _ familiar.Familiar = &Pod{}

type Pod struct {
	uns *unstructured.Unstructured
}

func (p *Pod) NeedMore() []*familiar.Id {
	serviceAccountName, exists, err := unstructured.NestedString(p.uns.Object, "spec", "serviceAccountName")
	if !exists || err != nil {
		log.Fatalf("%+v %+v", exists, err)
	}

	return []*familiar.Id{
		{
			GroupVersionKind: familiar.GroupVersionKind{
				Version: "v1",
				Kind:    "ServiceAccount",
			},
			Namespace: p.uns.GetNamespace(),
			Name:      serviceAccountName,
		},
	}
}
