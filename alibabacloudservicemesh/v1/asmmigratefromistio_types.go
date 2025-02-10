package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ASMMigrateFromIstioStste_Init                 = "Init"
	ASMMigrateFromIstioState_SetupIstioForMigrate = "SetupIstioForMigrate"
	ASMMigrateFromIstioState_Migrating            = "Migrating"
	ASMMigrateFromIstioState_Finished             = "Finished"
)

type ASMMigrateFromIstioAdvancedOptions struct {
	StopIstioSystemInjectionDisabling *bool `json:"stopIstioSystemInjectionDisabling,omitempty"`
}

type ASMMigrateFromIstioSpec struct {
	RetryCounter int32 `json:"retryCounter"`
	// +kubebuilder:validation:Enum=Init;SetupIstioForMigrate;Migrating;Finished
	DesiredState    string                              `json:"desiredState"`
	AdvancedOptions *ASMMigrateFromIstioAdvancedOptions `json:"advancedOptions,omitempty"`
}

type ASMMigrateFromIstioStatus struct {
	RetryCounter int32  `json:"retryCounter"`
	State        string `json:"state"`
	Message      string `json:"message"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// ASMMigrateFromIstio is the Schema for the ASMMigrateFromIstio API
// +genclient
type ASMMigrateFromIstio struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ASMMigrateFromIstioSpec   `json:"spec,omitempty"`
	Status            ASMMigrateFromIstioStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// ASMMigrateFromIstioList contains a list of ASMMigrateFromIstio
type ASMMigrateFromIstioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMMigrateFromIstio `json:"items"`
}
