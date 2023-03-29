package flower

import (
	"github.com/biosvos/resource-checker-go/flow/familiar"
	"github.com/biosvos/resource-checker-go/flow/monitor"
	"github.com/biosvos/structures"
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

type Checker struct {
	Group     string
	Version   string
	Kind      string
	Namespace string
	Name      string
}

func (c Checker) Identify() Checker {
	return c
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
	iterator := structures.NewIterator(mores...)
	checker := structures.NewSet[Checker, Checker]()
	var ret []*Resource
	for iterator.HasNext() {
		cur := iterator.Next()
		if cur.Name != "" {
			check := Checker{
				Group:     cur.Group,
				Version:   cur.Version,
				Kind:      cur.Kind,
				Namespace: cur.Namespace,
				Name:      cur.Name,
			}
			if checker.Has(check) {
				continue
			}
			checker.Add(check)
		}

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
				Manifest:  clientResource.Manifest,
			})
			factory, err := f.factory.Create(clientResource.Manifest)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			iterator.Add(factory.NeedMore()...)

			check := Checker{
				Group:     clientResource.Group,
				Version:   clientResource.Version,
				Kind:      clientResource.Kind,
				Namespace: clientResource.Namespace,
				Name:      clientResource.Name,
			}
			checker.Add(check)
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
