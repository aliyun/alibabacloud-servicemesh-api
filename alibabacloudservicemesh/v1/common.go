package v1

import (
	networking "istio.io/api/networking/v1alpha3"
	corev1 "k8s.io/api/core/v1"
)

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

type RouteConfigurationMatch struct {
	PortNumber *uint32                 `protobuf:"varint,0,opt,name=port_number,json=portNumber,proto3" json:"portNumber,omitempty"`
	PortName   *string                 `protobuf:"bytes,1,opt,name=port_name,json=portName,proto3" json:"portName,omitempty"`
	Gateway    *string                 `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Vhost      *CommonVirtualHostMatch `protobuf:"bytes,3,opt,name=vhost,proto3" json:"vhost,omitempty"`
	Name       *string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

type CommonVirtualHostMatch struct {
	Name  *string           `protobuf:"bytes,0,opt,name=name,proto3" json:"name,omitempty"`
	Route *CommonRouteMatch `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
}

type CommonRouteMatch struct {
	Name   *string `protobuf:"bytes,0,opt,name=name,proto3" json:"name,omitempty"`
	Action *string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
}

func (r *RouteConfigurationMatch) ToIstioRouteMatch() *networking.EnvoyFilter_EnvoyConfigObjectMatch_RouteConfiguration {

	routeMatch := &networking.EnvoyFilter_EnvoyConfigObjectMatch_RouteConfiguration{
		RouteConfiguration: &networking.EnvoyFilter_RouteConfigurationMatch{},
	}
	if r.PortNumber != nil {
		routeMatch.RouteConfiguration.PortNumber = *r.PortNumber
	}
	if r.PortName != nil {
		routeMatch.RouteConfiguration.PortName = *r.PortName
	}
	if r.Gateway != nil {
		routeMatch.RouteConfiguration.Gateway = *r.Gateway
	}
	if r.Vhost != nil {
		routeMatch.RouteConfiguration.Vhost = &networking.EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch{}
		if r.Vhost.Name != nil {
			routeMatch.RouteConfiguration.Vhost.Name = *r.Vhost.Name
		}
		if r.Vhost.Route != nil {
			routeMatch.RouteConfiguration.Vhost.Route = &networking.EnvoyFilter_RouteConfigurationMatch_RouteMatch{}
			if r.Vhost.Route.Name != nil {
				routeMatch.RouteConfiguration.Vhost.Route.Name = *r.Vhost.Route.Name
			}
			if r.Vhost.Route.Action != nil {
				switch *r.Vhost.Route.Action {
				case "ROUTE":
					{
						routeMatch.RouteConfiguration.Vhost.Route.Action = networking.EnvoyFilter_RouteConfigurationMatch_RouteMatch_ROUTE
					}
				case "REDIRECT":
					{
						routeMatch.RouteConfiguration.Vhost.Route.Action = networking.EnvoyFilter_RouteConfigurationMatch_RouteMatch_REDIRECT
					}
				case "DIRECT_RESPONSE":
					{
						routeMatch.RouteConfiguration.Vhost.Route.Action = networking.EnvoyFilter_RouteConfigurationMatch_RouteMatch_REDIRECT
					}
				default:
					{
						routeMatch.RouteConfiguration.Vhost.Route.Action = networking.EnvoyFilter_RouteConfigurationMatch_RouteMatch_ANY
					}
				}
			}
		}
	}
	return routeMatch
}
