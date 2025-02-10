package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ASMCustomMetricSpec defines the desired state of ASMCustomMetric
type ASMCustomMetricSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	WorkloadSelector *WorkloadSelector `json:"workloadSelector,omitempty"`
	ConfigPatch      *ConfigPatch      `json:"configPatch,omitempty"`
}

type ConfigPatch struct {
	InboundSidecar  *PluginConfig `json:"inboundSidecar,omitempty"`
	OutboundSidecar *PluginConfig `json:"outboundSidecar,omitempty"`
	Gateway         *PluginConfig `json:"gateway,omitempty"`
}

type PluginConfig struct {
	Debug bool `json:"debug,omitempty"`

	// prefix to add to stats emitted by the plugin.
	// DEPRECATED.
	StatPrefix string `json:"stat_prefix,omitempty"`

	// Optional: Disable using host header as a fallback if destination service is
	// not available from the controlplane. Disable the fallback if the host
	// header originates outsides the mesh, like at ingress.
	DisableHostHeaderFallback bool `json:"disable_host_header_fallback,omitempty"`

	Metrics []MetricConfig `json:"metrics,omitempty"`
}

type MetricConfig struct {
	// (Optional) Collection of tag names and tag expressions to include in the
	// metric. Conflicts are resolved by the tag name by overriding previously
	// supplied values.
	Dimensions map[string]string `json:"dimensions,omitempty"`

	// (Optional) Metric name to restrict the override to a metric. If not
	// specified, applies to all.
	Name string `json:"name,omitempty"`

	// (Optional) A list of tags to remove.
	TagsToRemove []string `json:"tags_to_remove,omitempty"`

	// NOT IMPLEMENTED. (Optional) Conditional enabling the override.
	Match string `json:"match,omitempty"`

	// (Optional) If this is set to true, the metric(s) selected by this
	// configuration will not be generated or reported.
	Drop bool `json:"drop,omitempty"`
}

// ASMCustomMetricStatus defines the observed state of ASMCustomMetric
type ASMCustomMetricStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status       CustomMetricState `json:"Status"`
	ErrorMessage string            `json:"ErrorMessage"`
	CompleteAt   string            `json:"CompleteAt"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMCustomMetric is the Schema for the ASMCustomMetrics API
// +genclient
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=asmcustommetrics,scope=Namespaced
type ASMCustomMetric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMCustomMetricSpec   `json:"spec,omitempty"`
	Status ASMCustomMetricStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMCustomMetricList contains a list of ASMCustomMetric
type ASMCustomMetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMCustomMetric `json:"items"`
}


type CustomMetricState string

const (
	CustomMetricInit       CustomMetricState = "Init"
	CustomMetricInprogress CustomMetricState = "Inprogress"
	CustomMetricReady      CustomMetricState = "Ready"
	CustomMetricFailed     CustomMetricState = "Failed"
	TelemetryV2Disabled    CustomMetricState = "TelemetryV2Disabled"
)