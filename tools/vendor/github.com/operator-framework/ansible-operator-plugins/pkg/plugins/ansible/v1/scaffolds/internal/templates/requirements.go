// Copyright 2020 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templates

import (
	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"
)

var _ machinery.Template = &RequirementsYml{}

// RequirementsYml - A requirements file for Ansible collection dependencies
type RequirementsYml struct {
	machinery.TemplateMixin
}

func (f *RequirementsYml) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = "requirements.yml"
	}
	f.TemplateBody = requirementsYmlTmpl
	return nil
}

const requirementsYmlTmpl = `---
collections:
  - name: operator_sdk.util
    version: "0.5.0"
  - name: kubernetes.core
    version: "3.2.0"
  - name: cloud.common
    version: "3.0.0"
  - name: community.docker
    version: "3.12.1"
`
