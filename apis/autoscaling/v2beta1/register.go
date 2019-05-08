package v2beta1

import "github.com/parryjacob/k8s"

func init() {
	k8s.Register("autoscaling", "v2beta1", "horizontalpodautoscalers", true, &HorizontalPodAutoscaler{})

	k8s.RegisterList("autoscaling", "v2beta1", "horizontalpodautoscalers", true, &HorizontalPodAutoscalerList{})
}
