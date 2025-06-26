package v1

import (
	"fmt"
	"sort"
	"strings"
)

type ApplyToTraffic string

func (a *ApplyToTraffic) String() string {
	if a == nil {
		return string(InboundTraffic)
	}
	return string(*a)
}

var (
	// SIDECAR INBOUND TRAFFIC
	// InboundTraffic defines the rate limiter configuration for inbound traffic.
	//
	// +kubebuilder:validation:Optional
	InboundTraffic ApplyToTraffic = "sidecar_inbound"

	// SIDECAR OUTBOUND TRAFFIC
	// vhost: tracing-analysis-dc-hz-internal.aliyuncs.com:8090
	// cluster: outbound|8090||tracing-analysis-dc-hz-internal.aliyuncs.com
	//
	// +kubebuilder:validation:Optional
	OutboundTraffic ApplyToTraffic = "sidecar_outbound"

	// GATEWAY INBOUND TRAFFIC
	//
	// +kubebuilder:validation:Optional
	GatewayTraffic ApplyToTraffic = "gateway"
)

type TargetServiceRef struct {
	// Group is the group of the referent. For example, "gateway.networking.k8s.io".
	// When unspecified or empty string, core API group is inferred.
	//
	// +optional
	// +kubebuilder:default=""
	Group *Group `json:"group,omitempty"`

	// Kind is the Kubernetes resource kind of the referent. For example
	// "Service".
	//
	// Defaults to "Service" when not specified.
	//
	// +optional
	// +kubebuilder:default=Service
	Kind *Kind `json:"kind,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Port *PortNumber `json:"port,omitempty"`

	// +kubebuilder:validation:Optional
	SectionName *SectionName `json:"section_name,omitempty"`
}

type MatchRequest struct {
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty"`

	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty"`

	// +kubebuilder:validation:Optional
	Headers map[string]string `json:"headers,omitempty"`

	// +kubebuilder:validation:Optional
	TargetService *TargetServiceRef `json:"target_service,omitempty"`
}

// +kubebuilder:validation:Minimum=1
// +kubebuilder:validation:Maximum=65535
type PortNumber int32

type WorkloadSelector struct {
	// One or more labels that indicate a specific set of pods/VMs
	// on which the configuration should be applied. The scope of
	// label search is restricted to the configuration namespace in which the
	// the resource is present.
	//
	// +kubebuilder:validation:Required
	Labels map[string]string `json:"labels,omitempty"`
}

func (w *WorkloadSelector) EqualsTo(to *WorkloadSelector) bool {
	if w.IsWildcard() || to.IsWildcard() {
		return w.IsWildcard() == to.IsWildcard()
	}

	if len(w.Labels) != len(to.Labels) {
		return false
	}

	for k, v := range w.Labels {
		if to.Labels[k] != v {
			return false
		}
	}

	return true
}

func (w *WorkloadSelector) IsWildcard() bool {
	return w == nil || len(w.Labels) == 0
}

