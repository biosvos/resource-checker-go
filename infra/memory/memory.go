package memory

import (
	"errors"
	"github.com/biosvos/resource-checker-go/flow/monitor"
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

func (m *Memory) Get(group string, version string, kind string, namespace string, name string) (*monitor.Resource, error) {
	ns := GroupVersionKindNamespace{
		Group:     group,
		Version:   version,
		Kind:      kind,
		Namespace: namespace,
	}
	names := GroupVersionKindNamespaceName{
		Group:     group,
		Version:   version,
		Kind:      kind,
		Namespace: namespace,
		Name:      name,
	}
	ret, ok := m.elements[ns][names]
	if !ok {
		return nil, errors.New("failed to get resource")
	}
	return ret, nil
}

func (m *Memory) AddResources(resources ...*monitor.Resource) {
	for _, resource := range resources {
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
		m.elements[namespace][name] = resource
	}
}
