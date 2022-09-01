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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ValdBenchmarkOperatorSpec defines the desired state of ValdBenchmarkOperator
type ValdBenchmarkOperatorSpec struct {
	Target  *BenchmarkTarget        `json:"target" yaml:"target"`
	Dataset *BenchmarkDataset       `json:"dataset" yaml:"dataset"`
	Jobs    []*ValdBenchmarkJobSpec `json:"jobs" yaml:"jobs"`
}

type ValdBenchmarkJobSpec struct {
	// +optional
	Target *BenchmarkTarget `json:"target" yaml:"target"`
	// +optional
	Dataset *BenchmarkDataset `json:"dataset" yaml:"dataset"`
	// +kubebuilder:default=1
	// +optional
	Replica int `json:"replica" yaml:"replica"`
	// +kubebuilder:default=1
	// +optional
	Repetition int    `json:"repetition" yaml:"repetition"`
	JobType    string `json:"job_type" yaml:"job_type"`
	Dimension  int    `json:"dimension" yaml:"dimension"`
	// +kubebuilder:default=0.2
	// +optional
	Epsilon float32 `json:"epsilon" yaml:"epsilon"`
	// +kubebuilder:default=-1
	// +optional
	Radius  float32 `json:"radius" yaml:"radius"`
	Iter    int     `json:"iter" yaml:"iter"`
	Num     int32   `json:"num" yaml:"num"`
	MinNUm  int32   `json:"min_num" yaml:"min_num"`
	Timeout string  `json:"timeout" yaml:"timeout"`
	// +optional
	Rules []*BenchmarkJobRule `json:"rules" yaml:"rules"`
}

type BenchmarkTarget struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

type BenchmarkDataset struct {
	Name    string                 `json:"name" yaml:"name"`
	Group   string                 `json:"group" yaml:"group"`
	Indexes int                    `json:"indexes" yaml:"indexes"`
	Range   *BenchmarkDatasetRange `json:"range" yaml:"range"`
}

type BenchmarkDatasetRange struct {
	Start int `json:"start" yaml:"start"`
	End   int `json:"end" yaml:"end"`
}

type BenchmarkJobRule struct {
	Name string `json:"name" yaml:"name"`
	Type string `json:"type" yaml:"type"`
}

// ValdBenchmarkOperatorStatus defines the observed state of ValdBenchmarkOperator
type ValdBenchmarkOperatorStatus string

const (
	ValdBenchmarkOperatorNotReady  = ValdBenchmarkOperatorStatus("NotReady")
	ValdBenchmarkOperatorAvailable = ValdBenchmarkOperatorStatus("Available")
	ValdBenchmarkOperatorHealthy   = ValdBenchmarkOperatorStatus("Healthy")
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:validation:Enum=NotReady;Available;Healthy
//+kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status"

// ValdBenchmarkOperator is the Schema for the valdbenchmarkoperators API
type ValdBenchmarkOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ValdBenchmarkOperatorSpec   `json:"spec,omitempty"`
	Status ValdBenchmarkOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ValdBenchmarkOperatorList contains a list of ValdBenchmarkOperator
type ValdBenchmarkOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ValdBenchmarkOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ValdBenchmarkOperator{}, &ValdBenchmarkOperatorList{})
}
