//go:build kubernetes

package kubernetes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKubernetesNew(t *testing.T) {
	client, err := NewClient()

	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestKubernetesList(t *testing.T) {
	client, _ := NewClient()

	list, err := client.List("apps", "v1", "deployment", "wow", nil)

	require.NoError(t, err)
	require.NotNil(t, list)
}

func TestKubernetesListWithSelector(t *testing.T) {
	client, _ := NewClient()

	resources, err := client.List("apps", "v1", "ReplicaSet", "wow", map[string]string{
		"name": "choco",
	})

	require.NoError(t, err)
	require.Equal(t, 1, len(resources))
}
