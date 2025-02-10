/*
Copyright 2024 Alibaba Cloud.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PacketInspectConditionType string

const (
	PacketInspectConditionInitializing PacketInspectConditionType = "Initializing"
	PacketInspectConditionInspecting   PacketInspectConditionType = "Inspecting"
	PacketInspectConditionComplete     PacketInspectConditionType = "Complete"
	PacketInspectConditionFailed       PacketInspectConditionType = "Failed"
)

// EDIT THIS FILE! THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type PodDescriptor struct {
	ClusterId string `json:"clusterId,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
}

// ASMPacketInspectorSpec defines the desired state of ASMPacketInspector
type ASMPacketInspectorSpec struct {
	Pod           PodDescriptor   `json:"pod,omitempty"`
	TcpDumpParams string          `json:"tcpDumpParams,omitempty"`
	Duration      metav1.Duration `json:"duration,omitempty"`
	FileName      string          `json:"fileName,omitempty"`
}

type ASMPacketInspectorCondition struct {
	Type    PacketInspectConditionType `json:"type,omitempty"`
	Time    metav1.Time                `json:"time,omitempty"`
	Reason  string                     `json:"reason,omitempty"`
	Message string                     `json:"message,omitempty"`
}

type NodeMetadata struct {
	Name   string `json:"name,omitempty"`
	Region string `json:"region,omitempty"`
	EcsId  string `json:"ecsId,omitempty"`
}

// ASMPacketInspectorStatus defines the observed state of ASMPacketInspector
type ASMPacketInspectorStatus struct {
	Phase         string                        `json:"phase,omitempty"`
	StartAt       metav1.Time                   `json:"startAt,omitempty"`
	CompletedAt   metav1.Time                   `json:"completedAt,omitempty"`
	Conditions    []ASMPacketInspectorCondition `json:"conditions,omitempty"`
	TaskId        string                        `json:"taskId,omitempty"`
	RunningOnNode *NodeMetadata                 `json:"runningOnNode,omitempty"`
	FilePath      string                        `json:"filePath,omitempty"`
}

// ASMPacketInspector is the Schema for the asmpacketinspectors API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type ASMPacketInspector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMPacketInspectorSpec   `json:"spec,omitempty"`
	Status ASMPacketInspectorStatus `json:"status,omitempty"`
}

// ASMPacketInspectorList contains a list of ASMPacketInspector
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ASMPacketInspectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMPacketInspector `json:"items"`
}