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
	List(group string, version string, kind string, namespace string, labels map[string]string) ([]*Resource, error)
	Get(group string, version string, kind string, namespace string, name string) (*Resource, error)
}
