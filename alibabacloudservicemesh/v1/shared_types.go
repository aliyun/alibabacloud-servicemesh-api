package v1

import corev1 "k8s.io/api/core/v1"

// Reconcile state
type ConfigState string

type ProxyInitResourceLimit struct {
	ResourceCPULimit    string `json:"resourceCPULimit,omitempty"`
	ResourceMemoryLimit string `json:"resourceMemoryLimit,omitempty"`
}

type ProxyInitResourceRequest struct {
	ResourceCPURequest    string `json:"resourceCPURequest,omitempty"`
	ResourceMemoryRequest string `json:"resourceMemoryRequest,omitempty"`
}

type ProxyStatsMatcher struct {
	// Proxy stats name prefix matcher for inclusion.
	InclusionPrefixes []string `json:"inclusionPrefixes,omitempty"`
	// Proxy stats name suffix matcher for inclusion.
	InclusionSuffixes []string `json:"inclusionSuffixes,omitempty"`
	// Proxy stats name regexps matcher for inclusion.
	InclusionRegexps []string `json:"inclusionRegexps,omitempty"`
}

type Int32 struct {
	Value int32 `json:"value,omitempty"`
}

// GatewaySDSConfiguration
type GatewaySDSConfiguration struct {
	Enabled   *bool                        `json:"enabled,omitempty"`
	Image     string                       `json:"image,omitempty"`
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

type SecretVolume struct {
	Name       string `json:"name,omitempty"`
	SecretName string `json:"secretName,omitempty"`
	MountPath  string `json:"mountPath,omitempty"`
}

type ConfigVolume struct {
	Name          string `json:"name,omitempty"`
	ConfigMapName string `json:"configMapName,omitempty"`
	MountPath     string `json:"mountPath,omitempty"`
}

type ResourceCalculationStrategy string

const (
	// calculate resource based on max container resource
	PercentageMaxContainerResource ResourceCalculationStrategy = "PercentageMax"
	// calculate resource based on sum of container resource
	PercentageSumContainerResource ResourceCalculationStrategy = "PercentageSum"
	// calculate resource based on specific container resource
	PercentageRefContainerResource ResourceCalculationStrategy = "PercentageRef"
)

type ScaledSidecarResource struct {
	ResourceCalculationStrategy ResourceCalculationStrategy `json:"resourceCalculationStrategy,omitempty"`
	ContainerRef                string                      `json:"containerRef,omitempty"`
	ResourcePercentage          uint                        `json:"resourcePercentage,omitempty"`
}

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
