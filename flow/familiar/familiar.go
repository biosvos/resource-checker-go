package familiar

type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

type Id struct {
	GroupVersionKind
	Namespace string
	Name      string
	Labels    map[string]string
}

type Familiar interface {
	NeedMore() []*Id
}

type Factory interface {
	Create(manifest string) (Familiar, error)
}
