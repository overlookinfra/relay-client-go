package v1

import (
	"encoding/json"
	"io"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTenantEngineMapping(t *testing.T) {
	eu, err := url.Parse("https://unit-testing.relay.sh/events")
	require.NoError(t, err)

	tu, err := url.Parse("https://unit-testing.relay.sh/workflow-run")
	require.NoError(t, err)

	mapper := NewDefaultTenantEngineMapper(
		WithNamespaceTenantOption("tenant-namespace"),
		WithNameTenantOption("test-tenant"),
		WithIDTenantOption("test-123"),
		WithWorkflowIDTenantOption("workflow-123"),
		WithEventURLTenantOption(eu),
		WithWorkflowExecutionURLTenantOption(tu),
		WithTokenSecretNameTenantOption("test-token-secret"),
	)

	manifest, err := mapper.ToRuntimeObjectsManifest()
	require.NoError(t, err)

	require.NotNil(t, manifest.Namespace)
	require.NotNil(t, manifest.Tenant)

	require.Equal(t, "test-tenant", manifest.Tenant.GetName())

	require.NotNil(t, manifest.Tenant.Spec.TriggerEventSink)
	require.Equal(t, eu.String(), manifest.Tenant.Spec.TriggerEventSink.API.URL)
	require.Equal(t, "test-token-secret", manifest.Tenant.Spec.TriggerEventSink.API.TokenFrom.SecretKeyRef.LocalObjectReference.Name)

	require.NotNil(t, manifest.Tenant.Spec.WorkflowExecutionSink)
	require.Equal(t, tu.String(), manifest.Tenant.Spec.WorkflowExecutionSink.API.URL)
	require.Equal(t, "test-token-secret", manifest.Tenant.Spec.WorkflowExecutionSink.API.TokenFrom.SecretKeyRef.LocalObjectReference.Name)

	require.NoError(t, json.NewEncoder(io.Discard).Encode(manifest.Namespace))
	require.NoError(t, json.NewEncoder(io.Discard).Encode(manifest.Tenant))
}
