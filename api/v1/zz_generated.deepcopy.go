//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BenchmarkDataset) DeepCopyInto(out *BenchmarkDataset) {
	*out = *in
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = new(BenchmarkDatasetRange)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BenchmarkDataset.
func (in *BenchmarkDataset) DeepCopy() *BenchmarkDataset {
	if in == nil {
		return nil
	}
	out := new(BenchmarkDataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BenchmarkDatasetRange) DeepCopyInto(out *BenchmarkDatasetRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BenchmarkDatasetRange.
func (in *BenchmarkDatasetRange) DeepCopy() *BenchmarkDatasetRange {
	if in == nil {
		return nil
	}
	out := new(BenchmarkDatasetRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BenchmarkJobRule) DeepCopyInto(out *BenchmarkJobRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BenchmarkJobRule.
func (in *BenchmarkJobRule) DeepCopy() *BenchmarkJobRule {
	if in == nil {
		return nil
	}
	out := new(BenchmarkJobRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BenchmarkTarget) DeepCopyInto(out *BenchmarkTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BenchmarkTarget.
func (in *BenchmarkTarget) DeepCopy() *BenchmarkTarget {
	if in == nil {
		return nil
	}
	out := new(BenchmarkTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValdBenchmarkJobSpec) DeepCopyInto(out *ValdBenchmarkJobSpec) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(BenchmarkTarget)
		**out = **in
	}
	if in.Dataset != nil {
		in, out := &in.Dataset, &out.Dataset
		*out = new(BenchmarkDataset)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]*BenchmarkJobRule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BenchmarkJobRule)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValdBenchmarkJobSpec.
func (in *ValdBenchmarkJobSpec) DeepCopy() *ValdBenchmarkJobSpec {
	if in == nil {
		return nil
	}
	out := new(ValdBenchmarkJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValdBenchmarkOperator) DeepCopyInto(out *ValdBenchmarkOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValdBenchmarkOperator.
func (in *ValdBenchmarkOperator) DeepCopy() *ValdBenchmarkOperator {
	if in == nil {
		return nil
	}
	out := new(ValdBenchmarkOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValdBenchmarkOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValdBenchmarkOperatorList) DeepCopyInto(out *ValdBenchmarkOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ValdBenchmarkOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValdBenchmarkOperatorList.
func (in *ValdBenchmarkOperatorList) DeepCopy() *ValdBenchmarkOperatorList {
	if in == nil {
		return nil
	}
	out := new(ValdBenchmarkOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValdBenchmarkOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValdBenchmarkOperatorSpec) DeepCopyInto(out *ValdBenchmarkOperatorSpec) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(BenchmarkTarget)
		**out = **in
	}
	if in.Dataset != nil {
		in, out := &in.Dataset, &out.Dataset
		*out = new(BenchmarkDataset)
		(*in).DeepCopyInto(*out)
	}
	if in.Jobs != nil {
		in, out := &in.Jobs, &out.Jobs
		*out = make([]*ValdBenchmarkJobSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ValdBenchmarkJobSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValdBenchmarkOperatorSpec.
func (in *ValdBenchmarkOperatorSpec) DeepCopy() *ValdBenchmarkOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(ValdBenchmarkOperatorSpec)
	in.DeepCopyInto(out)
	return out
}
