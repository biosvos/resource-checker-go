package unstructure

import (
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var _ familiar.Familiar = &ServiceAccount{}

type ServiceAccount struct {
	uns *unstructured.Unstructured
}

func (s *ServiceAccount) NeedMore() []*familiar.Id {
	return nil
}

func (s *ServiceAccount) String() string {
	return ""
}
