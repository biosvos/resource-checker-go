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
		return &OwnerDecorator{
			uns:  &uns,
			wrap: &Deployment{uns: &uns},
		}, nil
	case schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "ReplicaSet",
	}:
		return &OwnerDecorator{
			uns:  &uns,
			wrap: &ReplicaSet{uns: &uns},
		}, nil
	case schema.GroupVersionKind{
		Version: "v1",
		Kind:    "Pod",
	}:
		return &OwnerDecorator{
			uns:  &uns,
			wrap: &Pod{uns: &uns},
		}, nil
	case schema.GroupVersionKind{
		Version: "v1",
		Kind:    "ServiceAccount",
	}:
		return &OwnerDecorator{
			uns:  &uns,
			wrap: &ServiceAccount{uns: &uns},
		}, nil
	case schema.GroupVersionKind{
		Version: "v1",
		Kind:    "ConfigMap",
	}:
		return &OwnerDecorator{
			uns:  &uns,
			wrap: &Empty{},
		}, nil
	case schema.GroupVersionKind{
		Version: "v1",
		Kind:    "Secret",
	}:
		return &OwnerDecorator{
			uns:  &uns,
			wrap: &Empty{},
		}, nil
	default:
		panic(uns.GroupVersionKind())
	}
}
