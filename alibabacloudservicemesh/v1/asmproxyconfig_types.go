package v1

import (
	"github.com/gogo/protobuf/proto"
	"istio.io/api/type/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProxyTracingCustomTag_Named struct {
	Name         *string `json:"name,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
}

type ProxyTracingCustomTag_Literal struct {
	Value *string `json:"value,omitempty"`
}

type ProxyTracingCustomTag struct {
	Literal     *ProxyTracingCustomTag_Literal `json:"literal,omitempty"`
	Header      *ProxyTracingCustomTag_Named   `json:"header,omitempty"`
	Environment *ProxyTracingCustomTag_Named   `json:"environment,omitempty"`
}

type ProxyTracingZipkin struct {
	Address *string `json:"address,omitempty"`
}

type ProxyTracingConfiguration struct {
	Sampling         *float32                         `json:"sampling,omitempty"`
	CustomTags       map[string]ProxyTracingCustomTag `json:"custom_tags,omitempty"`
	MaxPathTagLength *uint32                          `json:"max_path_tag_length,omitempty"`
	Zipkin           *ProxyTracingZipkin              `json:"zipkin,omitempty"`
}

type CryptoMbConfiguration struct {
	PollDelay *string `json:"pollDelay,omitempty"`
	Fallback  *bool   `json:"fallback,omitempty"`
}

type PrivateKeyProviderConfiguration struct {
	CryptoMb *CryptoMbConfiguration `json:"cryptomb,omitempty"`
}

type SMCConfiguration struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// ASMProxyConfigSpec defines the desired state of ASMProxyConfig
// +k8s:openapi-gen=true
type ASMProxyConfigSpec struct {
	// Optional. Selectors specify the set of pods/VMs on which this `ProxyConfig` resource should be applied.
	// If not set, the `ProxyConfig` resource will be applied to all workloads in the namespace where this resource is defined.
	Selector *v1beta1.WorkloadSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The number of worker threads to run.
	// If unset, defaults to 2. If set to 0, this will be configured to use all cores on the machine using
	// CPU requests and limits to choose a value, with limits taking precedence over requests.
	// +k8s:openapi-gen=true
	Concurrency *int32 `protobuf:"bytes,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	//TerminationDrainDuration
	TerminationDrainDuration *string `protobuf:"bytes,3,opt,name=terminationDrainDuration,proto3" json:"terminationDrainDuration,omitempty"`
	//Lifecycle describes actions that the management system should take in response to container lifecycle events
	Lifecycle *corev1.Lifecycle `protobuf:"bytes,4,opt,name=lifecycle,proto3"  json:"lifecycle,omitempty"`
	//SidecarProxyInitResourceLimit describes the cpu/memory limit for init container of sidecar proxy
	SidecarProxyInitResource *corev1.ResourceRequirements `protobuf:"bytes,5,opt,name=sidecarProxyInitResource,proto3"   json:"sidecarProxyInitResource,omitempty"`
	SidecarProxyResource     *corev1.ResourceRequirements `protobuf:"bytes,6,opt,name=sidecarProxyResource,proto3"   json:"sidecarProxyResource,omitempty"`

	//ExcludeOutboundPorts: A comma separated list of outbound ports to be excluded from redirection to Envoy.
	ExcludeOutboundPorts *string `protobuf:"bytes,7,opt,name=excludeOutboundPorts,proto3" json:"excludeOutboundPorts,omitempty"`
	//ExcludeOutboundIPRanges: A comma separated list of IP ranges in CIDR form to be excluded from redirection. Only applies when all outbound traffic (i.e. '*') is being redirected.
	ExcludeOutboundIPRanges *string `protobuf:"bytes,8,opt,name=excludeOutboundIPRanges,proto3" json:"excludeOutboundIPRanges,omitempty"`
	//IncludeOutboundIPRanges: A comma separated list of IP ranges in CIDR form to redirect to Envoy (optional). The wildcard character '*' can be used to redirect all outbound traffic. An empty list will disable all outbound redirection.
	IncludeOutboundIPRanges *string `protobuf:"bytes,9,opt,name=includeOutboundIPRanges,proto3" json:"includeOutboundIPRanges,omitempty"`
	//ExcludeInboundPorts: A comma separated list of inbound ports to be excluded from redirection to Envoy. Only applies when all inbound traffic (i.e. '*') is being redirected.
	ExcludeInboundPorts *string `protobuf:"bytes,10,opt,name=excludeInboundPorts,proto3"  json:"excludeInboundPorts,omitempty"`
	//IncludeInboundPorts: A comma separated list of inbound ports for which traffic is to be redirected to Envoy. The wildcard character '*' can be used to configure redirection for all ports.
	//An empty list will disable all inbound redirection.
	IncludeInboundPorts *string `protobuf:"bytes,11,opt,name=includeInboundPorts,proto3" json:"includeInboundPorts,omitempty"`
	//IncludeOutboundPorts: A comma separated list of outbound ports for which traffic is to be redirected to Envoy, regardless of the destination IP.
	IncludeOutboundPorts *string `protobuf:"bytes,12,opt,name=includeOutboundPorts,proto3" json:"includeOutboundPorts,omitempty"`
	IstioDNSProxyEnabled *bool   `protobuf:"bytes,13,opt,name=istioDNSProxyEnabled,proto3"  json:"istioDNSProxyEnabled,omitempty"`
	//envoy metric
	ProxyStatsMatcher               *ProxyStatsMatcher               `protobuf:"bytes,14,opt,name=proxyStatsMatcher,proto3"  json:"proxyStatsMatcher,omitempty"  yaml:"proxyStatsMatcher,omitempty"`
	HoldApplicationUntilProxyStarts *bool                            `protobuf:"bytes,15,opt,name=holdApplicationUntilProxyStarts,proto3" json:"holdApplicationUntilProxyStarts,omitempty"  yaml:"holdApplicationUntilProxyStarts,omitempty"`
	LogLevel                        *string                          `protobuf:"bytes,16,opt,name=logLevel,proto3" json:"logLevel,omitempty"  yaml:"logLevel,omitempty"`
	Tracing                         *ProxyTracingConfiguration       `protobuf:"bytes,17,opt,name=tracing,proto3" json:"tracing,omitempty" yaml:"tracing,omitempty"`
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

func (in *ASMProxyConfigSpec) Reset() {
	*in = ASMProxyConfigSpec{}
}

func (in *ASMProxyConfigSpec) String() string {
	return proto.CompactTextString(in)
}

func (in *ASMProxyConfigSpec) ProtoMessage() {}

func (m *ASMProxyConfigSpec) GetSelector() *v1beta1.WorkloadSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

// ProxyImage The following values are used to construct proxy image url.
// $hub/$image_name/$tag-$image_type
// example: docker.io/istio/proxyv2:1.11.1 or docker.io/istio/proxyv2:1.11.1-distroless
// This information was previously part of the Values API.
// +k8s:openapi-gen=true
type ProxyImage struct {
	// The image type of the image.
	// Istio publishes default, debug, and distroless images.
	// Other values are allowed if those image types (example: centos) are published to the specified hub.
	// supported values: default, debug, distroless.
	ImageType string `protobuf:"bytes,1,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
}

// ASMProxyConfigStatus defines the observed state of ASMProxyConfig
// +k8s:openapi-gen=true
type ASMProxyConfigStatus struct {
}

// ASMProxyConfig is the Schema for the asmcustomproxyconfig API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type ASMProxyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMProxyConfigSpec   `json:"spec,omitempty"`
	Status ASMProxyConfigStatus `json:"status,omitempty"`
}

// ASMProxyConfigList contains a list of ASMProxyConfig
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ASMProxyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMProxyConfig `json:"items"`
}
