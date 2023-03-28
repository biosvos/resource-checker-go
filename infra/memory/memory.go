package memory

import (
	"github.com/biosvos/resource-checker-go/flow/monitor"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var _ monitor.Client = &Memory{}

type GroupVersionKindNamespaceName struct {
	Group     string
	Version   string
	Kind      string
	Namespace string
	Name      string
}

type GroupVersionKindNamespace struct {
	Group     string
	Version   string
	Kind      string
	Namespace string
}

type Memory struct {
	elements map[GroupVersionKindNamespace]map[GroupVersionKindNamespaceName]*monitor.Resource
}

func NewMemory() *Memory {
	return &Memory{
		elements: map[GroupVersionKindNamespace]map[GroupVersionKindNamespaceName]*monitor.Resource{},
	}
}

func (m *Memory) List(group string, version string, kind string, namespace string) ([]*monitor.Resource, error) {
	ns := GroupVersionKindNamespace{
		Group:     group,
		Version:   version,
		Kind:      kind,
		Namespace: namespace,
	}
	var ret []*monitor.Resource
	for _, resource := range m.elements[ns] {
		ret = append(ret, resource)
	}
	return ret, nil
}

func (m *Memory) AddResources(manifests ...string) {
	for _, manifest := range manifests {
		var uns unstructured.Unstructured
		err := uns.UnmarshalJSON([]byte(manifest))
		if err != nil {
			panic(err)
		}
		resource := monitor.Resource{
			Group:     uns.GroupVersionKind().Group,
			Version:   uns.GroupVersionKind().Version,
			Kind:      uns.GroupVersionKind().Kind,
			Namespace: uns.GetNamespace(),
			Name:      uns.GetName(),
		}
		namespace := GroupVersionKindNamespace{
			Group:     resource.Group,
			Version:   resource.Version,
			Kind:      resource.Kind,
			Namespace: resource.Namespace,
		}
		_, ok := m.elements[namespace]
		if !ok {
			m.elements[namespace] = map[GroupVersionKindNamespaceName]*monitor.Resource{}
		}
		name := GroupVersionKindNamespaceName{
			Group:     resource.Group,
			Version:   resource.Version,
			Kind:      resource.Kind,
			Namespace: resource.Namespace,
			Name:      resource.Name,
		}
		m.elements[namespace][name] = &resource
	}
}
