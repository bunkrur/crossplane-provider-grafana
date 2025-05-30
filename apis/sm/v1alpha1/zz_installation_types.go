/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstallationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Generated token to access the SM API.
	SmAccessToken *string `json:"smAccessToken,omitempty" tf:"sm_access_token,omitempty"`
}

type InstallationParameters struct {

	// Reference to a Stack in cloud to populate stackId.
	// +kubebuilder:validation:Optional
	CloudStackRef *v1.Reference `json:"cloudStackRef,omitempty" tf:"-"`

	// Selector for a Stack in cloud to populate stackId.
	// +kubebuilder:validation:Optional
	CloudStackSelector *v1.Selector `json:"cloudStackSelector,omitempty" tf:"-"`

	// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
	// +kubebuilder:validation:Required
	MetricsPublisherKeySecretRef v1.SecretKeySelector `json:"metricsPublisherKeySecretRef" tf:"-"`

	// The ID or slug of the stack to install SM on.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1.Stack
	// +crossplane:generate:reference:refFieldName=CloudStackRef
	// +crossplane:generate:reference:selectorFieldName=CloudStackSelector
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/synthetic-monitoring/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	// +kubebuilder:validation:Optional
	StackSmAPIURL *string `json:"stackSmApiUrl,omitempty" tf:"stack_sm_api_url,omitempty"`
}

// InstallationSpec defines the desired state of Installation
type InstallationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstallationParameters `json:"forProvider"`
}

// InstallationStatus defines the observed state of Installation.
type InstallationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstallationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Installation is the Schema for the Installations API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Installation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstallationSpec   `json:"spec"`
	Status            InstallationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstallationList contains a list of Installations
type InstallationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Installation `json:"items"`
}

// Repository type metadata.
var (
	Installation_Kind             = "Installation"
	Installation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Installation_Kind}.String()
	Installation_KindAPIVersion   = Installation_Kind + "." + CRDGroupVersion.String()
	Installation_GroupVersionKind = CRDGroupVersion.WithKind(Installation_Kind)
)

func init() {
	SchemeBuilder.Register(&Installation{}, &InstallationList{})
}
