package kubernetes

import (
	"context"
	"github.com/biosvos/resource-checker-go/flow/monitor"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/clientcmd"
	clientGo "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ monitor.Client = &Kubernetes{}

type Kubernetes struct {
	client clientGo.Client
}

func NewClient() (*Kubernetes, error) {
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientcmd.NewDefaultClientConfigLoadingRules(), nil)
	clientConfig, err := config.ClientConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	client, err := clientGo.New(clientConfig, clientGo.Options{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Kubernetes{
		client: client,
	}, nil
}

func (k *Kubernetes) List(group string, version string, kind string, namespace string, labels map[string]string) ([]*monitor.Resource, error) {
	list, err := getUnstructuredList(k.client, schema.GroupVersionKind{
		Group:   group,
		Version: version,
		Kind:    kind,
	}, clientGo.InNamespace(namespace), clientGo.MatchingLabels(labels))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return newResources(list), nil
}

func (k *Kubernetes) Get(group string, version string, kind string, namespace string, name string) (*monitor.Resource, error) {
	var uns unstructured.Unstructured
	uns.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   group,
		Version: version,
		Kind:    kind,
	})
	err := k.client.Get(context.Background(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, &uns)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return newResource(&uns), nil
}

func getUnstructuredList(client clientGo.Client, gvk schema.GroupVersionKind, opts ...clientGo.ListOption) (*unstructured.UnstructuredList, error) {
	unsList := unstructured.UnstructuredList{}
	unsList.SetGroupVersionKind(gvk)
	err := client.List(context.Background(), &unsList, opts...)
	return &unsList, errors.WithStack(err)
}
