// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactInfo) DeepCopyInto(out *ArtifactInfo) {
	*out = *in
	if in.ServicePorts != nil {
		in, out := &in.ServicePorts, &out.ServicePorts
		*out = make([]ServicePorts, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactInfo.
func (in *ArtifactInfo) DeepCopy() *ArtifactInfo {
	if in == nil {
		return nil
	}
	out := new(ArtifactInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CD) DeepCopyInto(out *CD) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CD.
func (in *CD) DeepCopy() *CD {
	if in == nil {
		return nil
	}
	out := new(CD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CD) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDList) DeepCopyInto(out *CDList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CD, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDList.
func (in *CDList) DeepCopy() *CDList {
	if in == nil {
		return nil
	}
	out := new(CDList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDSpec) DeepCopyInto(out *CDSpec) {
	*out = *in
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.ServiceImage != nil {
		in, out := &in.ServiceImage, &out.ServiceImage
		*out = new(string)
		**out = **in
	}
	if in.DeployNamespace != nil {
		in, out := &in.DeployNamespace, &out.DeployNamespace
		*out = new(string)
		**out = **in
	}
	in.ArtifactInfo.DeepCopyInto(&out.ArtifactInfo)
	if in.DeployType != nil {
		in, out := &in.DeployType, &out.DeployType
		*out = new(string)
		**out = **in
	}
	if in.FlowId != nil {
		in, out := &in.FlowId, &out.FlowId
		*out = new(string)
		**out = **in
	}
	if in.StepName != nil {
		in, out := &in.StepName, &out.StepName
		*out = new(string)
		**out = **in
	}
	if in.AckStates != nil {
		in, out := &in.AckStates, &out.AckStates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDSpec.
func (in *CDSpec) DeepCopy() *CDSpec {
	if in == nil {
		return nil
	}
	out := new(CDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CI) DeepCopyInto(out *CI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CI.
func (in *CI) DeepCopy() *CI {
	if in == nil {
		return nil
	}
	out := new(CI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CIList) DeepCopyInto(out *CIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIList.
func (in *CIList) DeepCopy() *CIList {
	if in == nil {
		return nil
	}
	out := new(CIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISpec) DeepCopyInto(out *CISpec) {
	*out = *in
	if in.GitURL != nil {
		in, out := &in.GitURL, &out.GitURL
		*out = new(string)
		**out = **in
	}
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.CommitID != nil {
		in, out := &in.CommitID, &out.CommitID
		*out = new(string)
		**out = **in
	}
	if in.RetryCount != nil {
		in, out := &in.RetryCount, &out.RetryCount
		*out = new(uint32)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(string)
		**out = **in
	}
	if in.FlowId != nil {
		in, out := &in.FlowId, &out.FlowId
		*out = new(string)
		**out = **in
	}
	if in.StepName != nil {
		in, out := &in.StepName, &out.StepName
		*out = new(string)
		**out = **in
	}
	if in.AckStates != nil {
		in, out := &in.AckStates, &out.AckStates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISpec.
func (in *CISpec) DeepCopy() *CISpec {
	if in == nil {
		return nil
	}
	out := new(CISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePorts) DeepCopyInto(out *ServicePorts) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePorts.
func (in *ServicePorts) DeepCopy() *ServicePorts {
	if in == nil {
		return nil
	}
	out := new(ServicePorts)
	in.DeepCopyInto(out)
	return out
}
