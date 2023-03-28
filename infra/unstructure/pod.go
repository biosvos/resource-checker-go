package unstructure

import (
	"fmt"
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var _ familiar.Familiar = &Pod{}

type Pod struct {
	uns *unstructured.Unstructured
}

func (p *Pod) NeedMore() []*familiar.Id {
	return nil
}

func (p *Pod) String() string {
	return fmt.Sprintf("%v %v", p.uns.GetNamespace(), p.uns.GetName())
}
