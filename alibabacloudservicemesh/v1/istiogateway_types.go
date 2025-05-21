package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ManagedByASM                    = "asm.alibabacloud.com/managed-by-asm"
	LOG_PROJECT_FOR_MANAGED_GATEWAY = "asm.alibabacloud.com/aliyun-logs-project"
	REUSED_LOGTAIL_CONFIG           = "asm.alibabacloud.com/aliyun-logtail-config"
	REUSED_MACHINE_GROUP            = "asm.alibabacloud.com/aliyun-logs-machinegroup"
	REPLICAS_MANAGED_BY_ASM         = "asm.alibabacloud.com/replicas-managed-by-asm"
	ENABLE_CANARY_DEPLOYMENT        = "asm.alibabacloud.com/enable-canary-deployment"
)

type GatewayType string

const (
	GatewayTypeIngress      GatewayType = "ingress"
	GatewayTypeEgress       GatewayType = "egress"
	GatewayTypeCrossNetwork GatewayType = "cross-network"
)

// IstioGatewaySpec defines the desired state of Istio
// +k8s:openapi-gen=true
type IstioGatewaySpec struct {
	// +kubebuilder:validation:Enum=ingress;egress
	GatewayType                    GatewayType `json:"gatewayType,omitempty"`
	IstioGatewayBasicConfiguration `json:",inline"`
	ClusterIds                     []string                                  `json:"clusterIds,omitempty"`
	clusterId                      string                                    `json:"clusterId,omitempty"`
	Overrides                      map[string]IstioGatewayBasicConfiguration `json:"overrides,omitempty"`
}

type KernelParameters struct {
	NetCoreSoMaxConn        *string `json:"net.core.somaxconn,omitempty"`
	NetCoreNetdevMaxBacklog *string `json:"net.core.netdev_max_backlog,omitempty"`
	NetIpv4TcpRMem          *string `json:"net.ipv4.tcp_rmem,omitempty"`
	NetIpv4TcpWMem          *string `json:"net.ipv4.tcp_wmem,omitempty"`
	// https://sauravomar01.medium.com/how-to-increase-the-port-range-in-linux-b7df2f0b0b58
	NetIpv4IpLocalPortRange      *string `json:"net.ipv4.ip_local_port_range,omitempty"`
	NetIpv4TcpFinTimeout         *string `json:"net.ipv4.tcp_fin_timeout,omitempty"`
	NetIpv4TcpTwTimeout          *string `json:"net.ipv4.tcp_tw_timeout,omitempty"`
	NetIpv4TcpTwReuse            *string `json:"net.ipv4.tcp_tw_reuse,omitempty"`
	NetIpv4TcpTwRecycle          *string `json:"net.ipv4.tcp_tw_recycle,omitempty"`
	NetIpv4TcpTimestamps         *string `json:"net.ipv4.tcp_timestamps,omitempty"`
	NetIpv4TcpRetries2           *string `json:"net.ipv4.tcp_retries2,omitempty"`
	NetIpv4TcpSlowStartAfterIdle *string `json:"net.ipv4.tcp_slow_start_after_idle,omitempty"`
	NetIpv4TcpMaxOrphans         *string `json:"net.ipv4.tcp_max_orphans,omitempty"`
	NetIpv4TcpMaxSynBacklog      *string `json:"net.ipv4.tcp_max_syn_backlog,omitempty"`
	NetIpv4TcpNoMetricsSave      *string `json:"net.ipv4.tcp_no_metrics_save,omitempty"`
	NetIpv4TcpAutocorking        *string `json:"net.ipv4.tcp_autocorking,omitempty"`
	KernelPrintk                 *string `json:"kernel.printk,omitempty"`
	VmSwappiness                 *string `json:"vm.swappiness,omitempty"`
}

type KernelConfiguration struct {
	Enabled    *bool            `json:"enabled,omitempty"`
	Parameters KernelParameters `json:"parameters,omitempty"`
}

type GzipCompressorConfiguration struct {
	MemoryLevel         *uint32 `json:"memory_level,omitempty"`
	CompressionLevel    *string `json:"compression_level,omitempty"`
	CompressionStrategy *string `json:"compression_strategy,omitempty"`
	WindowBits          *uint32 `json:"window_bits,omitempty"`
	ChunkSize           *uint32 `json:"chunk_size,omitempty"`
}

type CompressorConfiguration struct {
	Enabled                    *bool                        `json:"enabled,omitempty"`
	MinContentLength           *uint32                      `json:"min_content_length,omitempty"`
	ContentType                *[]string                    `json:"content_type,omitempty"`
	DisableOnETagHeader        *bool                        `json:"disable_on_etag_header,omitempty"`
	RemoveAcceptEncodingHeader *bool                        `json:"remove_accept_encoding_header,omitempty"`
	Gzip                       *GzipCompressorConfiguration `json:"gzip,omitempty"`
}

type ReadinessProbeConfiguration struct {
	FailureThreshold int32 `json:"failureThreshold,omitempty"`
	PeriodSeconds    int32 `json:"periodSeconds,omitempty"`
	SuccessThreshold int32 `json:"successThreshold,omitempty"`
}

