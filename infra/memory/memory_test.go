package memory

import (
	"github.com/biosvos/resource-checker-go/flow/monitor"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemoryNew(t *testing.T) {
	memory := NewMemory()

	require.NotNil(t, memory)
}

func TestMemoryList(t *testing.T) {
	memory := NewMemory()
	memory.AddResources(&monitor.Resource{
		Group:     "a",
		Version:   "b",
		Kind:      "c",
		Namespace: "d",
		Name:      "e",
	})

	list, err := memory.List("a", "b", "c", "d")

	require.NoError(t, err)
	require.Equal(t, 1, len(list))
	require.Equal(t, "e", list[0].Name)
}

func TestMemoryList_(t *testing.T) {
	memory := NewMemory()

	list, err := memory.List("a", "b", "c", "d")

	require.NoError(t, err)
	require.Equal(t, 0, len(list))
}

func TestMemoryGet(t *testing.T) {
	memory := NewMemory()
	memory.AddResources(&monitor.Resource{
		Group:     "a",
		Version:   "b",
		Kind:      "c",
		Namespace: "d",
		Name:      "e",
	})

	get, err := memory.Get("a", "b", "c", "d", "e")

	require.NoError(t, err)
	require.Equal(t, "e", get.Name)
}

func TestMemoryGet_(t *testing.T) {
	memory := NewMemory()

	_, err := memory.Get("a", "b", "c", "d", "e")

	require.Error(t, err)
}
