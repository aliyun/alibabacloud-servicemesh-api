package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ASMSecurityPolicySpec defines the desired state of ASMSecurityPolicy
type ASMSecurityPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Oidc            *OIDCConfig      `json:"oidcConfig,omitempty"`
	JwtConfig       *JWTConfig       `json:"jwtConfig,omitempty"`
	ASMExtAuthzHttp *ASMExtAuthzHttp `json:"asmExtAuthzHttp,omitempty"`
	ASMExtAuthzGrpc *ASMExtAuthzGrpc `json:"asmExtAuthzGrpc,omitempty"`
	SimpleAuthz     *SimpleAuthz     `json:"simpleAuthz,omitempty"`
}

type SimpleAuthz struct {
	WorkloadGroups []WorkloadGroup `json:"workloadGroups,omitempty"`
}

type ASMExtAuthzHttp struct {
	ExistedExtProviderName *string                                 `json:"existedExtProviderName,omitempty"`
	EnvoyExtAuthzHttp      *EnvoyExternalAuthorizationHttpProvider `json:"envoyExtAuthzHttp,omitempty"`
	WorkloadGroups         []WorkloadGroup                         `json:"workloadGroups,omitempty"`
}

type ASMExtAuthzGrpc struct {
	ExistedExtProviderName *string                                 `json:"existedExtProviderName,omitempty"`
	EnvoyExtAuthzGrpc      *EnvoyExternalAuthorizationGrpcProvider `json:"envoyExtAuthzGrpc,omitempty"`
	WorkloadGroups         []WorkloadGroup                         `json:"workloadGroups,omitempty"`
}

type WorkloadGroup struct {
	Name             string                  `json:"name,omitempty"`
	MatchMode        string                  `json:"matchMode,omitempty"`
	WorkloadSelector []ASMSPWorkloadSelector `json:"workloadSelector,omitempty"`
	RequestSelector  []RequestSelector       `json:"requestSelector,omitempty"`
}

type RequestSelector struct {
	Host          *string `json:"host,omitempty"`
	Port          *string `json:"port,omitempty"`
	Method        *string `json:"method,omitempty"`
	Path          *string `json:"path,omitempty"`
	IpBlock       *string `json:"ipBlock,omitempty"`
	RemoteIpBlock *string `json:"remoteIpBlock,omitempty"`
}

type OIDCConfig struct {
	ExistedExtProviderName *string         `json:"existedExtProviderName,omitempty"`
	RedirectProtocol       *string         `json:"redirectProtocol,omitempty"`
	RedirectASMGateway     *string         `json:"redirectASMGateway,omitempty"`
	RedirectASMGatewayIP   *string         `json:"redirectASMGatewayIP,omitempty"`
	RedirectDomain         *string         `json:"redirectDomain,omitempty"`
	RedirectPort           *uint32         `json:"redirectPort,omitempty"`
	IssuerURI              *string         `json:"issuerURI,omitempty"`
	ClientID               *string         `json:"clientID,omitempty"`
	ClientSecret           *string         `json:"clientSecret,omitempty"`
	CookieSecret           *string         `json:"cookieSecret,omitempty"`
	CookieRefresh          *string         `json:"cookieRefresh,omitempty"`
	CookieExpire           *string         `json:"cookieExpire,omitempty"`
	Scopes                 *[]string       `json:"scopes,omitempty"`
	SignoutEndpoint        *string         `json:"signoutEndpoint,omitempty"`
	WorkloadGroups         []WorkloadGroup `json:"workloadGroups,omitempty"`
}

type JWTConfig struct {
	JWTRules       []JWTRule       `json:"jwtRules,omitempty"`
	WorkloadGroups []WorkloadGroup `json:"workloadGroups,omitempty"`
}

type ASMSPWorkloadSelector struct {
	Name      string            `json:"name,omitempty"`
	Namespace string            `json:"namespace,omitempty"`
	Kind      string            `json:"kind,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
}

type JWTRule struct {
	Issuer                *string     `json:"issuer,omitempty"`
	Jwks                  *string     `json:"jwks,omitempty"`
	JwksUri               *string     `json:"jwksUri,omitempty"`
	Audiences             []string    `json:"audiences,omitempty"`
	FromHeaders           []JWTHeader `json:"fromHeaders,omitempty"`
	FromParams            []string    `json:"fromParams,omitempty"`
	OutputPayloadToHeader *string     `json:"outputPayloadToHeader,omitempty"`
	ForwardOriginalToken  *bool       `json:"forwardOriginalToken,omitempty"`
}

type JWTHeader struct {
	Name   string `json:"name,omitempty"`
	Prefix string `json:"prefix,omitempty"`
}

// ASMSecurityPolicyStatus defines the observed state of ASMSecurityPolicy
type ASMSecurityPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Status       ConfigState   `json:"status,omitempty"`
	SubRes       []SubResource `json:"subResource,omitempty"`
	ErrorMessage string        `json:"errormessage,omitempty"`
}

type SubResource struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMSecurityPolicy is the Schema for the asmsecurepolicies API
// +genclient
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=asmsecuritypolicies,scope=Cluster
type ASMSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMSecurityPolicySpec   `json:"spec,omitempty"`
	Status ASMSecurityPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMSecurityPolicyList contains a list of ASMSecurityPolicy
type ASMSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMSecurityPolicy `json:"items"`
}
