/*


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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&Pool{}, &PoolList{})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Target Namespace",type=string,JSONPath=`.status.namespace.name`
// +kubebuilder:printcolumn:name="apps",type=string,JSONPath=`.status.apps`
// +kubebuilder:printcolumn:name="quota",type=string,JSONPath=`.spec.appQuotaLimit`

// Pool is the Schema for the pools API
type Pool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PoolSpec   `json:"spec,omitempty"`
	Status PoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PoolList contains a list of Pool
type PoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pool `json:"items"`
}

// PoolSpec defines the desired state of Pool
type PoolSpec struct {
	// +kubebuilder:validation:MinLength=1
	NamespaceName string `json:"namespace"`

	AppQuotaLimit int `json:"appQuotaLimit"`

	IngressController IngressControllerSpec `json:"ingressController,omitempty"`
}

type PoolPhase string

const (
	PoolCreated PoolPhase = "Created"
	PoolFailed  PoolPhase = "Failed"
)

type Traefik struct {
	EntryPoints []string `json:"entryPoints"`
}

type IngressControllerSpec struct {
	ClassName       string   `json:"className,omitempty"`
	Domain          string   `json:"domain,omitempty"`
	ServiceEndpoint string   `json:"serviceEndpoint,omitempty"`
	Traefik         *Traefik `json:"traefik,omitempty"`
}

// PoolStatus defines the observed state of Pool
type PoolStatus struct {
	Phase   PoolPhase `json:"phase,omitempty"`
	Message string    `json:"message,omitempty"`

	Namespace *v1.ObjectReference `json:"namespace,omitempty"`
	Apps      []string            `json:"apps,omitempty"`
}

func (p *Pool) HasApp(name string) bool {
	for _, appName := range p.Status.Apps {
		if appName == name {
			return true
		}
	}
	return false
}