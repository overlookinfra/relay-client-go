package v1

import (
	"path"

	nebulav1 "github.com/puppetlabs/relay-core/pkg/apis/nebula.puppet.com/v1"
	"github.com/puppetlabs/relay-core/pkg/apis/relay.sh/v1beta1"
	"github.com/puppetlabs/relay-core/pkg/model"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DefaultRunEngineMapperOption is a func that takes a *DefaultRunEngineMapper and
// configures it. Each function is responsible for a small configuration, such
// as setting the name field.
type DefaultRunEngineMapperOption func(*DefaultRunEngineMapper)

func WithDomainIDRunOption(id string) DefaultRunEngineMapperOption {
	return func(m *DefaultRunEngineMapper) {
		m.domainID = id
	}
}

func WithWorkflowNameRunOption(name string) DefaultRunEngineMapperOption {
	return func(m *DefaultRunEngineMapper) {
		m.name = name
	}
}

func WithWorkflowRunNameRunOption(name string) DefaultRunEngineMapperOption {
	return func(m *DefaultRunEngineMapper) {
		m.runName = name
	}
}

func WithNamespaceRunOption(ns string) DefaultRunEngineMapperOption {
	return func(m *DefaultRunEngineMapper) {
		m.namespace = ns
	}
}

func WithRunParametersRunOption(params WorkflowRunParameters) DefaultRunEngineMapperOption {
	return func(m *DefaultRunEngineMapper) {
		m.runParameters = params
	}
}

func WithVaultEngineMountRunOption(mount string) DefaultRunEngineMapperOption {
	return func(m *DefaultRunEngineMapper) {
		m.vaultEngineMount = mount
	}
}

func WithTenantRunOption(tenant *v1beta1.Tenant) DefaultRunEngineMapperOption {
	return func(m *DefaultRunEngineMapper) {
		m.tenant = tenant
	}
}

func WithWorkflowRunOption(workflow *v1beta1.Workflow) DefaultRunEngineMapperOption {
	return func(m *DefaultRunEngineMapper) {
		m.workflow = workflow
	}
}

// DefaultRunEngineMapper maps a WorkflowRun to Kubernetes runtime objects. It
// is the default for relay-operator.
type DefaultRunEngineMapper struct {
	name             string
	runName          string
	namespace        string
	runParameters    WorkflowRunParameters
	domainID         string
	vaultEngineMount string
	tenant           *v1beta1.Tenant
	workflow         *v1beta1.Workflow
}

// ToRuntimeObjectsManifest returns a RunKubernetesObjectMapping that contains
// uncreated objects that map to relay-core CRDs and other kubernetes resources
// required to support a run.
func (m *DefaultRunEngineMapper) ToRuntimeObjectsManifest() (*RunKubernetesObjectMapping, error) {
	manifest := RunKubernetesObjectMapping{}

	if m.namespace != defaultNamespace {
		manifest.Namespace = mapNamespace(m.namespace)
	}

	wrp := map[string]interface{}{}
	for k, v := range m.runParameters {
		wrp[k] = v.Value
	}

	annotations := map[string]string{
		model.RelayDomainIDAnnotation: m.domainID,
		model.RelayTenantIDAnnotation: m.name,
	}

	if m.vaultEngineMount != "" {
		annotations[model.RelayVaultEngineMountAnnotation] = m.vaultEngineMount
		annotations[model.RelayVaultSecretPathAnnotation] = path.Join("workflows", m.name)
		annotations[model.RelayVaultConnectionPathAnnotation] = path.Join("connections", m.domainID)
	}

	manifest.WorkflowRun = &nebulav1.WorkflowRun{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.runName,
			Namespace: m.namespace,
			// TODO
			Annotations: annotations,
		},
		Spec: nebulav1.WorkflowRunSpec{
			Name:       m.runName,
			Parameters: v1beta1.NewUnstructuredObject(wrp),
		},
	}

	if m.tenant != nil {
		manifest.WorkflowRun.Spec.TenantRef = &corev1.LocalObjectReference{
			Name: m.tenant.GetName(),
		}
	}

	if m.workflow != nil {
		manifest.WorkflowRun.Spec.WorkflowRef = corev1.LocalObjectReference{
			Name: m.workflow.GetName(),
		}
	}

	return &manifest, nil
}

// NewDefaultRunEngineMapper takes any number of DefaultRunEngineMapperOption's and
// returns a configured KubernetesEngineMapper.
func NewDefaultRunEngineMapper(opts ...DefaultRunEngineMapperOption) *DefaultRunEngineMapper {
	m := &DefaultRunEngineMapper{
		name:      defaultWorkflowName,
		runName:   defaultWorkflowRunName,
		namespace: defaultNamespace,
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func mapNamespace(ns string) *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: ns,
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "nebula",
				// This label controls how network policies external to the API
				// allow access to the Vault server. Do not remove it!
				"nebula.puppet.com/network-policy.customer": "true",
				// This label controls RBAC access to the namespace by the
				// controller.
				"controller.relay.sh/tenant-workload": "true",
			},
		},
	}
}
