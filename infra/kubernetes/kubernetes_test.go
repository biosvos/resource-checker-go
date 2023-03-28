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

func TestKubernetesGet(t *testing.T) {
	client, _ := NewClient()

	resource, err := client.Get("apps", "v1", "deployment", "wow", "deployment-not-owner")

	require.NoError(t, err)
	require.NotNil(t, resource)
}

func TestKubernetesList(t *testing.T) {
	client, _ := NewClient()

	list, err := client.List("apps", "v1", "deployment", "wow")

	require.NoError(t, err)
	require.NotNil(t, list)
}
