package v1

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkflowMapper(t *testing.T) {
	ctx := context.Background()

	f, err := os.Open("testdata/valid.yaml")
	require.NoError(t, err)

	sd := NewDocumentStreamingDecoder(f, &YAMLDecoder{})

	wd, err := sd.DecodeStream(ctx)
	require.NoError(t, err)

	mapper := NewDefaultWorkflowMapper(
		WithNamespaceOption("valid-workflow"),
		WithWorkflowNameOption("valid-workflow-name"),
	)

	mapping, err := mapper.Map(wd)
	require.NoError(t, err)

	require.NotNil(t, mapping.Namespace)
	require.NotNil(t, mapping.Workflow)

	require.Equal(t, mapper.namespace, mapping.Namespace.GetName())
	require.Equal(t, mapper.name, mapping.Workflow.GetName())

	expectedSpec := map[string]interface{}{
		"tag": "v1",
	}

	expectedEnv := map[string]interface{}{
		"CI":      true,
		"RETRIES": 3,
	}

	require.Len(t, mapping.Workflow.Spec.Steps, 1)

	require.Equal(t, "step-1", mapping.Workflow.Spec.Steps[0].Name)

	require.Equal(t, "image-1", mapping.Workflow.Spec.Steps[0].Container.Image)

	require.Len(t, mapping.Workflow.Spec.Steps[0].Container.Spec, 1)
	require.Equal(t, expectedSpec, mapping.Workflow.Spec.Steps[0].Container.Spec.Value())

	require.Len(t, mapping.Workflow.Spec.Steps[0].Container.Env, 2)
	require.Equal(t, expectedEnv, mapping.Workflow.Spec.Steps[0].Container.Env.Value())

	require.Len(t, mapping.Workflow.Spec.Parameters, 1)
	require.Equal(t, "hi", mapping.Workflow.Spec.Parameters[0].Name)
	require.Equal(t, 5, mapping.Workflow.Spec.Parameters[0].Value.Value().(int))

	require.NoError(t, json.NewEncoder(io.Discard).Encode(mapping.Namespace))
	require.NoError(t, json.NewEncoder(io.Discard).Encode(mapping.Workflow))
}
