package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EnvoyExternalAuthorizationRequestBody struct {
	MaxRequestBytes *uint32 `json:"maxRequestBytes,omitempty"`
	// 如果body太大，就只发部分body给外部授权服务
	AllowPartialMessage *bool `json:"allowPartialMessage,omitempty"`
	PackAsBytes         *bool `json:"packAsBytes,omitempty"`
}

type EnvoyExternalAuthorizationBase struct {
	Service                   *string                                `json:"service,omitempty"`
	Port                      *uint32                                `json:"port,omitempty"`
	Timeout                   *string                                `json:"timeout,omitempty"`
	FailOpen                  *bool                                  `json:"failOpen,omitempty"`
	StatusOnError             *string                                `json:"statusOnError,omitempty"`
	IncludeRequestBodyInCheck *EnvoyExternalAuthorizationRequestBody `json:"includeRequestBodyInCheck,omitempty"`
}

type ASMOIDC struct {
	RedirectProtocol     *string                  `json:"redirectProtocol,omitempty"`
	RedirectASMGateway   *string                  `json:"redirectASMGateway,omitempty"`
	RedirectASMGatewayIP *string                  `json:"redirectASMGatewayIP,omitempty"`
	RedirectPort         *uint32                  `json:"redirectPort,omitempty"`
	RedirectDomain       *string                  `json:"redirectDomain,omitempty"`
	IssuerURI            *string                  `json:"issuerURI,omitempty"`
	ClientID             *string                  `json:"clientID,omitempty"`
	ClientSecret         *string                  `json:"clientSecret,omitempty"`
	CookieSecret         *string                  `json:"cookieSecret,omitempty"`
	CookieRefresh        *string                  `json:"cookieRefresh,omitempty"`
	CookieExpire         *string                  `json:"cookieExpire,omitempty"`
	Scopes               *[]string                `json:"scopes,omitempty"`
	ImageTag             *string                  `json:"imageTag,omitempty"`
	Replicas             *int32                   `json:"replicas,omitempty"`
	SignoutEndpoint      *string                  `json:"signoutEndpoint,omitempty"`
	ResourceRequirements *v1.ResourceRequirements `json:"resourceRequirements,omitempty"`
}

type EnvoyExternalAuthorizationHttpProvider struct {
	EnvoyExternalAuthorizationBase  `json:",inline"`
	PathPrefix                      *string            `json:"pathPrefix,omitempty"`
	IncludeRequestHeadersInCheck    *[]string          `json:"includeRequestHeadersInCheck,omitempty"`
	IncludeAdditionalHeadersInCheck *map[string]string `json:"includeAdditionalHeadersInCheck,omitempty"`
	HeadersToUpstreamOnAllow        *[]string          `json:"headersToUpstreamOnAllow,omitempty"`
	HeadersToDownstreamOnDeny       *[]string          `json:"headersToDownstreamOnDeny,omitempty"`

	OIDC *ASMOIDC `json:"oidc,omitempty"`
}
type EnvoyExternalAuthorizationGrpcProvider struct {
	EnvoyExternalAuthorizationBase `json:",inline"`
}

type ZipkinTracingProvider struct {
	Service      *string `json:"service,omitempty"`
	Path         *string `json:"path,omitempty"`
	Port         *uint32 `json:"port,omitempty"`
	MaxTagLength *uint32 `json:"maxTagLength,omitempty"`
}

type LightstepTracingProvider struct {
	Service      *string `json:"service,omitempty"`
	Port         *uint32 `json:"port,omitempty"`
	AccessToken  *string `json:"accessToken,omitempty"`
	MaxTagLength *uint32 `json:"maxTagLength,omitempty"`
}

type DatadogTracingProvider struct {
	Service      *string `json:"service,omitempty"`
	Port         *uint32 `json:"port,omitempty"`
	MaxTagLength *uint32 `json:"maxTagLength,omitempty"`
}

type StackdriverProvider_Logging struct {
	Labels *map[string]string `json:"labels,omitempty"`
}

type StackdriverProvider struct {
	MaxTagLength *uint32                      `json:"maxTagLength,omitempty"`
	Logging      *StackdriverProvider_Logging `json:"logging,omitempty"`
}

type OpenCensusAgentTracingProvider_TraceContext string

type OpenCensusAgentTracingProvider struct {
	Service      *string                                       `json:"service,omitempty"`
	Port         *uint32                                       `json:"port,omitempty"`
	Context      []OpenCensusAgentTracingProvider_TraceContext `json:"context,omitempty"`
	MaxTagLength *uint32                                       `json:"maxTagLength,omitempty"`
}

type SkywalkingTracingProvider struct {
	Service     *string `json:"service,omitempty"`
	Port        *uint32 `json:"port,omitempty"`
	AccessToken *string `json:"accessToken,omitempty"`
}

type OpenTelemetryProvider struct {
	Service      *string                            `json:"service,omitempty"`
	Port         *uint32                            `json:"port,omitempty"`
	MaxTagLength *uint32                            `json:"maxTagLength,omitempty"`
	Grpc         *OpenTelemetryProvider_GrpcService `json:"grpc,omitempty"`
	Http         *OpenTelemetryProvider_HttpService `json:"http,omitempty"`
}

