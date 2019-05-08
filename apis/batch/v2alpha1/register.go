package v2alpha1

import "github.com/parryjacob/k8s"

func init() {
	k8s.Register("batch", "v2alpha1", "cronjobs", true, &CronJob{})

	k8s.RegisterList("batch", "v2alpha1", "cronjobs", true, &CronJobList{})
}
