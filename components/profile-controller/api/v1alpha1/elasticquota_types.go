package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

// ElasticQuota sets elastic quota restrictions per namespace
type ElasticQuota struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	// ElasticQuotaSpec defines the Min and Max for Quota.
	// +kubebuilder:validation:Optional
	Spec ElasticQuotaSpec `json:"spec,omitempty"`

	// ElasticQuotaStatus defines the observed use.
	// +kubebuilder:validation:Optional
	Status ElasticQuotaStatus `json:"status,omitempty"`
}

// ElasticQuotaSpec defines the Min and Max for Quota.
type ElasticQuotaSpec struct {
	// Min is the set of desired guaranteed limits for each named resource.
	// +optional
	Min v1.ResourceList `json:"min,omitempty"`

	// Max is the set of desired max limits for each named resource. The usage of max is based on the resource configurations of
	// successfully scheduled pods.
	// +optional
	Max v1.ResourceList `json:"max,omitempty"`
}

// ElasticQuotaStatus defines the observed use.
type ElasticQuotaStatus struct {
	// Used is the current observed total usage of the resource in the namespace.
	// +kubebuilder:validation:Optional
	Used v1.ResourceList `json:"used,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticQuotaList is a list of ElasticQuota items.
type ElasticQuotaList struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of ElasticQuota objects.
	Items []ElasticQuota `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ElasticQuota{}, &ElasticQuotaList{})
}
