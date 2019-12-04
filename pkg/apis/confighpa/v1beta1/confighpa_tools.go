package v1beta1

import (
	"fmt"
)

func (chpa ConfigHpa) String() string {
	ret := fmt.Sprintf("Spec:%v Status:%v", chpa.Spec, chpa.Status)
	return ret
}

func (chpa_spec ConfigHpaSpec) String() string {
	minReplicas := "<nil>"
	if chpa_spec.MinReplicas != nil {
		minReplicas = fmt.Sprintf("%v", *chpa_spec.MinReplicas)
	}
	ret := fmt.Sprintf("{Ref:%v/%v DFWS:%v UFWS:%v SULF:%v SULM:%v T:%v MinR:%v MaxR:%v M:%v}",
		chpa_spec.ScaleTargetRef.Kind,
		chpa_spec.ScaleTargetRef.Name,
		chpa_spec.DownscaleForbiddenWindowSeconds,
		chpa_spec.UpscaleForbiddenWindowSeconds,
		chpa_spec.ScaleUpLimitFactor,
		chpa_spec.ScaleDownLimitFactor,
		chpa_spec.Tolerance,
		minReplicas,
		chpa_spec.MaxReplicas,
		chpa_spec.Metrics)
	return ret
}

func (chpa_status ConfigHpaStatus) String() string {
	lastScaleTime := "<nil>"
	if chpa_status.LastScaleTime != nil {
		lastScaleTime = fmt.Sprintf("%v", *chpa_status.LastScaleTime)
	}
	ret := fmt.Sprintf("{LST:%v CR:%v DR:%v}", lastScaleTime, chpa_status.CurrentReplicas, chpa_status.DesiredReplicas)
	return ret
}
