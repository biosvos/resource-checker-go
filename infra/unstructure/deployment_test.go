package unstructure

import (
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func TestDeployment(t *testing.T) {
	var uns unstructured.Unstructured
	err := uns.UnmarshalJSON([]byte(chocoFailedDeployJson))
	require.NoError(t, err)
	deployment := &Deployment{uns: &uns}

	more := deployment.NeedMore()
	require.Equal(t, 1, len(more))
}
