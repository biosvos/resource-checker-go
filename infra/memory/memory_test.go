package memory

import (
	"github.com/biosvos/resource-checker-go/flow/monitor"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func TestMemoryNew(t *testing.T) {
	memory := NewMemory()

	require.NotNil(t, memory)
}

func TestMemoryList(t *testing.T) {
	memory := NewMemory()
	memory.AddResources(chocoFailedDeployJson)

	list, err := memory.List("apps", "v1", "Deployment", "wow")

	require.NoError(t, err)
	require.Equal(t, 1, len(list))
	require.Equal(t, "choco", list[0].Name)
}

func TestMemoryList_(t *testing.T) {
	memory := NewMemory()

	list, err := memory.List("a", "b", "c", "d")

	require.NoError(t, err)
	require.Equal(t, 0, len(list))
}

func TestMemoryGetResourceStatusFailed(t *testing.T) {
	var uns unstructured.Unstructured
	_ = uns.UnmarshalJSON([]byte(chocoFailedDeployJson))

	status := decisionStatus(&uns)

	require.Equal(t, monitor.StatusFailed, status.Status)
	require.Condition(t, func() (success bool) {
		return len(status.Reasons) > 0
	})
}
