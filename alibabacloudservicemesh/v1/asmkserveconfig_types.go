package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ASMKServeConfigState string

const (
	/*
	 * States
	 */
	ASMKServeConfigStateInit       ASMKServeConfigState = "Init"
	ASMKServeConfigStateInprogress ASMKServeConfigState = "Inprogress"
	ASMKServeConfigStateReady      ASMKServeConfigState = "Ready"
	ASMKServeConfigStateFailed     ASMKServeConfigState = "Failed"
)

// ASMKServeConfigSpec defines the desired state of ASMKServeConfig
// +k8s:openapi-gen=true
type ASMKServeConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	KServeConfig `json:",inline"`
}

// ASMKServeConfigStatus defines the observed state of ASMKServeConfig
// +k8s:openapi-gen=true
type ASMKServeConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Status       ConfigState `json:"status,omitempty" protobuf:"bytes,1,rep,name=status"`
	ErrorMessage string      `json:"errorMessage,omitempty" protobuf:"bytes,2,rep,name=errorMessage"`
	CompleteAt   string      `json:"completeAt,omitempty" protobuf:"bytes,3,rep,name=completeAt"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMKServeConfig is the Schema for the asmserviceregistries API
// +genclient
// +k8s:openapi-gen=true
type ASMKServeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMKServeConfigSpec   `json:"spec,omitempty"`
	Status ASMKServeConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMKServeConfigList contains a list of ASMKServeConfig
type ASMKServeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMKServeConfig `json:"items"`
}


type KServeConfig struct {
	Enabled                             bool                    `json:"enabled,omitempty"`
	Tag                                 *string                 `json:"tag,omitempty"`
	CertManagerEnabled                  *bool                   `json:"certManagerEnabled,omitempty"`
	CertManagerUseExisting              *bool                   `json:"certManagerUseExisting,omitempty"`
	CertManagerPreserved                *bool                   `json:"certManagerPreserved,omitempty"` // deprecated，1.17版本后不再使用
	InferenceServiceConfig              *InferenceServiceConfig `json:"inferenceServiceConfig,omitempty"`
	CustomImageRepo                     *CustomImageRepo        `json:"customImageRepo,omitempty"`
	MultiModel                          *bool                   `json:"multiModel,omitempty"`
	SingleModelServerless               *bool                   `json:"singleModelServerless,omitempty"`
	MultiModelControllerNamespace       *string                 `json:"multiModelControllerNamespace,omitempty"`
	BuiltInServingRuntimeEnabled        *bool                   `json:"builtInServingRuntimeEnabled,omitempty"`
	BuiltInClusterServingRuntimeEnabled *bool                   `json:"builtInClusterServingRuntimeEnabled,omitempty"`
	EnableMmCallLocalModelFirst         *bool                   `json:"enableMmCallLocalModelFirst,omitempty"`
}

type InferenceServiceConfig struct {
	Agent              *string `json:"agent,omitempty"`
	Batcher            *string `json:"batcher,omitempty"`
	Credentials        *string `json:"credentials,omitempty"`
	Deploy             *string `json:"deploy,omitempty"`
	Explainers         *string `json:"explainers,omitempty"`
	Predictors         *string `json:"predictors,omitempty"`
	Transformers       *string `json:"transformers,omitempty"`
	Logger             *string `json:"logger,omitempty"`
	StorageInitializer *string `json:"storageInitializer,omitempty"`
	MetricsAggregator  *string `json:"metricsAggregator,omitempty"`
	Router             *string `json:"router,omitempty"`
}

type CustomImageRepo struct {
	Lgbserver         *string `json:"lgbserver,omitempty"`
	Mlserver          *string `json:"mlserver,omitempty"`
	Paddleserver      *string `json:"paddleserver,omitempty"`
	Pmmlserver        *string `json:"pmmlserver,omitempty"`
	Sklearnserver     *string `json:"sklearnserver,omitempty"`
	Tensorflowserving *string `json:"tensorflowserving,omitempty"`
	Torchserve        *string `json:"torchserve,omitempty"`
	Tritonserver      *string `json:"tritonserver,omitempty"`
	Xgbserver         *string `json:"xgbserver,omitempty"`
}