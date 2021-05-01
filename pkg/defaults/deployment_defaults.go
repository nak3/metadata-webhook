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
	// servingGroupName is the group name for knative labels and annotations
	servingGroupName = "serving.knative.dev"

	// servingRevisionLabelKey is the label key attached to k8s resources to indicate
	// which Revision triggered their creation.
	servingRevisionLabelKey = servingGroupName + "/revision"
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
	if r.Labels == nil {
		r.Labels = make(map[string]string)
	}

	if r.Spec.Template.Labels == nil {
		r.Spec.Template.Labels = make(map[string]string)
	}

	revisionName := r.Labels[servingRevisionLabelKey]
	if revisionName != "" {
		r.Labels["todo"] = revisionName
		r.Spec.Template.Labels["todo"] = revisionName
	}
}

// Validate returns nil due to no need for validation
func (r *TargetDeployment) Validate(ctx context.Context) *apis.FieldError {
	return nil
}