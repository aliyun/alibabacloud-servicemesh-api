package v1

import (
	networking "istio.io/api/networking/v1alpha3"
)

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
