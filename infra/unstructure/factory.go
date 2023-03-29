package unstructure

import (
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ familiar.Factory = &Factory{}

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Create(manifest string) (familiar.Familiar, error) {
	var uns unstructured.Unstructured
	err := uns.UnmarshalJSON([]byte(manifest))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	switch uns.GroupVersionKind() {
	case schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "Deployment",
	}:
		return &Deployment{uns: &uns}, nil
	case schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "ReplicaSet",
	}:
		return &ReplicaSet{uns: &uns}, nil
	case schema.GroupVersionKind{
		Version: "v1",
		Kind:    "Pod",
	}:
		return &Pod{uns: &uns}, nil
	case schema.GroupVersionKind{
		Version: "v1",
		Kind:    "ServiceAccount",
	}:
		return &ServiceAccount{uns: &uns}, nil
	default:
		panic(uns.GroupVersionKind())
	}
}