type OpenTelemetryProvider_HttpService struct {
	// REQUIRED. Specifies the path on the service.
	Path *string `json:"path,omitempty"`
	// Optional. Specifies the timeout for the HTTP request.
	Timeout *string `json:"timeout,omitempty"`
	// Optional. Allows specifying custom HTTP headers that will be added
	// to each HTTP request sent.
	Headers []*HeaderValue `json:"headers,omitempty"`
}

type OpenTelemetryProvider_GrpcService struct {
	// Optional. Specifies the timeout for the GRPC request.
	Timeout *string `json:"timeout,omitempty"`
	// Optional. Additional metadata to include in streams initiated to the GrpcService. This can be used for
	// scenarios in which additional ad hoc authorization headers (e.g. “x-foo-bar: baz-key“) are to
	// be injected.
	InitialMetadata []*HeaderValue `json:"initialMetadata,omitempty"`
}

type HeaderValue struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type EnvoyFileAccessLogProvider_LogFormat struct {
	Text   *string            `json:"text,omitempty"`
	Labels *map[string]string `json:"lables,omitempty"`
}

type EnvoyFileAccessLogProvider struct {
	Path      *string                               `json:"path,omitempty"`
	LogFormat *EnvoyFileAccessLogProvider_LogFormat `json:"logFormat,omitempty"`
}

type EnvoyHttpGrpcV3LogProvider struct {
	Service                         *string   `json:"service,omitempty"`
	Port                            *uint32   `json:"port,omitempty"`
	LogName                         *string   `json:"logName,omitempty"`
	FilterStateObjectsToLog         *[]string `json:"filterStateObjectsToLog,omitempty"`
	AdditionalRequestHeadersToLog   *[]string `json:"additionalRequestHeadersToLog,omitempty"`
	AdditionalResponseHeadersToLog  *[]string `json:"additionalResponseHeadersToLog,omitempty"`
	AdditionalResponseTrailersToLog *[]string `json:"additionalResponseTrailersToLog,omitempty"`
}

type EnvoyTcpGrpcV3LogProvider struct {
	Service                 *string   `json:"service,omitempty"`
	Port                    *uint32   `json:"port,omitempty"`
	LogName                 *string   `json:"logName,omitempty"`
	FilterStateObjectsToLog *[]string `json:"filterStateObjectsToLog,omitempty"`
}

type EnvoyOpenTelemetryLogProvider struct {
	Service   *string                               `json:"service,omitempty"`
	Port      *uint32                               `json:"port,omitempty"`
	LogName   *string                               `json:"logName,omitempty"`
	LogFormat *EnvoyFileAccessLogProvider_LogFormat `json:"logFormat,omitempty"`
}

type AggregateProvider struct {
	DefaultProvider *ProviderRef      `json:"defaultProvider,omitempty"`
	UniqueProviders []*UniqueProvider `json:"uniqueProviders,omitempty"`
}

type ProviderRef struct {
	Name *string `json:"name,omitempty"`
}

type UniqueProvider struct {
	Provider *ProviderRef `json:"provider,omitempty"`
	Clusters []ClusterRef `json:"clusters,omitempty"`
}

type ASMExtensionProviderSpec struct {
	EnvoyExtAuthzHttp  *EnvoyExternalAuthorizationHttpProvider `json:"envoyExtAuthzHttp,omitempty"`
	EnvoyExtAuthzGrpc  *EnvoyExternalAuthorizationGrpcProvider `json:"envoyExtAuthzGrpc,omitempty"`
	Zipkin             *ZipkinTracingProvider                  `json:"zipkin,omitempty"`
	Lightstep          *LightstepTracingProvider               `json:"lightstep,omitempty"`
	Datadog            *DatadogTracingProvider                 `json:"datadog,omitempty"`
	Stackdriver        *StackdriverProvider                    `json:"stackdriver,omitempty"`
	Opencensus         *OpenCensusAgentTracingProvider         `json:"opencensus,omitempty"`
	Skywalking         *SkywalkingTracingProvider              `json:"skywalking,omitempty"`
	OpenTelemetry      *OpenTelemetryProvider                  `json:"opentelemetry,omitempty"`
	EnvoyFileAccessLog *EnvoyFileAccessLogProvider             `json:"envoyFileAccessLog,omitempty"`
	EnvoyHttpAls       *EnvoyHttpGrpcV3LogProvider             `json:"envoyHttpAls,omitempty"`
	EnvoyTcpAls        *EnvoyTcpGrpcV3LogProvider              `json:"envoyTcpAls,omitempty"`
	EnvoyOtelAls       *EnvoyOpenTelemetryLogProvider          `json:"envoyOtelAls,omitempty"`
	Aggregate          *AggregateProvider                      `json:"aggregate,omitempty"`
}

type ASMExtensionProviderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Status defines the state of this instance
	Status string `json:"status,omitempty"`
	// Message defines the possible error message
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// ASMExtensionProvider is the Schema for the asmextensionprovider API
// +genclient
type ASMExtensionProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMExtensionProviderSpec   `json:"spec,omitempty"`
	Status ASMExtensionProviderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ASMExtensionProviderList contains a list of ASMExtensionProvider
type ASMExtensionProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []*ASMExtensionProvider `json:"items"`
}