package v1

import (
	nebulav1 "github.com/puppetlabs/relay-core/pkg/apis/nebula.puppet.com/v1"
	"github.com/puppetlabs/relay-core/pkg/apis/relay.sh/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

const (
	AccountIDLabel         = "managed.relay.sh/account-id"
	WorkflowIDLabel        = "managed.relay.sh/workflow-id"
	WorkflowTriggerIDLabel = "managed.relay.sh/workflow-trigger-id"
)

const (
	defaultWorkflowName    = "workflow"
	defaultWorkflowRunName = "workflow-run"
	defaultNamespace       = "default"
)

type RunKubernetesObjectMapping struct {
	Namespace   *corev1.Namespace
	WorkflowRun *nebulav1.WorkflowRun
}

type TenantKubernetesObjectMapping struct {
	Namespace *corev1.Namespace
	Tenant    *v1beta1.Tenant
}

type WebhookTriggerKubernetesObjectMapping struct {
	WebhookTrigger *v1beta1.WebhookTrigger
}

type WorkflowMapping struct {
	Namespace *corev1.Namespace
	Workflow  *v1beta1.Workflow
}

type RunKubernetesEngineMapper interface {
	ToRuntimeObjectsManifest(*WorkflowData) (*RunKubernetesObjectMapping, error)
}

type TenantKubernetesEngineMapper interface {
	ToRuntimeObjectsManifest(id string) (*TenantKubernetesObjectMapping, error)
}

type WebhookTriggerKubernetesEngineMapper interface {
	ToRuntimeObjectsManifest(source *WebhookWorkflowTriggerSource) (*WebhookTriggerKubernetesObjectMapping, error)
}

type WorkflowMapper interface {
	Map(*WorkflowData) (*WorkflowMapping, error)
}
