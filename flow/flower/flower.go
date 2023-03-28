package flower

type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

type Resource struct {
	GroupVersionKind
	Namespace string
	Name      string
}

type Flower interface {
	GetFamily(resource *Resource) ([]*Resource, error)
}
