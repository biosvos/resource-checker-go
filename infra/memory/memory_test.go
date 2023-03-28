package memory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemoryNew(t *testing.T) {
	memory := NewMemory()

	require.NotNil(t, memory)
}

func TestMemoryList(t *testing.T) {
	memory := NewMemory()
	memory.AddResources(chocoFailedDeployJson)

	list, err := memory.List("apps", "v1", "Deployment", "wow", nil)

	require.NoError(t, err)
	require.Equal(t, 1, len(list))
	require.Equal(t, "choco", list[0].Name)
}

func TestMemoryList_(t *testing.T) {
	memory := NewMemory()

	list, err := memory.List("a", "b", "c", "d", nil)

	require.NoError(t, err)
	require.Equal(t, 0, len(list))
}
