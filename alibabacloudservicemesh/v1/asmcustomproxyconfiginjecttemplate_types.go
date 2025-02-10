package v1

import (
	"istio.io/api/type/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ASMCustomProxyConfigInjectionTemplateSpec defines the desired state of ASMCustomProxyConfigInjectionTemplate
// +k8s:openapi-gen=true
type ASMCustomProxyConfigInjectionTemplateSpec struct {
	// Criteria used to select the specific set of pods/VMs on which this proxy Config should be applied.
	// If omitted, the proxy Config will be applied to all workload instances in the same namespace.
	WorkloadSelector *v1beta1.WorkloadSelector `json:"workloadSelector,omitempty"`
	ConfigPatches    ConfigPatches             `json:"configPatches,omitempty"`
}
type ConfigPatches struct {
	//TerminationDrainDuration
	TerminationDrainDuration *string `json:"terminationDrainDuration,omitempty" yaml:"terminationDrainDuration,omitempty"`
	//Lifecycle describes actions that the management system should take in response to container lifecycle events
	Lifecycle *corev1.Lifecycle `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
	//SidecarProxyInitResourceLimit describes the cpu/memory limit for init container of sidecar proxy
	SidecarProxyInitResourceLimit *ProxyInitResourceLimit `json:"sidecarProxyInitResourceLimit,omitempty" yaml:"sidecarProxyInitResourceLimit,omitempty"`
	//ProxyInitResourceRequest describes the cpu/memory request for init container of sidecar proxy
	SidecarProxyInitResourceRequest *ProxyInitResourceRequest `json:"sidecarProxyInitResourceRequest,omitempty" yaml:"sidecarProxyInitResourceRequest,omitempty"`
	//SidecarProxyResourceLimit describes the cpu/memory limit for sidecar proxy
	SidecarProxyResourceLimit *ProxyInitResourceLimit `json:"sidecarProxyResourceLimit,omitempty" yaml:"sidecarProxyResourceLimit,omitempty"`
	//SidecarProxyResourceRequest describes the cpu/memory request for sidecar proxy
	SidecarProxyResourceRequest *ProxyInitResourceRequest `json:"sidecarProxyResourceRequest,omitempty" yaml:"sidecarProxyResourceRequest,omitempty"`

	//ExcludeOutboundPorts: A comma separated list of outbound ports to be excluded from redirection to Envoy.
	ExcludeOutboundPorts *string `json:"excludeOutboundPorts,omitempty" yaml:"excludeOutboundPorts,omitempty"`
	//ExcludeOutboundIPRanges: A comma separated list of IP ranges in CIDR form to be excluded from redirection. Only applies when all outbound traffic (i.e. '*') is being redirected.
	ExcludeOutboundIPRanges *string `json:"excludeOutboundIPRanges,omitempty" yaml:"excludeOutboundIPRanges,omitempty"`
	//IncludeOutboundIPRanges: A comma separated list of IP ranges in CIDR form to redirect to Envoy (optional). The wildcard character '*' can be used to redirect all outbound traffic. An empty list will disable all outbound redirection.
	IncludeOutboundIPRanges *string `json:"includeOutboundIPRanges,omitempty" yaml:"includeOutboundIPRanges,omitempty"`
	//ExcludeInboundPorts: A comma separated list of inbound ports to be excluded from redirection to Envoy. Only applies when all inbound traffic (i.e. '*') is being redirected.
	ExcludeInboundPorts *string `json:"excludeInboundPorts,omitempty" yaml:"excludeInboundPorts,omitempty"`
	//IncludeInboundPorts: A comma separated list of inbound ports for which traffic is to be redirected to Envoy. The wildcard character '*' can be used to configure redirection for all ports.
	//An empty list will disable all inbound redirection.
	IncludeInboundPorts *string `json:"includeInboundPorts,omitempty" yaml:"includeInboundPorts,omitempty"`
	//IncludeOutboundPorts: A comma separated list of outbound ports for which traffic is to be redirected to Envoy, regardless of the destination IP.
	IncludeOutboundPorts *string `json:"includeOutboundPorts,omitempty" yaml:"includeOutboundPorts,omitempty"`
	IstioDNSProxyEnabled *bool   `json:"istioDNSProxyEnabled,omitempty" yaml:"istioDNSProxyEnabled,omitempty"`
	Concurrency          *int32  `json:"concurrency,omitempty" yaml:"concurrency,omitempty"`
	//envoy metric
	ProxyStatsMatcher               *ProxyStatsMatcher               `json:"proxyStatsMatcher,omitempty"  yaml:"proxyStatsMatcher,omitempty"`
	HoldApplicationUntilProxyStarts *bool                            `json:"holdApplicationUntilProxyStarts,omitempty"  yaml:"holdApplicationUntilProxyStarts,omitempty"`
	LogLevel                        *string                          `json:"logLevel,omitempty"  yaml:"logLevel,omitempty"`
	Tracing                         *ProxyTracingConfiguration       `json:"tracing,omitempty" yaml:"tracing,omitempty"`
	InterceptionMode                *string                          `json:"interceptionMode,omitempty" yaml:"interceptionMode,omitempty"`
	DrainDuration                   *string                          `json:"drainDuration,omitempty" yaml:"drainDuration,omitempty"`
	ParentShutdownDuration          *string                          `json:"parentShutdownDuration,omitempty" yaml:"parentShutdownDuration,omitempty"`
	ProxyMetadata                   map[string]string                `json:"proxyMetadata,omitempty" yaml:"proxyMetadata,omitempty"`
	ExtraStatTags                   []string                         `json:"extraStatTags,omitempty" yaml:"extraStatTags,omitempty"`
	PrivateKeyProvider              *PrivateKeyProviderConfiguration `json:"privateKeyProvider,omitempty" yaml:"privateKeyProvider,omitempty"`
	Privileged                      *bool                            `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	EnableCoreDump                  *bool                            `json:"enableCoreDump,omitempty" yaml:"enableCoreDump,omitempty"`
	ReadinessInitialDelaySeconds    *int32                           `json:"readinessInitialDelaySeconds,omitempty" yaml:"readinessInitialDelaySeconds,omitempty"`
	ReadinessPeriodSeconds          *int32                           `json:"readinessPeriodSeconds,omitempty" yaml:"readinessPeriodSeconds,omitempty"`
	ReadinessFailureThreshold       *int32                           `json:"readinessFailureThreshold,omitempty" yaml:"readinessFailureThreshold,omitempty"`

	SidecarProxyInitAckSloResource *corev1.ResourceRequirements `json:"sidecarProxyInitAckSloResource,omitempty" yaml:"sidecarProxyInitAckSloResource"`
	SidecarProxyAckSloResource     *corev1.ResourceRequirements `json:"sidecarProxyAckSloResource,omitempty" yaml:"sidecarProxyAckSloResource"`
	SMC                            *SMCConfiguration            `json:"smcConfiguration,omitempty" yaml:"smcConfiguration,omitempty"`
	RuntimeValues                  map[string]string            `json:"runtimeValues,omitempty" yaml:"runtimeValues,omitempty"`
}

// ASMCustomProxyConfigInjectionTemplateStatus defines the observed state of ASMCustomProxyConfigInjectionTemplate
// +k8s:openapi-gen=true
type ASMCustomProxyConfigInjectionTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Status       ConfigState `json:"Status"`
	ErrorMessage string      `json:"ErrorMessage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMCustomProxyConfigInjectionTemplate is the Schema for the asmcustomproxyconfiginjectiontemplate API
// +genclient
// +k8s:openapi-gen=true
type ASMCustomProxyConfigInjectionTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMCustomProxyConfigInjectionTemplateSpec   `json:"spec,omitempty"`
	Status ASMCustomProxyConfigInjectionTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMCustomProxyConfigInjectionTemplateList contains a list of ASMCustomProxyConfigInjectionTemplate
type ASMCustomProxyConfigInjectionTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMCustomProxyConfigInjectionTemplate `json:"items"`
}