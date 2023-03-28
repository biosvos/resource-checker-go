package monitor

type ResourceStatus string

const (
	StatusFailed  = ResourceStatus("Failed")
	StatusSuccess = ResourceStatus("Success")
)

type Status struct {
	Status ResourceStatus
	Reason string
}

type Resource struct {
	Group     string
	Version   string
	Kind      string
	Namespace string
	Name      string
	Status    Status
}

type Client interface {
	List(group string, version string, kind string, namespace string) ([]*Resource, error)
}
