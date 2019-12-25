package v1beta1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ConfigHpa struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigHpaSpec   `json:"spec,omitempty"`
	Status ConfigHpaStatus `json:"status,omitempty"`
}

// specification of horizontal pod autoscaler
// was copied from HorizontalPodAutoscalerSpec + HPAControllerConfiguration
type ConfigHpaSpec struct {
	// part of HorizontalController, see comments in the k8s repo: pkg/controller/podautoscaler/horizontal.go
	DownscaleForbiddenWindowSeconds int32 `json:"downscaleForbiddenWindowSeconds,omitempty"`
	UpscaleForbiddenWindowSeconds   int32 `json:"upscaleForbiddenWindowSeconds,omitempty"`
	// See the comment about this parameter above
	ScaleUpLimitFactor float64 `json:"scaleUpLimitFactor,omitempty"`
	// See the comment about this parameter above
	ScaleDownLimitFactor float64 `json:"scaleDownLimitFactor,omitempty"`
	Tolerance            float64 `json:"tolerance,omitempty"`

	// part of HorizontalPodAutoscalerSpec, see comments in the k8s-1.10.8 repo: staging/src/k8s.io/api/autoscaling/v1/types.go
	// reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption
	// and will set the desired number of pods by using its Scale subresource.
	ScaleTargetRef CrossVersionObjectReference `json:"scaleTargetRef"`
	// specifications that will be used to calculate the desired replica count
	Metrics     []autoscalingv2.MetricSpec `json:"metrics,omitempty"`
	MinReplicas *int32                     `json:"minReplicas,omitempty"`
	MaxReplicas int32                      `json:"maxReplicas"`
}

// CrossVersionObjectReference contains enough information to let you identify the referred resource.
type CrossVersionObjectReference struct {
	// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds"
	Kind string `json:"kind"`
	// Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name"`
	// API version of the referent
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`
}

// ConfigHpaStatus defines the observed state of CHPA
type ConfigHpaStatus struct {
	ObservedGeneration *int64                                           `json:"observedGeneration,omitempty"`
	LastScaleTime      *metav1.Time                                     `json:"lastScaleTime,omitempty"`
	CurrentReplicas    int32                                            `json:"currentReplicas"`
	DesiredReplicas    int32                                            `json:"desiredReplicas"`
	CurrentMetrics     []autoscalingv2.MetricStatus                     `json:"currentMetrics"`
	Conditions         []autoscalingv2.HorizontalPodAutoscalerCondition `json:"conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigHpaList contains a list of ConfigHpa
type ConfigHpaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigHpa `json:"items"`
}
