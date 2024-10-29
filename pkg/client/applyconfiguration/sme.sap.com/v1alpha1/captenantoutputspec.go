/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and cap-operator contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// CAPTenantOutputSpecApplyConfiguration represents a declarative configuration of the CAPTenantOutputSpec type for use
// with apply.
type CAPTenantOutputSpecApplyConfiguration struct {
	SubscriptionCallbackData *string `json:"subscriptionCallbackData,omitempty"`
}

// CAPTenantOutputSpecApplyConfiguration constructs a declarative configuration of the CAPTenantOutputSpec type for use with
// apply.
func CAPTenantOutputSpec() *CAPTenantOutputSpecApplyConfiguration {
	return &CAPTenantOutputSpecApplyConfiguration{}
}

// WithSubscriptionCallbackData sets the SubscriptionCallbackData field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubscriptionCallbackData field is set to the value of the last call.
func (b *CAPTenantOutputSpecApplyConfiguration) WithSubscriptionCallbackData(value string) *CAPTenantOutputSpecApplyConfiguration {
	b.SubscriptionCallbackData = &value
	return b
}