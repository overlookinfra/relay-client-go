package v1

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/puppetlabs/relay-core/pkg/spec"
	"github.com/stretchr/testify/require"
)

func TestWebhookTriggerMapping(t *testing.T) {
	tenantMapper := NewDefaultTenantEngineMapper(
		WithIDTenantOption("test-1234"),
		WithNameTenantOption("test-name"),
		WithWorkflowIDTenantOption("workflow-1234"),
		WithNamespaceTenantOption("test-tenant"),
	)

	tenant, err := tenantMapper.ToRuntimeObjectsManifest()
	require.NoError(t, err)

	mapper := NewDefaultWebhookTriggerEngineMapper(
		WithIDWebhookTriggerOption("test-1234"),
		WithNameWebhookTriggerOption("test-webhook-trigger"),
		WithImageWebhookTriggerOption("test-image"),
	)

	source := &WebhookWorkflowTriggerSource{
		ContainerMixin: ContainerMixin{
			Image: "test-image:latest",
			Spec: ExpressionMap{
				"tag": spec.JSONTree{Tree: "v1"},
			},
			Env: ExpressionMap{
				"CI":      spec.JSONTree{Tree: true},
				"RETRIES": spec.JSONTree{Tree: 3},
			},
		},
	}

	manifest, err := mapper.ToRuntimeObjectsManifest(tenant.Tenant, source)
	require.NoError(t, err)

	require.NotNil(t, manifest.WebhookTrigger)
	require.Equal(t, tenant.Tenant.GetNamespace(), manifest.WebhookTrigger.GetNamespace())

	require.Len(t, manifest.WebhookTrigger.Spec.Spec, 1)
	require.Len(t, manifest.WebhookTrigger.Spec.Env, 2)

	require.NoError(t, json.NewEncoder(io.Discard).Encode(manifest.WebhookTrigger))
}