func (w *WorkloadSelector) String() string {
	if w == nil || len(w.Labels) == 0 {
		return "*"
	}

	keys := make([]string, 0, len(w.Labels))
	for k := range w.Labels {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	str := ""
	for _, k := range keys {
		str = str + "," + fmt.Sprintf("%s=%s", k, w.Labels[k])
	}
	return strings.TrimPrefix(str, ",")
}

// Constants and structures based on the described logic
type Constants struct {
	MetricScrapeInterval      string
	CircuitEvaluationInterval string
}

type Driver struct {
	QueryString string
	Criteria    Criteria
	Selectors   []string
	Percentile  float64
	FluxMeter   FluxMeter
}

type Condition struct {
	Operator  string
	Threshold float64
}

type FluxMeter struct {
	Name      string
	Selectors *[]Selector `json:"selectors,omitempty"`
}

type Port struct {
	SignalName     *string         `json:"signal_name,omitempty"`
	ConstantSignal *ConstantSignal `json:"constant_signal,omitempty"`
}

type Query struct {
	Promql PromQL `json:"promql,omitempty"`
}

type PromQL struct {
	QueryString        string          `json:"query_string,omitempty"`
	EvaluationInterval string          `json:"evaluation_interval,omitempty"`
	OutPorts           map[string]Port `json:"out_ports,omitempty"`
}

type Decider struct {
	Operator string          `json:"operator,omitempty"`
	InPorts  map[string]Port `json:"in_ports,omitempty"`
	OutPorts map[string]Port `json:"out_ports,omitempty"`
}

type ConstantSignal struct {
	SpecialValue string  `json:"special_value,omitempty"`
	Value        float64 `json:"value,omitempty"`
}

type InPort struct {
	ConstantSignal *ConstantSignal `json:"constant_signal,omitempty"`
	SignalName     *string         `json:"signal_name,omitempty"`
}

type Component struct {
	Query        *Query                    `json:"query,omitempty"`
	Decider      *Decider                  `json:"decider,omitempty"`
	BoolVariable *BoolVariableComponent    `json:"bool_variable,omitempty"`
	OrVariable   *LogicalORComponent       `json:"or,omitempty"`
	AndVariable  *LogicalAndComponent      `json:"and,omitempty"`
	FirstValid   *FirstValidComponent      `json:"first_valid,omitempty"`
	Inverter     *LogicalInverterComponent `json:"inverter,omitempty"`
	FlowControl  *FlowControlComponent     `json:"flow_control,omitempty"`
}

type OutPort struct {
	SignalName string `json:"signal_name,omitempty"`
}

type Ins struct {
	Output OutPort `json:"output,omitempty"`
}

type BoolVariableComponent struct {
	ConfigKey      string           `json:"config_key,omitempty"`
	ConstantOutput bool             `json:"constant_output"`
	OutPorts       BoolVariableOuts `json:"out_ports,omitempty"`
}
type BoolVariableOuts struct {
	Output OutPort `json:"output,omitempty"`
}

type LogicalORComponent struct {
	InPorts  OrIns  `json:"in_ports,omitempty"`
	OutPorts OrOuts `json:"out_ports,omitempty"`
}
type OrIns struct {
	Inputs []InPort `json:"inputs,omitempty"`
}

type OrOuts struct {
	Output OutPort `json:"output,omitempty"`
}

type LogicalInverterComponent struct {
	InPorts  InverterIns  `json:"in_ports,omitempty"`
	OutPorts InverterOuts `json:"out_ports,omitempty"`
}
type InverterIns struct {
	Input InPort `json:"input,omitempty"`
}

type InverterOuts struct {
	Output OutPort `json:"output,omitempty"`
}

type FirstValidComponent struct {
	InPorts  FirstValidIns  `json:"in_ports,omitempty"`
	OutPorts FirstValidOuts `json:"out_ports,omitempty"`
}
type FirstValidIns struct {
	Inputs []InPort `json:"inputs,omitempty"`
}

type FirstValidOuts struct {
	Output OutPort `json:"output,omitempty"`
}

type LogicalAndComponent struct {
	InPorts  AndIns  `json:"in_ports,omitempty"`
	OutPorts AndOuts `json:"out_ports,omitempty"`
}
type AndIns struct {
	Inputs []InPort `json:"inputs,omitempty"`
}
type AndOuts struct {
	Output OutPort `json:"output,omitempty"`
}

type FlowControlComponent struct {
	LoadRamp *LoadRampComponent `json:"load_ramp,omitempty"`
}

type LoadRampComponent struct {
	InPorts                         *LoadRampIns        `json:"in_ports,omitempty"`
	OutPorts                        *LoadRampOuts       `json:"out_ports,omitempty"`
	Parameters                      *LoadRampParameters `json:"parameters,omitempty"`
	PassThroughLabelValues          *[]string           `json:"pass_through_label_values,omitempty"`
	PassThroughLabelValuesConfigKey string              `json:"pass_through_label_values_config_key,omitempty"`
}
type LoadRampIns struct {
	Backward *InPort `json:"backward,omitempty"`
	Forward  *InPort `json:"forward,omitempty"`
	Reset    *InPort `json:"reset,omitempty"`
}
type LoadRampOuts struct {
	AcceptPercentage *OutPort `json:"accept_percentage,omitempty"`
	AtEnd            *OutPort `json:"at_end,omitempty"`
	AtStart          *OutPort `json:"at_start,omitempty"`
}
type LoadRampParameters struct {
	Sampler *Sampler `json:"sampler,omitempty"`
	Steps   *[]Step  `json:"steps,omitempty"`
}

type DriverAccumulator struct {
	FluxMeters                   map[string]FluxMeter
	ForwardSignals               []string
	BackwardSignals              []string
	ResetSignals                 []string
	ForwardSignalsCount          int
	BackwardSignalsCount         int
	ResetSignalsCount            int
	Components                   []Component
	PromQLDriverCount            int
	AverageLatencyDriverCount    int
	PercentileLatencyDriverCount int
}

// Group refers to a Kubernetes Group. It must either be an empty string or a
// RFC 1123 subdomain.
//
// This validation is based off of the corresponding Kubernetes validation:
// https://github.com/kubernetes/apimachinery/blob/02cfb53916346d085a6c6c7c66f882e3c6b0eca6/pkg/util/validation/validation.go#L208
//
// Valid values include:
//
// * "" - empty string implies core Kubernetes API group
// * "gateway.networking.k8s.io"
// * "foo.example.com"
//
// Invalid values include:
//
// * "example.com/bar" - "/" is an invalid character
//
// +kubebuilder:validation:MaxLength=253
// +kubebuilder:validation:Pattern=`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
type Group string

// Kind refers to a Kubernetes Kind.
//
// Valid values include:
//
// * "Service"
// * "HTTPRoute"
//
// Invalid values include:
//
// * "invalid/kind" - "/" is an invalid character
//
// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=63
// +kubebuilder:validation:Pattern=`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`
type Kind string

// SectionName is the name of a section in a Kubernetes resource.
//
// In the following resources, SectionName is interpreted as the following:
//
// * Gateway: Listener name
// * HTTPRoute: HTTPRouteRule name
// * Service: Port name
//
// Section names can have a variety of forms, including RFC 1123 subdomains,
// RFC 1123 labels, or RFC 1035 labels.
//
// This validation is based off of the corresponding Kubernetes validation:
// https://github.com/kubernetes/apimachinery/blob/02cfb53916346d085a6c6c7c66f882e3c6b0eca6/pkg/util/validation/validation.go#L208
//
// Valid values include:
//
// * "example"
// * "foo-example"
// * "example.com"
// * "foo.example.com"
//
// Invalid values include:
//
// * "example.com/bar" - "/" is an invalid character
//
// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=253
type SectionName string
