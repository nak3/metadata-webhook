/*
Copyright 2020 The Knative Authors

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

package defaults

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"knative.dev/pkg/apis"
)

const (
	sidecarInject                = "sidecar.istio.io/inject"
	sidecarrewriteAppHTTPProbers = "sidecar.istio.io/rewriteAppHTTPProbers"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TargetDeployment is a wrapper around Deployment.
type TargetDeployment struct {
	appsv1.Deployment `json:",inline"`
}

// Verify that Deployment adheres to the appropriate interfaces.
var (
	// Check that Deployment can be defaulted.
	_ apis.Defaultable = (*TargetDeployment)(nil)
	_ apis.Validatable = (*TargetDeployment)(nil)
)

// SetDefaults implements apis.Defaultable
func (r *TargetDeployment) SetDefaults(ctx context.Context) {
	if r.Spec.Template.Annotations == nil {
		r.Spec.Template.Annotations = make(map[string]string)
	}

	r.Spec.Template.Annotations[sidecarInject] = "true"
	r.Spec.Template.Annotations[sidecarrewriteAppHTTPProbers] = "true"
}

// Validate returns nil due to no need for validation
func (r *TargetDeployment) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
