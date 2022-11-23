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

type ScheduleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScheduleParameters struct {

	// The URL of external iCal calendar which override primary events.
	// +kubebuilder:validation:Optional
	IcalURLOverrides *string `json:"icalUrlOverrides,omitempty" tf:"ical_url_overrides,omitempty"`

	// The URL of the external calendar iCal file.
	// +kubebuilder:validation:Optional
	IcalURLPrimary *string `json:"icalUrlPrimary,omitempty" tf:"ical_url_primary,omitempty"`

	// The schedule's name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The list of ID's of on-call shifts.
	// +kubebuilder:validation:Optional
	Shifts []*string `json:"shifts,omitempty" tf:"shifts,omitempty"`

	// The Slack-specific settings for a schedule.
	// +kubebuilder:validation:Optional
	Slack []ScheduleSlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// The schedule's time zone.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// The schedule's type.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ScheduleSlackObservation struct {
}

type ScheduleSlackParameters struct {

	// Slack channel id. Reminder about schedule shifts will be directed to this channel in Slack.
	// +kubebuilder:validation:Optional
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// Slack user group id. Members of user group will be updated when on-call users change.
	// +kubebuilder:validation:Optional
	UserGroupID *string `json:"userGroupId,omitempty" tf:"user_group_id,omitempty"`
}

// ScheduleSpec defines the desired state of Schedule
type ScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduleParameters `json:"forProvider"`
}

// ScheduleStatus defines the observed state of Schedule.
type ScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schedule is the Schema for the Schedules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Schedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduleSpec   `json:"spec"`
	Status            ScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduleList contains a list of Schedules
type ScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schedule `json:"items"`
}

// Repository type metadata.
var (
	Schedule_Kind             = "Schedule"
	Schedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Schedule_Kind}.String()
	Schedule_KindAPIVersion   = Schedule_Kind + "." + CRDGroupVersion.String()
	Schedule_GroupVersionKind = CRDGroupVersion.WithKind(Schedule_Kind)
)

func init() {
	SchemeBuilder.Register(&Schedule{}, &ScheduleList{})
}
