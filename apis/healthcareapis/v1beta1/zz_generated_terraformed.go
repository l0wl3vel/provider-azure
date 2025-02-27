/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this HealthcareDICOMService
func (mg *HealthcareDICOMService) GetTerraformResourceType() string {
	return "azurerm_healthcare_dicom_service"
}

// GetConnectionDetailsMapping for this HealthcareDICOMService
func (tr *HealthcareDICOMService) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HealthcareDICOMService
func (tr *HealthcareDICOMService) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HealthcareDICOMService
func (tr *HealthcareDICOMService) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HealthcareDICOMService
func (tr *HealthcareDICOMService) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HealthcareDICOMService
func (tr *HealthcareDICOMService) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HealthcareDICOMService
func (tr *HealthcareDICOMService) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HealthcareDICOMService using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HealthcareDICOMService) LateInitialize(attrs []byte) (bool, error) {
	params := &HealthcareDICOMServiceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HealthcareDICOMService) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this HealthcareFHIRService
func (mg *HealthcareFHIRService) GetTerraformResourceType() string {
	return "azurerm_healthcare_fhir_service"
}

// GetConnectionDetailsMapping for this HealthcareFHIRService
func (tr *HealthcareFHIRService) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HealthcareFHIRService
func (tr *HealthcareFHIRService) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HealthcareFHIRService
func (tr *HealthcareFHIRService) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HealthcareFHIRService
func (tr *HealthcareFHIRService) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HealthcareFHIRService
func (tr *HealthcareFHIRService) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HealthcareFHIRService
func (tr *HealthcareFHIRService) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HealthcareFHIRService using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HealthcareFHIRService) LateInitialize(attrs []byte) (bool, error) {
	params := &HealthcareFHIRServiceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HealthcareFHIRService) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this HealthcareMedtechService
func (mg *HealthcareMedtechService) GetTerraformResourceType() string {
	return "azurerm_healthcare_medtech_service"
}

// GetConnectionDetailsMapping for this HealthcareMedtechService
func (tr *HealthcareMedtechService) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HealthcareMedtechService
func (tr *HealthcareMedtechService) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HealthcareMedtechService
func (tr *HealthcareMedtechService) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HealthcareMedtechService
func (tr *HealthcareMedtechService) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HealthcareMedtechService
func (tr *HealthcareMedtechService) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HealthcareMedtechService
func (tr *HealthcareMedtechService) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HealthcareMedtechService using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HealthcareMedtechService) LateInitialize(attrs []byte) (bool, error) {
	params := &HealthcareMedtechServiceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HealthcareMedtechService) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this HealthcareMedtechServiceFHIRDestination
func (mg *HealthcareMedtechServiceFHIRDestination) GetTerraformResourceType() string {
	return "azurerm_healthcare_medtech_service_fhir_destination"
}

// GetConnectionDetailsMapping for this HealthcareMedtechServiceFHIRDestination
func (tr *HealthcareMedtechServiceFHIRDestination) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HealthcareMedtechServiceFHIRDestination
func (tr *HealthcareMedtechServiceFHIRDestination) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HealthcareMedtechServiceFHIRDestination
func (tr *HealthcareMedtechServiceFHIRDestination) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HealthcareMedtechServiceFHIRDestination
func (tr *HealthcareMedtechServiceFHIRDestination) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HealthcareMedtechServiceFHIRDestination
func (tr *HealthcareMedtechServiceFHIRDestination) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HealthcareMedtechServiceFHIRDestination
func (tr *HealthcareMedtechServiceFHIRDestination) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HealthcareMedtechServiceFHIRDestination using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HealthcareMedtechServiceFHIRDestination) LateInitialize(attrs []byte) (bool, error) {
	params := &HealthcareMedtechServiceFHIRDestinationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HealthcareMedtechServiceFHIRDestination) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this HealthcareService
func (mg *HealthcareService) GetTerraformResourceType() string {
	return "azurerm_healthcare_service"
}

// GetConnectionDetailsMapping for this HealthcareService
func (tr *HealthcareService) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HealthcareService
func (tr *HealthcareService) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HealthcareService
func (tr *HealthcareService) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HealthcareService
func (tr *HealthcareService) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HealthcareService
func (tr *HealthcareService) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HealthcareService
func (tr *HealthcareService) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HealthcareService using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HealthcareService) LateInitialize(attrs []byte) (bool, error) {
	params := &HealthcareServiceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HealthcareService) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this HealthcareWorkspace
func (mg *HealthcareWorkspace) GetTerraformResourceType() string {
	return "azurerm_healthcare_workspace"
}

// GetConnectionDetailsMapping for this HealthcareWorkspace
func (tr *HealthcareWorkspace) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HealthcareWorkspace
func (tr *HealthcareWorkspace) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HealthcareWorkspace
func (tr *HealthcareWorkspace) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HealthcareWorkspace
func (tr *HealthcareWorkspace) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HealthcareWorkspace
func (tr *HealthcareWorkspace) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HealthcareWorkspace
func (tr *HealthcareWorkspace) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HealthcareWorkspace using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HealthcareWorkspace) LateInitialize(attrs []byte) (bool, error) {
	params := &HealthcareWorkspaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HealthcareWorkspace) GetTerraformSchemaVersion() int {
	return 0
}
