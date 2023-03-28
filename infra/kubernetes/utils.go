package kubernetes

import (
	"github.com/biosvos/resource-checker-go/flow/monitor"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func newResource(uns *unstructured.Unstructured) *monitor.Resource {
	return &monitor.Resource{
		Group:     uns.GroupVersionKind().Group,
		Version:   uns.GroupVersionKind().Version,
		Kind:      uns.GroupVersionKind().Kind,
		Namespace: uns.GetNamespace(),
		Name:      uns.GetName(),
	}
}

func newResources(unsList *unstructured.UnstructuredList) []*monitor.Resource {
	var ret []*monitor.Resource
	for _, uns := range unsList.Items {
		ret = append(ret, newResource(&uns))
	}
	return ret
}
