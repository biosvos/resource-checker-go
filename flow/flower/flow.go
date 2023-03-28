package flower

import (
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"github.com/biosvos/resource-checker-go/flow/monitor"
	"github.com/pkg/errors"
)

var _ Flower = &Flow{}

type Flow struct {
	client  monitor.Client
	factory familiar.Factory
}

func NewFlow(client monitor.Client, factory familiar.Factory) *Flow {
	return &Flow{client: client, factory: factory}
}

type Iterator[T any] struct {
	elements []T
}

func NewIterator[T any](elements ...T) *Iterator[T] {
	return &Iterator[T]{elements: elements}
}

func (i *Iterator[T]) HasNext() bool {
	return len(i.elements) > 0
}

func (i *Iterator[T]) Next() T {
	ret := i.elements[0]
	i.elements = i.elements[1:]
	return ret
}

func (i *Iterator[T]) Add(element ...T) {
	i.elements = append(i.elements, element...)
}

func (f *Flow) GetFamily(resource *Resource) ([]*Resource, error) {
	mores := []*familiar.Id{
		{
			GroupVersionKind: familiar.GroupVersionKind{
				Group:   resource.Group,
				Version: resource.Version,
				Kind:    resource.Kind,
			},
			Namespace: resource.Namespace,
			Name:      resource.Name,
		},
	}
	iterator := NewIterator(mores...)
	var ret []*Resource
	for iterator.HasNext() {
		cur := iterator.Next()
		resources, err := f.requestMore(cur)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for _, clientResource := range resources {
			ret = append(ret, &Resource{
				GroupVersionKind: GroupVersionKind{
					Group:   clientResource.Group,
					Version: clientResource.Version,
					Kind:    clientResource.Kind,
				},
				Namespace: clientResource.Namespace,
				Name:      clientResource.Name,
			})
			factory, err := f.factory.Create(clientResource.Manifest)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			iterator.Add(factory.NeedMore()...)
		}
	}
	return ret, nil
}

func (f *Flow) requestMore(more *familiar.Id) ([]*monitor.Resource, error) {
	var ret []*monitor.Resource
	if more.Name == "" {
		resources, err := f.client.List(more.Group, more.Version, more.Kind, more.Namespace, more.Labels)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		ret = resources
	} else {
		resource, err := f.client.Get(more.Group, more.Version, more.Kind, more.Namespace, more.Name)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		ret = append(ret, resource)
	}
	return ret, nil
}
