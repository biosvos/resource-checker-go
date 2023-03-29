package unstructure

import (
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ familiar.Familiar = &OwnerDecorator{}

type OwnerDecorator struct {
	uns  *unstructured.Unstructured
	wrap familiar.Familiar
}

func (o *OwnerDecorator) NeedMore() []*familiar.Id {
	var ret []*familiar.Id
	for _, reference := range o.uns.GetOwnerReferences() {
		gv, err := schema.ParseGroupVersion(reference.APIVersion)
		if err != nil {
			panic(err)
		}
		ret = append(ret, &familiar.Id{
			GroupVersionKind: familiar.GroupVersionKind{
				Group:   gv.Group,
				Version: gv.Version,
				Kind:    reference.Kind,
			},
			Namespace: o.uns.GetNamespace(),
			Name:      reference.Name,
		})
	}
	ret = append(ret, o.wrap.NeedMore()...)
	return ret
}
