package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ASMKnativeConfigSpec defines the desired state of ASMKnativeConfig
type ASMKnativeConfigSpec struct {
	Enabled              bool            `json:"enabled,omitempty"`
	Tag                  *string         `json:"tag,omitempty"`
	UseExisting          *bool           `json:"useExisting,omitempty"`
	DomainMappingEnabled *bool           `json:"domainMappingEnabled,omitempty"`
	DomainConfig         *DomainConfig   `json:"domainConfig,omitempty"`
	Autoscaler           *Autoscaler     `json:"autoscaler,omitempty"`
	DefaultsConfig       *DefaultsConfig `json:"defaults,omitempty"`
	FeaturesConfig       *FeaturesConfig `json:"features,omitempty"`
	IstioConfig          *IstioConfig    `json:"istio,omitempty"`
	NetworkConfig        *NetworkConfig  `json:"network,omitempty"`
}

// DomainConfig defines the custom domain
type DomainConfig struct {
	DomainName     *string `json:"domainName,omitempty"`
	CredentialName *string `json:"credentialName,omitempty"`
}

// DefaultsConfig defines the config of Defaults
type DefaultsConfig struct {
	RevisionTimeoutSeconds          *string `json:"revision-timeout-seconds,omitempty"`
	MaxRevisionTimeoutSeconds       *string `json:"max-revision-timeout-seconds,omitempty"`
	RevisionCpuRequest              *string `json:"revision-cpu-request,omitempty"`
	RevisionMemoryRequest           *string `json:"revision-memory-request,omitempty"`
	RevisionEphemeralStorageRequest *string `json:"revision-ephemeral-storage-request,omitempty"`
	RevisionCpuLimit                *string `json:"revision-cpu-limit,omitempty"`
	RevisionMemoryLimit             *string `json:"revision-memory-limit,omitempty"`
	RevisionEphemeralStorageLimit   *string `json:"revision-ephemeral-storage-limit,omitempty"`
	ContainerNameTemplate           *string `json:"container-name-template,omitempty"`
	ContainerConcurrency            *string `json:"container-concurrency,omitempty"`
	ContainerConcurrencyMaxLimit    *string `json:"container-concurrency-max-limit,omitempty"`
	AllowContainerConcurrencyZero   *string `json:"allow-container-concurrency-zero,omitempty"`
	EnableServiceLinks              *string `json:"enable-service-links,omitempty"`
}

// FeaturesConfig defines the config of multiple features
type FeaturesConfig struct {
	TagHeaderBasedRouting        *string `json:"tag-header-based-routing,omitempty"`
	PodspecVolumesEmptydir       *string `json:"kubernetes.podspec-volumes-emptydir,omitempty"`
	PodspecPersistentVolumeClaim *string `json:"kubernetes.podspec-persistent-volume-claim,omitempty"`
	PodspecPersistentVolumeWrite *string `json:"kubernetes.podspec-persistent-volume-write,omitempty"`
	PodspecInitContainers        *string `json:"kubernetes.podspec-init-containers,omitempty"`
}

// IstioConfig defines the config of Istio
type IstioConfig struct {
	Gatway *string `json:"gateway.knative-serving.knative-ingress-gateway,omitempty"`
	//LocalGatway *string `json:"local-gateway.knative-serving.knative-local-gateway,omitempty"`
}

// NetworkConfig defines the config of network
type NetworkConfig struct {
	HttpProtocol          *string `json:"httpProtocol,omitempty"`
	DefaultExternalScheme *string `json:"defaultExternalScheme,omitempty"`
	RolloutDuration       *string `json:"rolloutDuration,omitempty"`
	DomainTemplate        *string `json:"domainTemplate,omitempty"`
}

// Autoscaler defines the config of AutoScaler
type Autoscaler struct {
	ContainerConcurrencyTargetPercentage *string `json:"container-concurrency-target-percentage,omitempty"`
	ContainerConcurrencyTargetDefault    *string `json:"container-concurrency-target-default,omitempty"`
	RequestsPerSecondTargetDefault       *string `json:"requests-per-second-target-default,omitempty"`
	TargetBurstCapacity                  *string `json:"target-burst-capacity,omitempty"`
	StableWindow                         *string `json:"stable-window,omitempty"`
	PanicWindowPercentage                *string `json:"panic-window-percentage,omitempty"`
	PanicThresholdPercentage             *string `json:"panic-threshold-percentage,omitempty"`
	MaxScaleUpRate                       *string `json:"max-scale-up-rate,omitempty"`
	MaxScaleDownRate                     *string `json:"max-scale-down-rate,omitempty"`
	EnableScaleToZero                    *string `json:"enable-scale-to-zero,omitempty"`
	ScaleToZeroGracePeriod               *string `json:"scale-to-zero-grace-period,omitempty"`
	ScaleToZeroPodRetentionPeriod        *string `json:"scale-to-zero-pod-retention-period,omitempty"`
}

// ASMKnativeConfigStatus defines the observed state of ASMKnativeConfig
type ASMKnativeConfigStatus struct {
	Status       ConfigState `json:"status,omitempty" protobuf:"bytes,1,rep,name=status"`
	ErrorMessage string      `json:"errorMessage,omitempty" protobuf:"bytes,2,rep,name=errorMessage"`
	CompleteAt   string      `json:"completeAt,omitempty" protobuf:"bytes,3,rep,name=completeAt"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMKnativeConfig is the Schema for the asmknativeconfigs API
// +genclient
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=asmknativeconfigs,scope=Namespaced
// KnativeConfig defines the configuration of knative related artifacts
type ASMKnativeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMKnativeConfigSpec   `json:"spec,omitempty"`
	Status ASMKnativeConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMKnativeConfigList contains a list of ASMKnativeConfig
type ASMKnativeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMKnativeConfig `json:"items"`
}

func (r *ASMKnativeConfig) GetDomainConfig() *DomainConfig {
	var domainConfig *DomainConfig
	if r != nil && r.Spec.DomainConfig != nil {
		domainConfig = r.Spec.DomainConfig
	}
	return domainConfig
}

func (r *ASMKnativeConfig) GetDomainName() string {
	domainConfig := r.GetDomainConfig()
	domainName := "example.com"
	if domainConfig != nil {
		if domainConfig.DomainName != nil {
			domainName = *domainConfig.DomainName
		}
	}
	return domainName
}

func (r *ASMKnativeConfig) GetIstioConfig() *IstioConfig {
	var istioConfig *IstioConfig
	if r != nil && r.Spec.IstioConfig != nil {
		istioConfig = r.Spec.IstioConfig
	}
	return istioConfig
}

func (r *ASMKnativeConfig) GetIstioGatewayService() string {
	istioConfig := r.GetIstioConfig()
	gateway := "istio-ingressgateway.istio-system.svc.cluster.local"
	if istioConfig != nil {
		if istioConfig.Gatway != nil && len(*istioConfig.Gatway) > 0 {
			gateway = *istioConfig.Gatway
		}
	}
	return gateway
}