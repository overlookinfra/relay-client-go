/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// WorkflowStep - An individual workflow step
type WorkflowStep struct {
	ApprovalWorkflowStep *ApprovalWorkflowStep
	ContainerWorkflowStep *ContainerWorkflowStep
}

// ApprovalWorkflowStepAsWorkflowStep is a convenience function that returns ApprovalWorkflowStep wrapped in WorkflowStep
func ApprovalWorkflowStepAsWorkflowStep(v *ApprovalWorkflowStep) WorkflowStep {
	return WorkflowStep{ ApprovalWorkflowStep: v}
}

// ContainerWorkflowStepAsWorkflowStep is a convenience function that returns ContainerWorkflowStep wrapped in WorkflowStep
func ContainerWorkflowStepAsWorkflowStep(v *ContainerWorkflowStep) WorkflowStep {
	return WorkflowStep{ ContainerWorkflowStep: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WorkflowStep) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApprovalWorkflowStep
	err = json.Unmarshal(data, &dst.ApprovalWorkflowStep)
	if err == nil {
		jsonApprovalWorkflowStep, _ := json.Marshal(dst.ApprovalWorkflowStep)
		if string(jsonApprovalWorkflowStep) == "{}" { // empty struct
			dst.ApprovalWorkflowStep = nil
		} else {
			match++
		}
	} else {
		dst.ApprovalWorkflowStep = nil
	}

	// try to unmarshal data into ContainerWorkflowStep
	err = json.Unmarshal(data, &dst.ContainerWorkflowStep)
	if err == nil {
		jsonContainerWorkflowStep, _ := json.Marshal(dst.ContainerWorkflowStep)
		if string(jsonContainerWorkflowStep) == "{}" { // empty struct
			dst.ContainerWorkflowStep = nil
		} else {
			match++
		}
	} else {
		dst.ContainerWorkflowStep = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApprovalWorkflowStep = nil
		dst.ContainerWorkflowStep = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(WorkflowStep)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(WorkflowStep)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WorkflowStep) MarshalJSON() ([]byte, error) {
	if src.ApprovalWorkflowStep != nil {
		return json.Marshal(&src.ApprovalWorkflowStep)
	}

	if src.ContainerWorkflowStep != nil {
		return json.Marshal(&src.ContainerWorkflowStep)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkflowStep) GetActualInstance() (interface{}) {
	if obj.ApprovalWorkflowStep != nil {
		return obj.ApprovalWorkflowStep
	}

	if obj.ContainerWorkflowStep != nil {
		return obj.ContainerWorkflowStep
	}

	// all schemas are nil
	return nil
}

type NullableWorkflowStep struct {
	value *WorkflowStep
	isSet bool
}

func (v NullableWorkflowStep) Get() *WorkflowStep {
	return v.value
}

func (v *NullableWorkflowStep) Set(val *WorkflowStep) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStep) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStep(val *WorkflowStep) *NullableWorkflowStep {
	return &NullableWorkflowStep{value: val, isSet: true}
}

func (v NullableWorkflowStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


