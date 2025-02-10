package v1

import corev1 "k8s.io/api/core/v1"

const (
	UpgradeType_Unset = iota
	UpgradeType_Inplace
	UpgradeType_Canary
)

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

// GatewaySDSConfiguration
type GatewaySDSConfiguration struct {
	Enabled   *bool                        `json:"enabled,omitempty"`
	Image     string                       `json:"image,omitempty"`
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// GatewayConfiguration
type GatewayConfiguration struct {
	Enabled          *bool  `json:"enabled,omitempty"`
	GatewayName      string `json:"gatewayName,omitempty"`
	IstioGatewaySpec `json:",inline"`
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

type ConfigState string
