package memory

import (
	"encoding/json"
	"github.com/biosvos/resource-checker-go/flow/monitor"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
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
	elements map[GroupVersionKindNamespace]map[GroupVersionKindNamespaceName]*unstructured.Unstructured
}

func NewMemory() *Memory {
	return &Memory{
		elements: map[GroupVersionKindNamespace]map[GroupVersionKindNamespaceName]*unstructured.Unstructured{},
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
	for _, uns := range m.elements[ns] {
		resource := newResource(uns)
		ret = append(ret, resource)
	}
	return ret, nil
}

func (m *Memory) AddResources(manifests ...string) {
	for _, manifest := range manifests {
		uns := newUnstructured(manifest)
		resource := newResource(uns)
		namespace := GroupVersionKindNamespace{
			Group:     resource.Group,
			Version:   resource.Version,
			Kind:      resource.Kind,
			Namespace: resource.Namespace,
		}
		_, ok := m.elements[namespace]
		if !ok {
			m.elements[namespace] = map[GroupVersionKindNamespaceName]*unstructured.Unstructured{}
		}
		name := GroupVersionKindNamespaceName{
			Group:     resource.Group,
			Version:   resource.Version,
			Kind:      resource.Kind,
			Namespace: resource.Namespace,
			Name:      resource.Name,
		}
		m.elements[namespace][name] = uns
	}
}

func newResource(uns *unstructured.Unstructured) *monitor.Resource {
	return &monitor.Resource{
		Group:     uns.GroupVersionKind().Group,
		Version:   uns.GroupVersionKind().Version,
		Kind:      uns.GroupVersionKind().Kind,
		Namespace: uns.GetNamespace(),
		Name:      uns.GetName(),
	}
}

func newUnstructured(manifest string) *unstructured.Unstructured {
	var uns unstructured.Unstructured
	err := uns.UnmarshalJSON([]byte(manifest))
	if err != nil {
		panic(err)
	}
	return &uns
}

func decisionStatus(uns *unstructured.Unstructured) monitor.Status {
	switch uns.GroupVersionKind() {
	case schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "Deployment",
	}:
		return decisionDeploymentStatus(uns)
	default:
		panic(uns.GroupVersionKind())
	}
}

type deploymentStatus struct {
	Status struct {
		Conditions []struct {
			LastTransitionTime time.Time `json:"lastTransitionTime"`
			LastUpdateTime     time.Time `json:"lastUpdateTime"`
			Message            string    `json:"message"`
			Reason             string    `json:"reason"`
			Status             string    `json:"status"`
			Type               string    `json:"type"`
		} `json:"conditions"`
		ObservedGeneration  int `json:"observedGeneration"`
		Replicas            int `json:"replicas"`
		UnavailableReplicas int `json:"unavailableReplicas"`
		UpdatedReplicas     int `json:"updatedReplicas"`
	} `json:"status"`
}

func decisionDeploymentStatus(uns *unstructured.Unstructured) monitor.Status {
	marshal, err := json.Marshal(uns.Object)
	if err != nil {
		panic(err)
	}
	var status deploymentStatus
	err = json.Unmarshal(marshal, &status)
	if err != nil {
		panic(err)
	}
	if status.Status.UnavailableReplicas > 0 {
		var reasons []string
		for _, condition := range status.Status.Conditions {
			if condition.Status == "False" {
				reasons = append(reasons, condition.Reason)
			}
		}

		return monitor.Status{
			Status:  monitor.StatusFailed,
			Reasons: reasons,
		}
	}

	return monitor.Status{
		Status: monitor.StatusSuccess,
	}
}
