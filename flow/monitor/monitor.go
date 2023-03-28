package monitor

type Resource struct {
	Group     string
	Version   string
	Kind      string
	Namespace string
	Name      string
	Manifest  string
}

type Client interface {
	List(group string, version string, kind string, namespace string) ([]*Resource, error)
}
