package monitor

type Resource struct {
	Group     string
	Version   string
	Kind      string
	Namespace string
	Name      string
}

type Client interface {
	List() ([]*Resource, error)
	Get() (*Resource, error)
}
