package v1

import (
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IngressType string

type ConfigState string

const (
	IngressTypeK8sIngress   IngressType = "ingress"
	IngressTypeIstioGateway IngressType = "ASM"
)

type RoutingStrategy string

const (
	RoutingStrategy_RuleBased     = "rule_based"
	RoutingStrategy_WeightedBased = "weighted"
)

type ClusterRef struct {
	Region string `json:"region,omitempty"`
	Name   string `json:"name,omitempty"`
	Id     string `json:"id,omitempty"`
}

type ClusterServiceRef struct {
	Name               string        `json:"name,omitempty"`
	Namespace          string        `json:"namespace,omitempty"`
	Cluster            ClusterRef    `json:"cluster,omitempty"`
	TrafficPolicyPatch apiextv1.JSON `json:"trafficPolicyPatch,omitempty"`
	HTTPRoutePatch     apiextv1.JSON `json:"httpRoutePatch,omitempty"`
}

type SwimLaneGatewayConfiguration struct {
	Name      string      `json:"name,omitempty"`
	Namespace string      `json:"namespace,omitempty"`
	Type      IngressType `json:"type,omitempty"`
}

type WeightedSwimLaneIngressConfiguration struct {
	Hosts          []string                      `json:"hosts,omitempty"`
	RequestMatches []SwimLaneIngressRequestMatch `json:"requestMatches,omitempty"`
}

type SwimLaneGroupIngressRouteConfiguration struct {
	WeightedRoutingRule    WeightedSwimLaneIngressConfiguration `json:"weightedRoutingRule,omitempty"`
	IngressRoutingStrategy RoutingStrategy                      `json:"ingressRoutingStrategy,omitempty"`
}

type SwimLaneGroupIngressConfiguration struct {
	Gateway        SwimLaneGatewayConfiguration           `json:"gateway,omitempty"`
	IngressRouting SwimLaneGroupIngressRouteConfiguration `json:"ingressRouting,omitempty"`
}

type PermissiveSwimLaneGroupConfiguration struct {
	// A swimlane's name, indicating a fallback target, namely baseline siwmlane.
	// When target deploy did not exist in the swimlane, the request will be routed to this fallback target.
	FallbackTarget string `json:"fallbackTarget,omitempty"`
	// Newly added in 1.21. A service level fallback map, map from service namespacedName to fallback target swimlane
	// We must make sure that the fallback target version of the service exists.
	// this will override the basic FallbackTarget when a service's fallback target is set.
	ServiceLevelFallbackTarget map[string]string `json:"serviceLevelFallbackTarget,omitempty"`
	// In permissive mode, trace header is used for context propagation in a trace.
	// context propagation could have following scenarios:
	//
	// 1. trace header is a trace id (e.g. x-request-id), it varies for each request. In this case, trace header != route header, trace header is used to propagate the route header, using trafficlabel.
	//
	// 2. trace header is "baggage" (https://www.w3.org/TR/2020/WD-baggage-20201020/). In this case, ASMHeaderPropagation could help for propagate route header using baggage.
	//
	// 3. trace header = route header. This only occurs when user propagate the route header themselves (not a commo case). In this case, we don't need to help propate route header, just use it.
	TraceHeader string `json:"traceHeader,omitempty"`
	// In permissive mode, route header is a header BEING propagted leveaging the context propagation mechanism provided by trace header.
	// It could be any common header such as "version", but not a trace id nor "baggage". We will propagate this header to each call on the calling trace.
	// Then we could just using virtualservice to match this header and route to diffrenct swimlanes.
	RouteHeader string `json:"routeHeader,omitempty"`
}

// ASMSwimLaneGroupSpec defines the desired state of ASMSwimLaneGroup
// +k8s:openapi-gen=true
type ASMSwimLaneGroupSpec struct {
	Services                    []ClusterServiceRef                  `json:"services,omitempty"`
	Ingress                     SwimLaneGroupIngressConfiguration    `json:"ingress,omitempty"`
	PermissiveModeConfiguration PermissiveSwimLaneGroupConfiguration `json:"permissiveModeConfiguration,omitempty"`
	IsPermissive                bool                                 `json:"isPermissive,omitempty"`
	DefaultTrafficPolicyPatch   apiextv1.JSON                        `json:"defaultTrafficPolicyPatch,omitempty"`
	DefaultHTTPRoutePatch       apiextv1.JSON                        `json:"defaultHttpRoutePatch,omitempty"`
	AutoUpdate                  bool                                 `json:"autoUpdate,omitempty"`
}

type ServiceLabelSummary struct {
	ServiceSelector map[string]string            `json:"serviceSelector"`
	PodLabels       map[string]map[string]string `json:"podLabels"`
}

type SwimLaneServices struct {
	ServiceSummaries map[string]*ServiceLabelSummary `json:"serviceSummaries,omitempty"`
	// CommonLabels are podLabels that have potention to be used to distinguish different swimlanes
	CommonLabels map[string][]string `json:"commonLabels,omitempty"`
}

// ASMSwimLaneGroupStatus defines the observed state of ASMSwimLaneGroup
// +k8s:openapi-gen=true
type ASMSwimLaneGroupStatus struct {
	Status                  ConfigState       `json:"status,omitempty"`
	Error                   string            `json:"error,omitempty"`
	ServicesInSwimLaneGroup *SwimLaneServices `json:"servicesInSwimLaneGroup,omitempty"`
}

// ASMSwimLaneGroup is the Schema for the asmswimlanegroups API
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type ASMSwimLaneGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMSwimLaneGroupSpec   `json:"spec,omitempty"`
	Status ASMSwimLaneGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ASMSwimLaneGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMSwimLaneGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMSwimLaneGroup{}, &ASMSwimLaneGroupList{})
}
