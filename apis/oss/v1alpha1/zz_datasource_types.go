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

type DataSourceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataSourceParameters struct {

	// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
	// +kubebuilder:validation:Optional
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// Whether to enable basic auth for the data source. Defaults to `false`.
	// +kubebuilder:validation:Optional
	BasicAuthEnabled *bool `json:"basicAuthEnabled,omitempty" tf:"basic_auth_enabled,omitempty"`

	// Basic auth username. Defaults to “.
	// +kubebuilder:validation:Optional
	BasicAuthUsername *string `json:"basicAuthUsername,omitempty" tf:"basic_auth_username,omitempty"`

	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to “.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Serialized JSON string containing the json data. Replaces the json_data attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	// +kubebuilder:validation:Optional
	JSONDataEncoded *string `json:"jsonDataEncoded,omitempty" tf:"json_data_encoded,omitempty"`

	// A unique name for the data source.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Serialized JSON string containing the secure json data. Replaces the secure_json_data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	// +kubebuilder:validation:Optional
	SecureJSONDataEncodedSecretRef *v1.SecretKeySelector `json:"secureJsonDataEncodedSecretRef,omitempty" tf:"-"`

	// The data source type. Must be one of the supported data source keywords.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Unique identifier. If unset, this will be automatically generated.
	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// (Required by some data source types) The username to use to authenticate to the data source. Defaults to “.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// DataSourceSpec defines the desired state of DataSource
type DataSourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSourceParameters `json:"forProvider"`
}

// DataSourceStatus defines the observed state of DataSource.
type DataSourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSource is the Schema for the DataSources API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type DataSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSourceSpec   `json:"spec"`
	Status            DataSourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSourceList contains a list of DataSources
type DataSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSource `json:"items"`
}

// Repository type metadata.
var (
	DataSource_Kind             = "DataSource"
	DataSource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSource_Kind}.String()
	DataSource_KindAPIVersion   = DataSource_Kind + "." + CRDGroupVersion.String()
	DataSource_GroupVersionKind = CRDGroupVersion.WithKind(DataSource_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSource{}, &DataSourceList{})
}
