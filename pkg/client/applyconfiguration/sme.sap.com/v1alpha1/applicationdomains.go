/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and cap-operator contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ApplicationDomainsApplyConfiguration represents a declarative configuration of the ApplicationDomains type for use
// with apply.
type ApplicationDomainsApplyConfiguration struct {
	Primary                   *string                       `json:"primary,omitempty"`
	Secondary                 []string                      `json:"secondary,omitempty"`
	DnsTarget                 *string                       `json:"dnsTarget,omitempty"`
	IstioIngressGatewayLabels []NameValueApplyConfiguration `json:"istioIngressGatewayLabels,omitempty"`
}

// ApplicationDomainsApplyConfiguration constructs a declarative configuration of the ApplicationDomains type for use with
// apply.
func ApplicationDomains() *ApplicationDomainsApplyConfiguration {
	return &ApplicationDomainsApplyConfiguration{}
}

// WithPrimary sets the Primary field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Primary field is set to the value of the last call.
func (b *ApplicationDomainsApplyConfiguration) WithPrimary(value string) *ApplicationDomainsApplyConfiguration {
	b.Primary = &value
	return b
}

// WithSecondary adds the given value to the Secondary field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Secondary field.
func (b *ApplicationDomainsApplyConfiguration) WithSecondary(values ...string) *ApplicationDomainsApplyConfiguration {
	for i := range values {
		b.Secondary = append(b.Secondary, values[i])
	}
	return b
}

// WithDnsTarget sets the DnsTarget field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DnsTarget field is set to the value of the last call.
func (b *ApplicationDomainsApplyConfiguration) WithDnsTarget(value string) *ApplicationDomainsApplyConfiguration {
	b.DnsTarget = &value
	return b
}

// WithIstioIngressGatewayLabels adds the given value to the IstioIngressGatewayLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IstioIngressGatewayLabels field.
func (b *ApplicationDomainsApplyConfiguration) WithIstioIngressGatewayLabels(values ...*NameValueApplyConfiguration) *ApplicationDomainsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithIstioIngressGatewayLabels")
		}
		b.IstioIngressGatewayLabels = append(b.IstioIngressGatewayLabels, *values[i])
	}
	return b
}
