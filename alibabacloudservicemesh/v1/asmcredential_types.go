package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	TYPE_TLS                         string      = "tls"
	TYPE_MTLS                        string      = "mtls"
	ASMCredentialStatusAboutToExpire ConfigState = "AboutToExpire"
	ASMCredentialStatusExpired       ConfigState = "Expired"
)

// ASMCredentialSpec defines the desired state of ASMCredential
type ASMCredentialSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Type   *string `json:"type,omitempty"`
	Cert   *string `json:"cert,omitempty"`
	Key    *string `json:"key,omitempty"`
	CaCert *string `json:"caCert,omitempty"`
}

// ASMCredentialStatus defines the observed state of ASMCredential
type ASMCredentialStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Status       ConfigState            `json:"status,omitempty"`
	ErrorMessage string                 `json:"errormessage,omitempty"`
	CertInfo     *ASMCredentialCertInfo `json:"certInfo,omitempty"`
}

type ASMCredentialCertInfo struct {
	SubjectInfo        *SubjectInfo `json:"subjectInfo,omitempty"`
	IssuerInfo         *IssuerInfo  `json:"issuerInfo,omitempty"`
	SerialNumber       *string      `json:"serialNumber,omitempty"`
	PublicKeyAlgorithm *string      `json:"publicKeyAlgorithm,omitempty"`
	SignatureAlgorithm *string      `json:"signatureAlgorithm,omitempty"`
	NotBefore          *string      `json:"notBefore,omitempty"`
	NotAfter           *string      `json:"notAfter,omitempty"`
	FingerPrint        *string      `json:"fingerPrint,omitempty"`
}

type SubjectInfo struct {
	CommonName       *string   `json:"commonName,omitempty"`
	Country          *[]string `json:"country,omitempty"`
	Province         *[]string `json:"province,omitempty"`
	Locality         *[]string `json:"locality,omitempty"`
	Organization     *[]string `json:"organization,omitempty"`
	OrganizationUnit *[]string `json:"organizationUnit,omitempty"`
}

type IssuerInfo struct {
	CommonName   *string   `json:"commonName,omitempty"`
	Organization *[]string `json:"organization,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMCredential is the Schema for the asmcredentials API
// +genclient
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=asmcredentials,scope=Namespaced
type ASMCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMCredentialSpec   `json:"spec,omitempty"`
	Status ASMCredentialStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMCredentialList contains a list of ASMCredential
type ASMCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMCredential `json:"items"`
}