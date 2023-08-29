//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and cap-operator contributors
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDomains) DeepCopyInto(out *ApplicationDomains) {
	*out = *in
	if in.Secondary != nil {
		in, out := &in.Secondary, &out.Secondary
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IstioIngressGatewayLabels != nil {
		in, out := &in.IstioIngressGatewayLabels, &out.IstioIngressGatewayLabels
		*out = make([]NameValue, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDomains.
func (in *ApplicationDomains) DeepCopy() *ApplicationDomains {
	if in == nil {
		return nil
	}
	out := new(ApplicationDomains)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BTP) DeepCopyInto(out *BTP) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ServiceInfo, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BTP.
func (in *BTP) DeepCopy() *BTP {
	if in == nil {
		return nil
	}
	out := new(BTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BTPTenantIdentification) DeepCopyInto(out *BTPTenantIdentification) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BTPTenantIdentification.
func (in *BTPTenantIdentification) DeepCopy() *BTPTenantIdentification {
	if in == nil {
		return nil
	}
	out := new(BTPTenantIdentification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPApplication) DeepCopyInto(out *CAPApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPApplication.
func (in *CAPApplication) DeepCopy() *CAPApplication {
	if in == nil {
		return nil
	}
	out := new(CAPApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPApplicationList) DeepCopyInto(out *CAPApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CAPApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPApplicationList.
func (in *CAPApplicationList) DeepCopy() *CAPApplicationList {
	if in == nil {
		return nil
	}
	out := new(CAPApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPApplicationSpec) DeepCopyInto(out *CAPApplicationSpec) {
	*out = *in
	in.Domains.DeepCopyInto(&out.Domains)
	out.Provider = in.Provider
	in.BTP.DeepCopyInto(&out.BTP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPApplicationSpec.
func (in *CAPApplicationSpec) DeepCopy() *CAPApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(CAPApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPApplicationStatus) DeepCopyInto(out *CAPApplicationStatus) {
	*out = *in
	in.GenericStatus.DeepCopyInto(&out.GenericStatus)
	in.LastFullReconciliationTime.DeepCopyInto(&out.LastFullReconciliationTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPApplicationStatus.
func (in *CAPApplicationStatus) DeepCopy() *CAPApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(CAPApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPApplicationVersion) DeepCopyInto(out *CAPApplicationVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPApplicationVersion.
func (in *CAPApplicationVersion) DeepCopy() *CAPApplicationVersion {
	if in == nil {
		return nil
	}
	out := new(CAPApplicationVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPApplicationVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPApplicationVersionList) DeepCopyInto(out *CAPApplicationVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CAPApplicationVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPApplicationVersionList.
func (in *CAPApplicationVersionList) DeepCopy() *CAPApplicationVersionList {
	if in == nil {
		return nil
	}
	out := new(CAPApplicationVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPApplicationVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPApplicationVersionSpec) DeepCopyInto(out *CAPApplicationVersionSpec) {
	*out = *in
	if in.RegistrySecrets != nil {
		in, out := &in.RegistrySecrets, &out.RegistrySecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]WorkloadDetails, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TenantOperations != nil {
		in, out := &in.TenantOperations, &out.TenantOperations
		*out = new(TenantOperations)
		(*in).DeepCopyInto(*out)
	}
	if in.ContentJobs != nil {
		in, out := &in.ContentJobs, &out.ContentJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPApplicationVersionSpec.
func (in *CAPApplicationVersionSpec) DeepCopy() *CAPApplicationVersionSpec {
	if in == nil {
		return nil
	}
	out := new(CAPApplicationVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPApplicationVersionStatus) DeepCopyInto(out *CAPApplicationVersionStatus) {
	*out = *in
	in.GenericStatus.DeepCopyInto(&out.GenericStatus)
	if in.FinishedJobs != nil {
		in, out := &in.FinishedJobs, &out.FinishedJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPApplicationVersionStatus.
func (in *CAPApplicationVersionStatus) DeepCopy() *CAPApplicationVersionStatus {
	if in == nil {
		return nil
	}
	out := new(CAPApplicationVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPTenant) DeepCopyInto(out *CAPTenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPTenant.
func (in *CAPTenant) DeepCopy() *CAPTenant {
	if in == nil {
		return nil
	}
	out := new(CAPTenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPTenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPTenantList) DeepCopyInto(out *CAPTenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CAPTenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPTenantList.
func (in *CAPTenantList) DeepCopy() *CAPTenantList {
	if in == nil {
		return nil
	}
	out := new(CAPTenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPTenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPTenantOperation) DeepCopyInto(out *CAPTenantOperation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPTenantOperation.
func (in *CAPTenantOperation) DeepCopy() *CAPTenantOperation {
	if in == nil {
		return nil
	}
	out := new(CAPTenantOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPTenantOperation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPTenantOperationList) DeepCopyInto(out *CAPTenantOperationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CAPTenantOperation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPTenantOperationList.
func (in *CAPTenantOperationList) DeepCopy() *CAPTenantOperationList {
	if in == nil {
		return nil
	}
	out := new(CAPTenantOperationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPTenantOperationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPTenantOperationSpec) DeepCopyInto(out *CAPTenantOperationSpec) {
	*out = *in
	out.BTPTenantIdentification = in.BTPTenantIdentification
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]CAPTenantOperationStep, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPTenantOperationSpec.
func (in *CAPTenantOperationSpec) DeepCopy() *CAPTenantOperationSpec {
	if in == nil {
		return nil
	}
	out := new(CAPTenantOperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPTenantOperationStatus) DeepCopyInto(out *CAPTenantOperationStatus) {
	*out = *in
	in.GenericStatus.DeepCopyInto(&out.GenericStatus)
	if in.CurrentStep != nil {
		in, out := &in.CurrentStep, &out.CurrentStep
		*out = new(uint32)
		**out = **in
	}
	if in.ActiveJob != nil {
		in, out := &in.ActiveJob, &out.ActiveJob
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPTenantOperationStatus.
func (in *CAPTenantOperationStatus) DeepCopy() *CAPTenantOperationStatus {
	if in == nil {
		return nil
	}
	out := new(CAPTenantOperationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPTenantOperationStep) DeepCopyInto(out *CAPTenantOperationStep) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPTenantOperationStep.
func (in *CAPTenantOperationStep) DeepCopy() *CAPTenantOperationStep {
	if in == nil {
		return nil
	}
	out := new(CAPTenantOperationStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPTenantSpec) DeepCopyInto(out *CAPTenantSpec) {
	*out = *in
	out.BTPTenantIdentification = in.BTPTenantIdentification
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPTenantSpec.
func (in *CAPTenantSpec) DeepCopy() *CAPTenantSpec {
	if in == nil {
		return nil
	}
	out := new(CAPTenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPTenantStatus) DeepCopyInto(out *CAPTenantStatus) {
	*out = *in
	in.GenericStatus.DeepCopyInto(&out.GenericStatus)
	if in.PreviousCAPApplicationVersions != nil {
		in, out := &in.PreviousCAPApplicationVersions, &out.PreviousCAPApplicationVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.LastFullReconciliationTime.DeepCopyInto(&out.LastFullReconciliationTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPTenantStatus.
func (in *CAPTenantStatus) DeepCopy() *CAPTenantStatus {
	if in == nil {
		return nil
	}
	out := new(CAPTenantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerDetails) DeepCopyInto(out *ContainerDetails) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerDetails.
func (in *ContainerDetails) DeepCopy() *ContainerDetails {
	if in == nil {
		return nil
	}
	out := new(ContainerDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentDetails) DeepCopyInto(out *DeploymentDetails) {
	*out = *in
	in.ContainerDetails.DeepCopyInto(&out.ContainerDetails)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Ports, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentDetails.
func (in *DeploymentDetails) DeepCopy() *DeploymentDetails {
	if in == nil {
		return nil
	}
	out := new(DeploymentDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericStatus) DeepCopyInto(out *GenericStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericStatus.
func (in *GenericStatus) DeepCopy() *GenericStatus {
	if in == nil {
		return nil
	}
	out := new(GenericStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDetails) DeepCopyInto(out *JobDetails) {
	*out = *in
	in.ContainerDetails.DeepCopyInto(&out.ContainerDetails)
	if in.BackoffLimit != nil {
		in, out := &in.BackoffLimit, &out.BackoffLimit
		*out = new(int32)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDetails.
func (in *JobDetails) DeepCopy() *JobDetails {
	if in == nil {
		return nil
	}
	out := new(JobDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameValue) DeepCopyInto(out *NameValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameValue.
func (in *NameValue) DeepCopy() *NameValue {
	if in == nil {
		return nil
	}
	out := new(NameValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ports) DeepCopyInto(out *Ports) {
	*out = *in
	if in.AppProtocol != nil {
		in, out := &in.AppProtocol, &out.AppProtocol
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ports.
func (in *Ports) DeepCopy() *Ports {
	if in == nil {
		return nil
	}
	out := new(Ports)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInfo) DeepCopyInto(out *ServiceInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInfo.
func (in *ServiceInfo) DeepCopy() *ServiceInfo {
	if in == nil {
		return nil
	}
	out := new(ServiceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantOperationWorkloadReference) DeepCopyInto(out *TenantOperationWorkloadReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantOperationWorkloadReference.
func (in *TenantOperationWorkloadReference) DeepCopy() *TenantOperationWorkloadReference {
	if in == nil {
		return nil
	}
	out := new(TenantOperationWorkloadReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantOperations) DeepCopyInto(out *TenantOperations) {
	*out = *in
	if in.Provisioning != nil {
		in, out := &in.Provisioning, &out.Provisioning
		*out = make([]TenantOperationWorkloadReference, len(*in))
		copy(*out, *in)
	}
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = make([]TenantOperationWorkloadReference, len(*in))
		copy(*out, *in)
	}
	if in.Deprovisioning != nil {
		in, out := &in.Deprovisioning, &out.Deprovisioning
		*out = make([]TenantOperationWorkloadReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantOperations.
func (in *TenantOperations) DeepCopy() *TenantOperations {
	if in == nil {
		return nil
	}
	out := new(TenantOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadDetails) DeepCopyInto(out *WorkloadDetails) {
	*out = *in
	if in.ConsumedBTPServices != nil {
		in, out := &in.ConsumedBTPServices, &out.ConsumedBTPServices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeploymentDefinition != nil {
		in, out := &in.DeploymentDefinition, &out.DeploymentDefinition
		*out = new(DeploymentDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.JobDefinition != nil {
		in, out := &in.JobDefinition, &out.JobDefinition
		*out = new(JobDetails)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadDetails.
func (in *WorkloadDetails) DeepCopy() *WorkloadDetails {
	if in == nil {
		return nil
	}
	out := new(WorkloadDetails)
	in.DeepCopyInto(out)
	return out
}