// IstioGatewayBasicConfiguration defines the basic configuration of gateway service
// +k8s:openapi-gen=true
type IstioGatewayBasicConfiguration struct {
	ReplicaCount     int32 `json:"replicaCount,omitempty"`
	AutoscaleEnabled bool  `json:"autoscaleEnabled,omitempty"`
	MinReplicas      int32 `json:"minReplicas,omitempty"`
	MaxReplicas      int32 `json:"maxReplicas,omitempty"`
	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer
	ServiceType           corev1.ServiceType `json:"serviceType,omitempty"`
	LoadBalancerClass     *string            `json:"loadBalancerClass,omitempty"`
	LoadBalancerIP        string             `json:"loadBalancerIP,omitempty"`
	ExternalTrafficPolicy string             `json:"externalTrafficPolicy,omitempty"`
	ServiceAnnotations    map[string]string  `json:"serviceAnnotations,omitempty"`
	PodAnnotations        map[string]string  `json:"podAnnotations,omitempty"`
	ServiceLabels         map[string]string  `json:"serviceLabels,omitempty"`
	// begin to add Pod Labels
	PodLabels map[string]string `json:"podLabels,omitempty"`
	// end to add Pod Labels
	SDS                   GatewaySDSConfiguration      `json:"sds,omitempty"`
	Resources             *corev1.ResourceRequirements `json:"resources,omitempty"`
	EnvVars               []corev1.EnvVar              `json:"env,omitempty"`
	Ports                 []corev1.ServicePort         `json:"ports,omitempty"`
	MeshExpansionPorts    []corev1.ServicePort         `json:"meshExpansionPorts,omitempty"`
	NodeSelector          map[string]string            `json:"nodeSelector,omitempty"`
	ApplicationPorts      string                       `json:"applicationPorts,omitempty"`
	RequestedNetworkView  string                       `json:"requestedNetworkView,omitempty"`
	Affinity              *corev1.Affinity             `json:"affinity,omitempty"`
	Tolerations           []corev1.Toleration          `json:"tolerations,omitempty"`
	SecretVolumes         []SecretVolume               `json:"secretVolumes,omitempty"`
	ConfigVolumes         []ConfigVolume               `json:"configVolumes,omitempty"`
	RollingMaxSurge       *string                      `json:"rollingMaxSurge,omitempty"`
	RollingMaxUnavailable *string                      `json:"rollingMaxUnavailable,omitempty"`
	CPU                   HpaCPU                       `json:"cpu,omitempty"`
	Memory                HpaMemory                    `json:"memory,omitempty"`
	Kernel                KernelConfiguration          `json:"kernel,omitempty"`
	Compressor            CompressorConfiguration      `json:"compression,omitempty"`
	// Whether to run the gateway in a privileged container
	RunAsRoot *bool `json:"runAsRoot,omitempty"`
	// LifeCycle
	Lifecycle *corev1.Lifecycle `json:"lifecycle,omitempty"`
	// ReadinessProbe
	ReadinessProbe              ReadinessProbeConfiguration `json:"readinessProbe,omitempty"`
	HostNetwork                 *bool                       `json:"hostNetwork,omitempty"`
	DnsPolicy                   corev1.DNSPolicy            `json:"dnsPolicy,omitempty"`
	AutoCreateGatewayYaml       *bool                       `json:"autoCreateGatewayYaml,omitempty"`
	DisableContainerPortExposed *bool                       `json:"disableContainerPortExposed,omitempty"`
}

// HPACPU defines the cpu metrics of hpa
// +k8s:openapi-gen=true
type HpaCPU struct {
	TargetAverageUtilization *int32 `json:"targetAverageUtilization,omitempty"`
}

// HPACPU defines the memory metrics of hpa
// +k8s:openapi-gen=true
type HpaMemory struct {
	TargetAverageUtilization *int32 `json:"targetAverageUtilization,omitempty"`
}

// IstioGatewayStatus defines the observed state of IstioGateway
// +k8s:openapi-gen=true
type IstioGatewayStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Status                     ConfigState      `json:"Status,omitempty"`
	GatewayAddress             []string         `json:"GatewayAddress,omitempty"`
	ClusterIdGatewayAddressMap []GatewayAddress `json:"ClusterIdGatewayAddressMap,omitempty"`

	ErrorMessage string `json:"ErrorMessage,omitempty"`
}

// GatewayAddress defines the ipaddress and clusterId pair
// +k8s:openapi-gen=true
type GatewayAddress struct {
	IPAddress string `json:"ipAddress,omitempty"`
	ClusterId string `json:"clusterId,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IstioGateway is the Schema for the istiogateways API
// +genclient
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type IstioGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IstioGatewaySpec   `json:"spec,omitempty"`
	Status IstioGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IstioGatewayList contains a list of IstioGateway
type IstioGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IstioGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IstioGateway{}, &IstioGatewayList{})
}
