package v1

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkflowRunEngineMapping(t *testing.T) {
	mapper := NewDefaultRunEngineMapper(
		WithNamespaceRunOption("valid-workflow"),
		WithWorkflowNameRunOption("valid-workflow-name"),
		WithWorkflowRunNameRunOption("valid-workflow-run-name"),
	)

	manifest, err := mapper.ToRuntimeObjectsManifest()
	require.NoError(t, err)

	require.NotNil(t, manifest.Namespace)
	require.NotNil(t, manifest.WorkflowRun)

	require.Equal(t, "valid-workflow", manifest.Namespace.GetName())
	require.Equal(t, "valid-workflow-run-name", manifest.WorkflowRun.GetName())

	require.NoError(t, json.NewEncoder(ioutil.Discard).Encode(manifest.Namespace))
	require.NoError(t, json.NewEncoder(ioutil.Discard).Encode(manifest.WorkflowRun))
}
