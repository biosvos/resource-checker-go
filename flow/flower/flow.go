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

func (f *Flow) GetFamily(resource *Resource) ([]*Resource, error) {
	clientResource, err := f.client.Get(resource.Group, resource.Version, resource.Kind, resource.Namespace, resource.Name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	famil, err := f.factory.Create(clientResource.Manifest)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, more := range mores {

	}
}
