/*
Copyright 2022.

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

//
// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BenchmarkOperatorSpec defines the desired state of BenchmarkOperator
type BenchmarkOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Target  *BenchmarkTarget    `json:"target"`
	Dataset *BenchmarkDataset   `json:"dataset"`
	Jobs    []*BenchmarkJobSpec `json:"jobs,omitempty"`
}

// BenchmarkTarget defines the desired state of BenchmarkTarget
type BenchmarkTarget struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// BenchmarkDataset defines the desired state of BenchmarkDateset
type BenchmarkDataset struct {
	Name    string                 `json:"name"`
	Group   string                 `json:"group"`
	Indexes int                    `json:"indexes"`
	Range   *BenchmarkDatasetRange `json:"range"`
}

// BenchmarkDatasetRange defines the desired state of BenchmarkDatesetRange
type BenchmarkDatasetRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// BenchmarkJobSpec defines the desired state of BenchmarkJob
type BenchmarkJobSpec struct {
	// +optional
	Target *BenchmarkTarget `json:"target"`
	// +optional
	Dataset *BenchmarkDataset `json:"dataset"`
	// +optional
	Replica int `json:"replica"`
	// +optional
	Repetition int    `json:"repetition"`
	JobType    string `json:"job_type"`
	Dimension  int    `json:"dimension"`
	// +optional
	Epsilon float32 `json:"epsilon"`
	// +optional
	Radius  float32 `json:"radius"`
	Iter    int     `json:"iter"`
	Num     int32   `json:"num"`
	MinNUm  int32   `json:"min_num"`
	Timeout string  `json:"timeout"`
	// +optional
	Rules []*BenchmarkJobRule `json:"rules,omitempty"`
}

// BenchmarkJobRule defines the desired state of BenchmarkJobRule
type BenchmarkJobRule struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// BenchmarkOperatorStatus defines the observed state of BenchmarkOperator
//+kubebuilder:validation:Enum=NotReady;Available;Healthy
type BenchmarkOperatorStatus string

const (
	BenchmarkOperatorNotReady  BenchmarkOperatorStatus = "NotReady"
	BenchmarkOperatorAvailable BenchmarkOperatorStatus = "Available"
	BenchmarkOperatorHealthy   BenchmarkOperatorStatus = "Healthy"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status"

// BenchmarkOperator is the Schema for the markdownviews API
type BenchmarkOperator struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BenchmarkOperatorSpec   `json:"spec,omitempty"`
	Status BenchmarkOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BenchmarkOperatorList contains a list of BenchmarkOperator
type BenchmarkOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BenchmarkOperator `json:"items"`
}

// BenchmarkJob is the Schema for the markdownviews API
type BenchmarkJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BenchmarkJobSpec        `json:"spec,omitempty"`
	Status BenchmarkOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BenchmarkJobList contains a list of BenchmarkOperator
type BenchmarkJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BenchmarkJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&BenchmarkOperator{},
		&BenchmarkOperatorList{},
	)
}
