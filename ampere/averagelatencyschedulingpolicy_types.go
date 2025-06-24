// +kubebuilder:object:generate=true
// +groupName=istio.alibabacloud.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AverageLatencySchedulingPolicyStatus defines the observed state of Policy.
type AverageLatencySchedulingPolicyStatus struct {
	Status       string `json:"status,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:storageversion
// Policy is the Schema for the policies API.
// +genclient
type AverageLatencySchedulingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	//+kubebuilder:pruning:PreserveUnknownFields
	Spec   AverageLatencySchedulingPolicySpec   `json:"spec,omitempty"`
	Status AverageLatencySchedulingPolicyStatus `json:"status,omitempty"`
}

type Extractor struct {
	Address       *AddressExtractor    `json:"address,omitempty"`
	From          *string              `json:"from,omitempty"`
	Json          *JSONExtractor       `json:"json,omitempty"`
	JWT           *JWTExtractor        `json:"jwt,omitempty"`
	PathTemplates *PathTemplateMatcher `json:"path_templates,omitempty"`
}

type AddressExtractor struct {
	From *string `json:"from,omitempty"`
}

type JSONExtractor struct {
	From    *string `json:"from,omitempty"`
	Pointer *string `json:"pointer,omitempty"`
}
type JWTExtractor struct {
	From        *string `json:"from,omitempty"`
	JsonPointer *string `json:"json_pointer,omitempty"`
}

type PathTemplateMatcher struct {
	TemplateValues *map[string]string `json:"template_values,omitempty"`
}

type LabelMatcher struct {
	Expression *Expression `json:"expression,omitempty"`
	//MatchExpressions []MatchRequirement `json:"match_expressions,omitempty"`
	MatchLabels *map[string]string  `json:"match_labels,omitempty"`
	MatchList   *[]MatchRequirement `json:"match_list,omitempty"`
}

// +kubebuilder:pruning:PreserveUnknownFields
type ExpressionList struct {
	Of *[]Expression `json:"of,omitempty"`
}

type Expression struct {
	All          *ExpressionList    `json:"all,omitempty"`
	Any          *ExpressionList    `json:"any,omitempty"`
	LabelEquals  *EqualsExpression  `json:"label_equals,omitempty"`
	LabelExists  *string            `json:"label_exists,omitempty"`
	LabelMatches *MatchesExpression `json:"label_matches,omitempty"`
	Not          *NotExpression     `json:"not,omitempty"`
}
type NotExpression struct {
	All          *ExpressionList    `json:"all,omitempty"`
	Any          *ExpressionList    `json:"any,omitempty"`
	LabelEquals  *EqualsExpression  `json:"label_equals,omitempty"`
	LabelExists  *string            `json:"label_exists,omitempty"`
	LabelMatches *MatchesExpression `json:"label_matches,omitempty"`
}

type EqualsExpression struct {
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
}

type MatchesExpression struct {
	Label *string `json:"label,omitempty"`
	Regex *string `json:"regex,omitempty"`
}

type MatchRequirement struct {
	Key      *string   `json:"key,omitempty"`
	Operator *string   `json:"operator,omitempty"`
	Values   *[]string `json:"values,omitempty"`
}
type Selector struct {
	AgentGroup   *string       `json:"agent_group,omitempty"`
	ControlPoint *string       `json:"control_point,omitempty"`
	LabelMatcher *LabelMatcher `json:"label_matcher,omitempty"`
	Service      *string       `json:"service,omitempty"`
}

type Rule struct {
	Extractor *Extractor `json:"extractor,omitempty"`
}

type Classifier struct {
	Rules     *map[string]Rule `json:"rules,omitempty"`
	Selectors *[]Selector      `json:"selectors,omitempty"`
}

type FlowControlResources struct {
	Classifiers *[]Classifier `json:"classifiers,omitempty"`
}

type Resources struct {
	FlowControl *FlowControlResources `json:"flow_control,omitempty"`
}

type GradientControllerParameters struct {
	MaxGradient *float64 `json:"max_gradient,omitempty"`
	MinGradient *float64 `json:"min_gradient,omitempty"`
	Slope       *float64 `json:"slope,omitempty"`
}

type Scheduler struct {
	DecisionDeadlineMargin    *string                      `json:"decision_deadline_margin,omitempty"`
	DefaultWorkloadParameters *SchedulerWorkloadParameters `json:"default_workload_parameters,omitempty"`
	DeniedResponseStatusCode  *string                      `json:"denied_response_status_code,omitempty"`
	FairnessLabelKey          *string                      `json:"fairness_label_key,omitempty"`
	PriorityLabelKey          *string                      `json:"priority_label_key,omitempty"`
	TokensLabelKey            *string                      `json:"tokens_label_key,omitempty"`
	WorkloadLabelKey          *string                      `json:"workload_label_key,omitempty"`
	Workloads                 *[]SchedulerWorkload         `json:"workloads,omitempty"`
}

type StatusCode struct {
}
type SchedulerWorkload struct {
	LabelMatcher *LabelMatcher                `json:"label_matcher,omitempty"`
	Name         *string                      `json:"name,omitempty"`
	Parameters   *SchedulerWorkloadParameters `json:"parameters,omitempty"`
}
type SchedulerWorkloadParameters struct {
	Priority     *float64 `json:"priority,omitempty"`
	QueueTimeout *string  `json:"queue_timeout,omitempty"`
	Tokens       *float64 `json:"tokens,omitempty"`
}
type LoadSchedulerParameters struct {
	Scheduler                  *Scheduler     `json:"scheduler,omitempty"`
	Selectors                  *[]Selector    `json:"selectors,omitempty"`
	MatchRequests              []MatchRequest `json:"match_requests,omitempty"`
	WorkloadLatencyBasedTokens *bool          `json:"workload_latency_based_tokens,omitempty"`
}

type AimdLoadScheduler struct {
	Gradient                      *GradientControllerParameters `json:"gradient,omitempty"`
	LoadMultiplierLinearIncrement *float64                      `json:"load_multiplier_linear_increment,omitempty"`
	LoadScheduler                 *LoadSchedulerParameters      `json:"load_scheduler,omitempty"`
	MaxLoadMultiplier             *float64                      `json:"max_load_multiplier,omitempty"`
}

type LoadSchedulingCore struct {
	AimdLoadScheduler *AimdLoadScheduler `json:"aimd_load_scheduler,omitempty"`
}

type FluxMeterExponentialBuckets struct {
	Count  *int32   `json:"count,omitempty"`
	Factor *float64 `json:"factor,omitempty"`
	Start  *float64 `json:"start,omitempty"`
}

type FluxMeterExponentialBucketsRange struct {
	Count *int32   `json:"count,omitempty"`
	Max   *float64 `json:"max,omitempty"`
	Min   *float64 `json:"min,omitempty"`
}

type FluxMeterLinearBuckets struct {
	Count *int32   `json:"count,omitempty"`
	Width *float64 `json:"width,omitempty"`
	Start *float64 `json:"start,omitempty"`
}

type FluxMeterStaticBuckets struct {
	Buckets *[]float64 `json:"buckets,omitempty"`
}

type Meter struct {
	AttributeKey            *string                           `json:"attribute_key,omitempty"`
	ExponentialBuckets      *FluxMeterExponentialBuckets      `json:"exponential_buckets,omitempty"`
	ExponentialBucketsRange *FluxMeterExponentialBucketsRange `json:"exponential_buckets_range,omitempty"`
	LinearBuckets           *FluxMeterLinearBuckets           `json:"linear_buckets,omitempty"`
	Selectors               *[]Selector                       `json:"selectors,omitempty"`
	MatchRequests           []MatchRequest                    `json:"match_requests,omitempty"`
	StaticBuckets           *FluxMeterStaticBuckets           `json:"static_buckets,omitempty"`
}

type LatencyBaseliner struct {
	LatencyToleranceMultiplier    *float64 `json:"latency_tolerance_multiplier,omitempty"`
	LongTermQueryInterval         *string  `json:"long_term_query_interval,omitempty"`
	LongTermQueryPeriodicInterval *string  `json:"long_term_query_periodic_interval,omitempty"`
	Meter                         *Meter   `json:"meter,omitempty"`
}

type AverageLatencySchedulingPolicySpec struct {
	Resources          *Resources          `json:"resources,omitempty"`
	LoadSchedulingCore *LoadSchedulingCore `json:"load_scheduling_core,omitempty"`
	LatencyBaseliner   *LatencyBaseliner   `json:"latency_baseliner,omitempty"`
	WorkloadSelector   *WorkloadSelector   `json:"workload_selector,omitempty"`
	ApplyToTraffic     *ApplyToTraffic     `json:"apply_to_traffic,omitempty"`
}

//+kubebuilder:object:root=true

// PolicyList contains a list of Policy.
type AverageLatencySchedulingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AverageLatencySchedulingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AverageLatencySchedulingPolicy{}, &AverageLatencySchedulingPolicyList{})
}
