package v1

import (
	"path"

	relayv1beta1 "github.com/puppetlabs/relay-core/pkg/apis/relay.sh/v1beta1"
	"github.com/puppetlabs/relay-core/pkg/model"
	"github.com/puppetlabs/relay-core/pkg/spec"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DefaultWorkflowMapperOption func(*DefaultWorkflowMapper)

func WithDomainIDOption(id string) DefaultWorkflowMapperOption {
	return func(m *DefaultWorkflowMapper) {
		m.domainID = id
	}
}

func WithWorkflowNameOption(name string) DefaultWorkflowMapperOption {
	return func(m *DefaultWorkflowMapper) {
		m.name = name
	}
}

func WithNamespaceOption(ns string) DefaultWorkflowMapperOption {
	return func(m *DefaultWorkflowMapper) {
		m.namespace = ns
	}
}

func WithVaultEngineMountOption(mount string) DefaultWorkflowMapperOption {
	return func(m *DefaultWorkflowMapper) {
		m.vaultEngineMount = mount
	}
}

func WithTenantOption(tenant *relayv1beta1.Tenant) DefaultWorkflowMapperOption {
	return func(m *DefaultWorkflowMapper) {
		m.tenant = tenant
	}
}

type DefaultWorkflowMapper struct {
	name             string
	namespace        string
	domainID         string
	vaultEngineMount string
	tenant           *relayv1beta1.Tenant
}

func (m *DefaultWorkflowMapper) Map(wd *WorkflowData) (*WorkflowMapping, error) {
	mapping := WorkflowMapping{}

	if m.namespace != defaultNamespace {
		mapping.Namespace = mapNamespace(m.namespace)
	}

	wp := make([]*relayv1beta1.Parameter, 0)
	for k, v := range wd.Parameters {
		if def, ok := v.Default(); ok {
			value := relayv1beta1.AsUnstructured(def)
			wp = append(wp, &relayv1beta1.Parameter{
				Name:  k,
				Value: &value,
			})
		}
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

	mapping.Workflow = &relayv1beta1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			Name:        m.name,
			Namespace:   m.namespace,
			Annotations: annotations,
		},
		Spec: relayv1beta1.WorkflowSpec{
			Parameters: wp,
			Steps:      mapWorkflowSteps(wd),
		},
	}

	if m.tenant != nil {
		mapping.Workflow.Spec.TenantRef = corev1.LocalObjectReference{
			Name: m.tenant.GetName(),
		}
	}

	return &mapping, nil
}

func NewDefaultWorkflowMapper(opts ...DefaultWorkflowMapperOption) *DefaultWorkflowMapper {
	m := &DefaultWorkflowMapper{
		name:      defaultWorkflowName,
		namespace: defaultNamespace,
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func mapWorkflowSteps(wd *WorkflowData) []*relayv1beta1.Step {
	var workflowSteps []*relayv1beta1.Step

	for _, value := range wd.Steps {
		when := relayv1beta1.AsUnstructured(value.When.Tree)
		workflowStep := relayv1beta1.Step{
			Name:      value.Name,
			DependsOn: value.DependsOn,
			When:      &when,
		}

		switch variant := value.Variant.(type) {
		case *ContainerWorkflowStep:
			workflowStep.Container = relayv1beta1.Container{
				Image:   variant.Image,
				Spec:    mapSpec(variant.Spec),
				Env:     mapSpec(variant.Env),
				Input:   variant.Input,
				Command: variant.Command,
				Args:    variant.Args,
			}
		}

		workflowSteps = append(workflowSteps, &workflowStep)
	}

	return workflowSteps
}

func mapSpec(jm map[string]spec.JSONTree) relayv1beta1.UnstructuredObject {
	uo := make(relayv1beta1.UnstructuredObject, len(jm))
	for k, v := range jm {
		// The inner data type has to be compatible with transfer.JSONInterface
		// here, hence the explicit cast to interface{}.
		uo[k] = relayv1beta1.AsUnstructured(interface{}(v.Tree))
	}

	return uo
}
