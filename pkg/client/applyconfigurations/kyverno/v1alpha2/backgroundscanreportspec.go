/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/policyreport/v1alpha2"
)

// BackgroundScanReportSpecApplyConfiguration represents an declarative configuration of the BackgroundScanReportSpec type for use
// with apply.
type BackgroundScanReportSpecApplyConfiguration struct {
	Summary *v1alpha2.PolicyReportSummaryApplyConfiguration `json:"summary,omitempty"`
	Results []v1alpha2.PolicyReportResultApplyConfiguration `json:"results,omitempty"`
}

// BackgroundScanReportSpecApplyConfiguration constructs an declarative configuration of the BackgroundScanReportSpec type for use with
// apply.
func BackgroundScanReportSpec() *BackgroundScanReportSpecApplyConfiguration {
	return &BackgroundScanReportSpecApplyConfiguration{}
}

// WithSummary sets the Summary field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Summary field is set to the value of the last call.
func (b *BackgroundScanReportSpecApplyConfiguration) WithSummary(value *v1alpha2.PolicyReportSummaryApplyConfiguration) *BackgroundScanReportSpecApplyConfiguration {
	b.Summary = value
	return b
}

// WithResults adds the given value to the Results field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Results field.
func (b *BackgroundScanReportSpecApplyConfiguration) WithResults(values ...*v1alpha2.PolicyReportResultApplyConfiguration) *BackgroundScanReportSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResults")
		}
		b.Results = append(b.Results, *values[i])
	}
	return b
}
