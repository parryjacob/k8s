package v1beta1

import "github.com/parryjacob/k8s"

func init() {
	k8s.Register("batch", "v1beta1", "cronjobs", true, &CronJob{})

	k8s.RegisterList("batch", "v1beta1", "cronjobs", true, &CronJobList{})
}
